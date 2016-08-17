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
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/prometheus/common/model"
	"golang.org/x/net/context"

	"github.com/prometheus/alertmanager/nflog/nflogpb"
	"github.com/prometheus/alertmanager/types"
	"github.com/satori/go.uuid"
)

type failStage struct{}

func (s failStage) Exec(ctx context.Context, as ...*types.Alert) (context.Context, []*types.Alert, error) {
	return ctx, nil, fmt.Errorf("some error")
}

func TestDedupStageHasUpdate(t *testing.T) {
	var (
		n        = &DedupStage{}
		now      = time.Now()
		interval = 100 * time.Minute
	)
	cases := []struct {
		inAlert      *types.Alert
		inNotifyInfo *types.NotificationInfo
		result       bool
	}{
		// A new alert about which there's no previous notification information.
		{
			inAlert: &types.Alert{
				Alert: model.Alert{
					Labels:   model.LabelSet{"alertname": "a"},
					StartsAt: now.Add(-10 * time.Minute),
				},
			},
			inNotifyInfo: nil,
			result:       true,
		},
		// A new alert about which there's no previous notification information.
		// It is already resolved, so there's no use in sending a notification.
		{
			inAlert: &types.Alert{
				Alert: model.Alert{
					Labels:   model.LabelSet{"alertname": "a"},
					StartsAt: now.Add(-10 * time.Minute),
					EndsAt:   now,
				},
			},
			inNotifyInfo: nil,
			result:       false,
		},
		// An alert that has been firing is now resolved for the first time.
		{
			inAlert: &types.Alert{
				Alert: model.Alert{
					Labels:   model.LabelSet{"alertname": "a"},
					StartsAt: now.Add(-10 * time.Minute),
					EndsAt:   now,
				},
			},
			inNotifyInfo: &types.NotificationInfo{
				Alert:     model.LabelSet{"alertname": "a"}.Fingerprint(),
				Resolved:  false,
				Timestamp: now.Add(-time.Minute),
			},
			result: true,
		},
		// A resolved alert for which we have already sent a resolved notification.
		{
			inAlert: &types.Alert{
				Alert: model.Alert{
					Labels:   model.LabelSet{"alertname": "a"},
					StartsAt: now.Add(-10 * time.Minute),
					EndsAt:   now,
				},
			},
			inNotifyInfo: &types.NotificationInfo{
				Alert:     model.LabelSet{"alertname": "a"}.Fingerprint(),
				Resolved:  true,
				Timestamp: now.Add(-time.Minute),
			},
			result: false,
		},
		// An alert that was resolved last time but is now firing again.
		{
			inAlert: &types.Alert{
				Alert: model.Alert{
					Labels:   model.LabelSet{"alertname": "a"},
					StartsAt: now.Add(-3 * time.Minute),
				},
			},
			inNotifyInfo: &types.NotificationInfo{
				Alert:     model.LabelSet{"alertname": "a"}.Fingerprint(),
				Resolved:  true,
				Timestamp: now.Add(-4 * time.Minute),
			},
			result: true,
		},
		// A firing alert about which we already notified. The last notification
		// is less than the repeat interval ago.
		{
			inAlert: &types.Alert{
				Alert: model.Alert{
					Labels:   model.LabelSet{"alertname": "a"},
					StartsAt: now.Add(-10 * time.Minute),
				},
			},
			inNotifyInfo: &types.NotificationInfo{
				Alert:     model.LabelSet{"alertname": "a"}.Fingerprint(),
				Resolved:  false,
				Timestamp: now.Add(-15 * time.Minute),
			},
			result: false,
		},
		// A firing alert about which we already notified. The last notification
		// is more than the repeat interval ago.
		{
			inAlert: &types.Alert{
				Alert: model.Alert{
					Labels:   model.LabelSet{"alertname": "a"},
					StartsAt: now.Add(-10 * time.Minute),
				},
			},
			inNotifyInfo: &types.NotificationInfo{
				Alert:     model.LabelSet{"alertname": "a"}.Fingerprint(),
				Resolved:  false,
				Timestamp: now.Add(-115 * time.Minute),
			},
			result: true,
		},
	}

	for i, c := range cases {
		if n.hasUpdate(c.inAlert, c.inNotifyInfo, now, interval) != c.result {
			t.Errorf("unexpected hasUpdates result for case %d", i)
		}
	}
}

func TestMultiStage(t *testing.T) {
	var (
		alerts1 = []*types.Alert{{}}
		alerts2 = []*types.Alert{{}, {}}
		alerts3 = []*types.Alert{{}, {}, {}}
	)

	stage := MultiStage{
		StageFunc(func(ctx context.Context, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
			if !reflect.DeepEqual(alerts, alerts1) {
				t.Fatal("Input not equal to input of MultiStage")
			}
			return ctx, alerts2, nil
		}),
		StageFunc(func(ctx context.Context, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
			if !reflect.DeepEqual(alerts, alerts2) {
				t.Fatal("Input not equal to output of previous stage")
			}
			return ctx, alerts3, nil
		}),
	}

	_, alerts, err := stage.Exec(context.Background(), alerts1...)
	if err != nil {
		t.Fatalf("Exec failed: %s", err)
	}

	if !reflect.DeepEqual(alerts, alerts3) {
		t.Fatal("Output of MultiStage is not equal to the output of the last stage")
	}
}

func TestMultiStageFailure(t *testing.T) {
	var (
		ctx   = context.Background()
		s1    = failStage{}
		stage = MultiStage{s1}
	)

	_, _, err := stage.Exec(ctx, nil)
	if err.Error() != "some error" {
		t.Fatal("Errors were not propagated correctly by MultiStage")
	}
}

func TestRoutingStage(t *testing.T) {
	var (
		alerts1 = []*types.Alert{{}}
		alerts2 = []*types.Alert{{}, {}}
	)

	stage := RoutingStage{
		"name": StageFunc(func(ctx context.Context, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
			if !reflect.DeepEqual(alerts, alerts1) {
				t.Fatal("Input not equal to input of RoutingStage")
			}
			return ctx, alerts2, nil
		}),
		"not": failStage{},
	}

	ctx := WithReceiverName(context.Background(), "name")

	_, alerts, err := stage.Exec(ctx, alerts1...)
	if err != nil {
		t.Fatalf("Exec failed: %s", err)
	}

	if !reflect.DeepEqual(alerts, alerts2) {
		t.Fatal("Output of RoutingStage is not equal to the output of the inner stage")
	}
}

func TestSetNotifiesStage(t *testing.T) {
	var (
		notifies = newTestInfos()
		recv     = &nflogpb.Receiver{GroupName: "name"}
		stage    = NewSetNotifiesStage(notifies, recv)
		ctx      = context.Background()
	)
	now := time.Now()

	ctx = WithRepeatInterval(ctx, time.Duration(100*time.Minute))
	ctx = WithNow(ctx, now)

	alerts := []*types.Alert{
		{
			Alert: model.Alert{
				Labels: model.LabelSet{"alertname": "0"},
			},
		}, {
			Alert: model.Alert{
				Labels: model.LabelSet{"alertname": "1"},
				EndsAt: now.Add(-5 * time.Minute),
			},
		},
	}

	// Set an initial NotifyInfo to ensure that on notification failure
	// nothing changes.
	nsBefore := []*types.NotificationInfo{
		nil,
		{
			Alert:     alerts[1].Fingerprint(),
			Receiver:  recv.String(),
			Resolved:  false,
			Timestamp: now.Add(-10 * time.Minute),
		},
	}

	if err := notifies.Set(nsBefore...); err != nil {
		t.Fatalf("Setting notifies failed: %s", err)
	}

	_, _, err := stage.Exec(ctx, alerts...)
	if err != nil {
		t.Fatalf("Exec failed: %s", err)
	}

	nsCur, err := notifies.Get(recv.String(), alerts[0].Fingerprint(), alerts[1].Fingerprint())
	if err != nil {
		t.Fatalf("Error getting notifies: %s", err)
	}

	nsAfter := []*types.NotificationInfo{
		{
			Alert:     alerts[0].Fingerprint(),
			Receiver:  recv.String(),
			Resolved:  false,
			Timestamp: now,
		},
		{
			Alert:     alerts[1].Fingerprint(),
			Receiver:  recv.String(),
			Resolved:  true,
			Timestamp: now,
		},
	}

	for i, after := range nsAfter {
		cur := nsCur[i]

		// Hack correct timestamps back in if they are sane.
		if cur != nil && after.Timestamp.IsZero() {
			if cur.Timestamp.Before(now) {
				t.Fatalf("Wrong timestamp for notify %v", cur)
			}
			after.Timestamp = cur.Timestamp
		}

		if !reflect.DeepEqual(after, cur) {
			t.Errorf("Unexpected notifies, expected: %v, got: %v", after, cur)
		}
	}
}

func TestSilenceStage(t *testing.T) {
	// Mute all label sets that have a "mute" key.
	muter := types.MuteFunc(func(lset model.LabelSet) bool {
		_, ok := lset["mute"]
		return ok
	})

	marker := types.NewMarker()
	silencer := NewSilenceStage(muter, marker)

	in := []model.LabelSet{
		{},
		{"test": "set"},
		{"mute": "me"},
		{"foo": "bar", "test": "set"},
		{"foo": "bar", "mute": "me"},
		{},
		{"not": "muted"},
	}
	out := []model.LabelSet{
		{},
		{"test": "set"},
		{"foo": "bar", "test": "set"},
		{},
		{"not": "muted"},
	}

	var inAlerts []*types.Alert
	for _, lset := range in {
		inAlerts = append(inAlerts, &types.Alert{
			Alert: model.Alert{Labels: lset},
		})
	}

	// Set the second alert als previously silenced. It is expected to have
	// the WasSilenced flag set to true afterwards.
	marker.SetSilenced(inAlerts[1].Fingerprint(), uuid.NewV4())

	_, alerts, err := silencer.Exec(nil, inAlerts...)
	if err != nil {
		t.Fatalf("Exec failed: %s", err)
	}

	var got []model.LabelSet
	for i, a := range alerts {
		got = append(got, a.Labels)
		if a.WasSilenced != (i == 1) {
			t.Errorf("Expected WasSilenced to be %v for %d, was %v", i == 1, i, a.WasSilenced)
		}
	}

	if !reflect.DeepEqual(got, out) {
		t.Fatalf("Muting failed, expected: %v\ngot %v", out, got)
	}
}

func TestInhibitStage(t *testing.T) {
	// Mute all label sets that have a "mute" key.
	muter := types.MuteFunc(func(lset model.LabelSet) bool {
		_, ok := lset["mute"]
		return ok
	})

	marker := types.NewMarker()
	inhibitor := NewInhibitStage(muter, marker)

	in := []model.LabelSet{
		{},
		{"test": "set"},
		{"mute": "me"},
		{"foo": "bar", "test": "set"},
		{"foo": "bar", "mute": "me"},
		{},
		{"not": "muted"},
	}
	out := []model.LabelSet{
		{},
		{"test": "set"},
		{"foo": "bar", "test": "set"},
		{},
		{"not": "muted"},
	}

	var inAlerts []*types.Alert
	for _, lset := range in {
		inAlerts = append(inAlerts, &types.Alert{
			Alert: model.Alert{Labels: lset},
		})
	}

	// Set the second alert as previously inhibited. It is expected to have
	// the WasInhibited flag set to true afterwards.
	marker.SetInhibited(inAlerts[1].Fingerprint(), true)

	_, alerts, err := inhibitor.Exec(nil, inAlerts...)
	if err != nil {
		t.Fatalf("Exec failed: %s", err)
	}

	var got []model.LabelSet
	for i, a := range alerts {
		got = append(got, a.Labels)
		if a.WasInhibited != (i == 1) {
			t.Errorf("Expected WasInhibited to be %v for %d, was %v", i == 1, i, a.WasInhibited)
		}
	}

	if !reflect.DeepEqual(got, out) {
		t.Fatalf("Muting failed, expected: %v\ngot %v", out, got)
	}
}

type testInfos struct {
	mtx sync.RWMutex
	m   map[string]map[model.Fingerprint]*types.NotificationInfo
}

func newTestInfos() *testInfos {
	return &testInfos{
		m: map[string]map[model.Fingerprint]*types.NotificationInfo{},
	}
}

// Set implements the Notifies interface.
func (n *testInfos) Set(ns ...*types.NotificationInfo) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	for _, v := range ns {

		if v == nil {
			continue
		}
		am, ok := n.m[v.Receiver]
		if !ok {
			am = map[model.Fingerprint]*types.NotificationInfo{}
			n.m[v.Receiver] = am
		}
		am[v.Alert] = v
	}
	return nil
}

// Get implements the Notifies interface.
func (n *testInfos) Get(dest string, fps ...model.Fingerprint) ([]*types.NotificationInfo, error) {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	res := make([]*types.NotificationInfo, len(fps))

	ns, ok := n.m[dest]
	if !ok {
		return res, nil
	}

	for i, fp := range fps {
		res[i] = ns[fp]
	}

	return res, nil
}
