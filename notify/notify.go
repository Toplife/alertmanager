// Copyright 2015 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package notify

import (
	"fmt"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/model"
	"golang.org/x/net/context"

	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/inhibit"
	"github.com/prometheus/alertmanager/provider"
	meshprov "github.com/prometheus/alertmanager/provider/mesh"
	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/alertmanager/types"
)

var (
	numNotifications = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "alertmanager",
		Name:      "notifications_total",
		Help:      "The total number of attempted notifications.",
	}, []string{"integration"})

	numFailedNotifications = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "alertmanager",
		Name:      "notifications_failed_total",
		Help:      "The total number of failed notifications.",
	}, []string{"integration"})
)

func init() {
	prometheus.Register(numNotifications)
	prometheus.Register(numFailedNotifications)
}

// MinTimeout is the minimum timeout that is set for the context of a call
// to a notification pipeline.
const MinTimeout = 10 * time.Second

// notifyKey defines a custom type with which a context is populated to
// avoid accidental collisions.
type notifyKey int

const (
	keyReceiver notifyKey = iota
	keyRepeatInterval
	keyGroupLabels
	keyGroupKey
	keyNow
)

// WithReceiver populates a context with a receiver.
func WithReceiver(ctx context.Context, rcv string) context.Context {
	return context.WithValue(ctx, keyReceiver, rcv)
}

// WithRepeatInterval populates a context with a repeat interval.
func WithRepeatInterval(ctx context.Context, t time.Duration) context.Context {
	return context.WithValue(ctx, keyRepeatInterval, t)
}

// WithGroupKey populates a context with a group key.
func WithGroupKey(ctx context.Context, fp model.Fingerprint) context.Context {
	return context.WithValue(ctx, keyGroupKey, fp)
}

// WithGroupLabels populates a context with grouping labels.
func WithGroupLabels(ctx context.Context, lset model.LabelSet) context.Context {
	return context.WithValue(ctx, keyGroupLabels, lset)
}

// WithNow populates a context with a now timestamp.
func WithNow(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, keyNow, t)
}

func receiver(ctx context.Context) string {
	recv, ok := Receiver(ctx)
	if !ok {
		log.Error("missing receiver")
	}
	return recv
}

// Receiver extracts a receiver from the context. Iff none exists, the
// second argument is false.
func Receiver(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(keyReceiver).(string)
	return v, ok
}

// RepeatInterval extracts a repeat interval from the context. Iff none exists, the
// second argument is false.
func RepeatInterval(ctx context.Context) (time.Duration, bool) {
	v, ok := ctx.Value(keyRepeatInterval).(time.Duration)
	return v, ok
}

// GroupKey extracts a group key from the context. Iff none exists, the
// second argument is false.
func GroupKey(ctx context.Context) (model.Fingerprint, bool) {
	v, ok := ctx.Value(keyGroupKey).(model.Fingerprint)
	return v, ok
}

func groupLabels(ctx context.Context) model.LabelSet {
	groupLabels, ok := GroupLabels(ctx)
	if !ok {
		log.Error("missing group labels")
	}
	return groupLabels
}

// GroupLabels extracts grouping label set from the context. Iff none exists, the
// second argument is false.
func GroupLabels(ctx context.Context) (model.LabelSet, bool) {
	v, ok := ctx.Value(keyGroupLabels).(model.LabelSet)
	return v, ok
}

// Now extracts a now timestamp from the context. Iff none exists, the
// second argument is false.
func Now(ctx context.Context) (time.Time, bool) {
	v, ok := ctx.Value(keyNow).(time.Time)
	return v, ok
}

// A Stage processes alerts under the constraints of the given context.
type Stage interface {
	Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error)
}

// StageFunc wraps a function to represent a Stage.
type StageFunc func(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error)

// Exec implements Stage interface.
func (f StageFunc) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	return f(ctx, alerts...)
}

// BuildPipeline builds a map of receivers to Stages.
func BuildPipeline(
	confs []*config.Receiver,
	tmpl *template.Template,
	meshWait func() time.Duration,
	inhibitor *inhibit.Inhibitor,
	silences *meshprov.Silences,
	ni *meshprov.NotificationInfos,
	marker types.Marker,
) RoutingStage {
	rs := RoutingStage{}

	for _, rc := range confs {
		rs[rc.Name] = createStage(BuildReceiverIntegrations(rc, tmpl), meshWait, inhibitor, silences, ni, marker)
	}

	return rs
}

// createStage creates a pipeline of stages for a receiver.
func createStage(
	receiverIntegrations []Integration,
	meshWait func() time.Duration,
	inhibitor *inhibit.Inhibitor,
	silences *meshprov.Silences,
	ni *meshprov.NotificationInfos,
	marker types.Marker,
) Stage {
	var ms MultiStage
	ms = append(ms, NewLogStage(log.With("step", "inhibit")))
	ms = append(ms, NewInhibitStage(inhibitor, marker))
	ms = append(ms, NewLogStage(log.With("step", "silence")))
	ms = append(ms, NewSilenceStage(silences, marker))

	var fs = FanoutStage{}
	for _, i := range receiverIntegrations {
		var s MultiStage
		s = append(s, NewLogStage(log.With("step", "wait")))
		s = append(s, NewWaitStage(meshWait))
		s = append(s, NewLogStage(log.With("step", "filterResolved")))
		s = append(s, NewFilterResolvedStage(i.conf))
		s = append(s, NewLogStage(log.With("step", "dedup")))
		s = append(s, NewDedupStage(ni))
		s = append(s, NewLogStage(log.With("step", "integration")))
		s = append(s, NewRetryStage(i))
		s = append(s, NewLogStage(log.With("step", "newNotifies")))
		s = append(s, NewSetNotifiesStage(ni))
		fs[fmt.Sprintf("%s/%d", i.name, i.idx)] = s
	}

	return append(ms, fs)
}

// RoutingStage executes the inner stages based on the receiver specified in
// the context.
type RoutingStage map[string]Stage

// Exec implements the Stage interface.
func (rs RoutingStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	receiver, ok := Receiver(ctx)
	if !ok {
		return nil, fmt.Errorf("receiver missing")
	}

	s, ok := rs[receiver]
	if !ok {
		return nil, fmt.Errorf("stage for receiver missing")
	}

	return s.Exec(ctx, alerts...)
}

// A MultiStage executes a series of stages sequencially.
type MultiStage []Stage

// Exec implements the Stage interface.
func (ms MultiStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	var err error
	for _, s := range ms {
		if len(alerts) == 0 {
			return nil, nil
		}

		alerts, err = s.Exec(ctx, alerts...)
		if err != nil {
			return nil, err
		}
	}
	return alerts, nil
}

// FanoutStage executes its stages concurrently
type FanoutStage map[string]Stage

// Exec attempts to execute all stages concurrently. It returns a
// types.MultiError if any of them fails.
func (fs FanoutStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	var (
		wg sync.WaitGroup
		me types.MultiError
	)
	wg.Add(len(fs))

	receiver, ok := Receiver(ctx)
	if !ok {
		return nil, fmt.Errorf("receiver missing")
	}

	for suffix, s := range fs {
		// Suffix the receiver with the unique key for the fanout.
		foCtx := WithReceiver(ctx, fmt.Sprintf("%s/%s", receiver, suffix))

		go func(s Stage) {
			_, err := s.Exec(foCtx, alerts...)
			if err != nil {
				me.Add(err)
				log.Errorf("Error on notify: %s", err)
			}
			wg.Done()
		}(s)
	}

	wg.Wait()

	if me.Len() > 0 {
		return nil, &me
	}

	return alerts, nil
}

// LogStage logs the passed alerts with a given logger.
type LogStage struct {
	log log.Logger
}

func NewLogStage(log log.Logger) *LogStage {
	return &LogStage{log: log}
}

// Exec implements the Stage interface.
func (l *LogStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	l.log.Debugf("notify %v", alerts)

	return alerts, nil
}

// InhibitStage filters alerts through an inhibition muter.
type InhibitStage struct {
	muter  types.Muter
	marker types.Marker
}

// NewInhibitStage return a new InhibitStage.
func NewInhibitStage(m types.Muter, mk types.Marker) *InhibitStage {
	return &InhibitStage{
		muter:  m,
		marker: mk,
	}
}

// Exec implements the Stage interface.
func (n *InhibitStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	var filtered []*types.Alert
	for _, a := range alerts {
		ok := n.marker.Inhibited(a.Fingerprint())
		// TODO(fabxc): increment total alerts counter.
		// Do not send the alert if the silencer mutes it.
		if !n.muter.Mutes(a.Labels) {
			// TODO(fabxc): increment muted alerts counter.
			filtered = append(filtered, a)
			// Store whether a previously inhibited alert is firing again.
			a.WasInhibited = ok
		}
	}

	return filtered, nil
}

// SilenceStage filters alerts through a silence muter.
type SilenceStage struct {
	muter  types.Muter
	marker types.Marker
}

// NewSilenceStage returns a new SilenceStage.
func NewSilenceStage(m types.Muter, mk types.Marker) *SilenceStage {
	return &SilenceStage{
		muter:  m,
		marker: mk,
	}
}

// Exec implements the Stage interface.
func (n *SilenceStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	var filtered []*types.Alert
	for _, a := range alerts {
		_, ok := n.marker.Silenced(a.Fingerprint())
		// TODO(fabxc): increment total alerts counter.
		// Do not send the alert if the silencer mutes it.
		if !n.muter.Mutes(a.Labels) {
			// TODO(fabxc): increment muted alerts counter.
			filtered = append(filtered, a)
			// Store whether a previously silenced alert is firing again.
			a.WasSilenced = ok
		}
	}

	return filtered, nil
}

// WaitStage waits for a certain amount of time before continuing or until the
// context is done.
type WaitStage struct {
	wait func() time.Duration
}

// NewWaitStage returns a new WaitStage.
func NewWaitStage(wait func() time.Duration) *WaitStage {
	return &WaitStage{
		wait: wait,
	}
}

// Exec implements the Stage interface.
func (ws *WaitStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	select {
	case <-time.After(ws.wait()):
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return alerts, nil
}

// FilterResolvedStage filters alerts based on a given notifierConfig. Either
// returns all alerts or only those that are not resolved.
type FilterResolvedStage struct {
	conf notifierConfig
}

// NewFilterRecolvedStage returns a new instance of a FilterResolvedStage.
func NewFilterResolvedStage(conf notifierConfig) *FilterResolvedStage {
	return &FilterResolvedStage{
		conf: conf,
	}
}

// Exec implements the Stage interface.
func (fr *FilterResolvedStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	var res []*types.Alert

	if fr.conf.SendResolved() {
		res = alerts
	} else {
		for _, a := range alerts {
			if a.Status() != model.AlertResolved {
				res = append(res, a)
			}
		}
	}

	return res, nil
}

// DedupStage filters alerts.
// Filtering happens based on a provider of NotifyInfos.
type DedupStage struct {
	notifies provider.Notifies
}

// NewDedupStage wraps a DedupStage that runs against the given NotifyInfo provider.
func NewDedupStage(notifies provider.Notifies) *DedupStage {
	return &DedupStage{notifies}
}

// hasUpdates checks an alert against the last notification that was made
// about it.
func (n *DedupStage) hasUpdate(alert *types.Alert, last *types.NotificationInfo, now time.Time, interval time.Duration) bool {
	if last != nil {
		if alert.Resolved() {
			if last.Resolved {
				return false
			}
		} else if !last.Resolved {
			// Do not send again if last was delivered unless
			// the repeat interval has already passed.
			if !now.After(last.Timestamp.Add(interval)) {
				return false
			}
		}
	} else if alert.Resolved() {
		// If the alert is resolved but we never notified about it firing,
		// there is nothing to do.
		return false
	}
	return true
}

// Exec implements the Stage interface.
func (n *DedupStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	name, ok := Receiver(ctx)
	if !ok {
		return nil, fmt.Errorf("notifier name missing")
	}

	repeatInterval, ok := RepeatInterval(ctx)
	if !ok {
		return nil, fmt.Errorf("repeat interval missing")
	}

	now, ok := Now(ctx)
	if !ok {
		return nil, fmt.Errorf("now time missing")
	}

	var fps []model.Fingerprint
	for _, a := range alerts {
		fps = append(fps, a.Fingerprint())
	}

	notifyInfo, err := n.notifies.Get(name, fps...)
	if err != nil {
		return nil, err
	}

	// If we have to notify about any of the alerts, we send a notification
	// for the entire batch.
	for i, alert := range alerts {
		if n.hasUpdate(alert, notifyInfo[i], now, repeatInterval) {
			return alerts, nil
		}
	}

	return nil, nil
}

// RetryStage notifies via passed integration with exponential backoff until it
// succeeds. It aborts if the context is canceled or timed out.
type RetryStage struct {
	integration Integration
}

// NewRetryStage returns a new instance of a RetryStage.
func NewRetryStage(i Integration) *RetryStage {
	return &RetryStage{
		integration: i,
	}
}

// Exec implements the Stage interface.
func (r RetryStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	var (
		i    = 0
		b    = backoff.NewExponentialBackOff()
		tick = backoff.NewTicker(b)
	)
	defer tick.Stop()

	for {
		i++
		// Always check the context first to not notify again.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		select {
		case <-tick.C:
			if err := r.integration.Notify(ctx, alerts...); err != nil {
				numFailedNotifications.WithLabelValues(r.integration.name).Inc()
				log.Warnf("Notify attempt %d failed: %s", i, err)
			} else {
				numNotifications.WithLabelValues(r.integration.name).Inc()
				return alerts, nil
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// SetNotifiesStage sets the notification information about passed alerts. The
// passed alerts should have already been sent to the receivers.
type SetNotifiesStage struct {
	notifies provider.Notifies
}

// NewSetNotifiesStage returns a new instance of a SetNotifiesStage.
func NewSetNotifiesStage(notifies provider.Notifies) *SetNotifiesStage {
	return &SetNotifiesStage{
		notifies: notifies,
	}
}

// Exec implements the Stage interface.
func (n SetNotifiesStage) Exec(ctx context.Context, alerts ...*types.Alert) ([]*types.Alert, error) {
	name, ok := Receiver(ctx)
	if !ok {
		return nil, fmt.Errorf("notifier name missing")
	}

	now, ok := Now(ctx)
	if !ok {
		return nil, fmt.Errorf("now time missing")
	}

	var newNotifies []*types.NotificationInfo

	for _, a := range alerts {
		newNotifies = append(newNotifies, &types.NotificationInfo{
			Alert:     a.Fingerprint(),
			Receiver:  name,
			Resolved:  a.Resolved(),
			Timestamp: now,
		})
	}

	return alerts, n.notifies.Set(newNotifies...)
}
