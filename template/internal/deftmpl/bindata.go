// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3b\x79\x6f\xdb\xc6\xf2\xff\xf3\x53\x4c\x59\xfc\xd0\x18\x10\x45\xdb\x39\x50\x1f\xf2\x0f\x8a\x4c\xc7\xc2\x93\x25\x43\x92\x93\x06\x45\x11\x50\xe4\x4a\xda\x84\xe4\xb2\xdc\xa5\x65\xb7\xaf\xdf\xfd\xcd\x2c\xa9\x83\x12\x65\xcb\x46\x6a\xeb\xbd\x3a\xe9\xa1\x1d\xee\xdc\xb3\x33\xb3\xdc\xe5\x9f\x7f\x82\xcf\x86\x3c\x62\x60\x7e\xf9\xe2\x06\x2c\x51\xa1\x1b\xb9\x23\x96\x98\xf0\xd7\x5f\x75\x1a\x5f\x64\x63\x9c\xc8\x22\x1f\x81\xc6\x5a\x94\xab\x6e\x8b\xb0\xf0\x79\xd5\xb9\x51\x2c\x89\xdc\x00\x41\x08\xb1\x7f\xb4\xf5\x3c\xf9\xff\x09\xf3\x18\xbf\x66\x49\x8d\x26\x75\xf3\x41\x86\x93\x53\x2f\x92\x97\xe9\xe0\x2b\xf3\x14\x91\xfd\x95\x50\x7a\xca\x55\xa9\x84\x7f\x83\x12\x57\x71\x3c\x45\xe5\x43\x60\xbf\xcf\x1e\x9a\x43\x9e\xf0\x68\x44\x38\x87\x84\xa3\xb5\x90\xd5\x33\x0d\x45\xd4\x80\x45\x8b\x1c\x7f\x03\x9a\xf4\x21\x11\x69\xdc\x72\x07\x2c\x90\xd5\x9e\x48\x14\xf3\x2f\x5d\x9e\xc8\xea\x47\x37\x48\x19\x31\xfc\x2a\x78\x04\x26\x10\x55\xc8\x58\x8e\x14\xbc\x22\x5a\xd5\x86\x08\x43\x11\x65\xc8\x3b\x39\x6c\x81\xde\x0e\xa2\xbc\x42\x94\x09\x57\xe3\xe2\x64\xb4\x40\x28\xae\x59\x91\x7b\xdb\x0d\x91\x61\x66\xc6\x32\xee\x33\xc1\x77\x66\xbf\xd6\xf8\xc6\x67\xd2\x4b\x78\xac\xb8\x88\xcc\x3b\x6c\xac\xd8\x8d\xca\xfc\xf8\x25\xe0\x52\xe5\x53\x13\x37\x1a\xa1\x64\x38\xc8\xe4\x3a\x34\xe6\xc0\x55\x3b\x91\x55\x2c\x6d\x48\x12\x9f\x46\x35\x98\x29\x90\x0b\x96\x31\xaf\x47\x91\x40\x3f\xa1\x4c\x05\x92\x0b\xe0\xc7\xd1\xed\x89\x34\xf1\xd8\x61\xe6\x4c\x16\xb1\xc4\x55\x22\xc9\xc2\xcf\x28\x31\x54\xc1\x06\x32\x70\xbd\x6f\x55\x1c\xb9\x69\xa0\xaa\x8a\xab\x80\xe5\x56\x50\x2c\x8c\x03\x57\x15\x63\xb1\xba\xce\xe4\x45\x3a\xa9\xa4\x25\x10\x96\x91\x2a\x2e\xb4\x0d\xe9\x0d\xdd\x20\x18\x20\x60\x85\x5e\xa9\xf8\x44\x14\x03\xe7\xbe\x89\x01\x8f\xbe\x6d\x2c\x41\x9c\x30\x0a\x16\x73\xb3\xd9\x0b\xf4\xef\x34\x80\x4e\x1b\x1b\x4a\xc0\x3d\x11\xe1\x9a\xf9\xca\x37\x94\x81\xe6\xa7\x49\xb0\xa9\xc4\x2b\xca\x15\xc2\x64\xcc\x63\x6f\xec\xaa\xb9\x43\x12\x11\x3e\xde\xb9\xcb\xd4\x70\xd5\x4b\x44\xd9\x3c\xf0\x0a\xb2\xc5\xc4\xcd\x4f\xd5\xed\x8c\xde\xea\xea\x7f\x58\x30\xaf\x52\xf4\x02\xce\x22\xf5\x78\x8d\xd7\x51\x9c\xd7\x8d\xc7\x85\xc8\x2a\x5d\x1e\x49\xe5\x46\x1e\x93\x25\x74\x57\xd2\xdd\x1d\x56\x15\xb1\x1c\xb1\x88\xb3\xc7\x3b\xe9\x2e\x62\xab\x1e\xca\xab\xc3\x9a\x64\x58\x5a\x0e\x8c\xa5\x62\x54\xa8\x76\x3b\xb0\x0b\x16\xce\xc9\x80\x90\x01\x75\xda\xbd\xdb\x22\xc5\x92\xa9\x99\x58\x0b\x1a\x95\xf0\xeb\x32\x29\x82\x6b\xe6\x2f\x71\x9c\x82\x37\xe7\x39\xc5\x58\xe1\x6a\x6d\x62\x52\xa9\xab\xc0\xc3\xa3\xa9\xe0\xf5\x6b\xee\x61\xed\x40\xda\x0f\x75\xfb\x72\xbe\x7d\x48\x10\xaf\x32\x7d\x44\x7a\x29\xa8\xc1\x42\x97\x07\x73\xcb\xcc\x3b\xa9\x07\x47\x6e\x91\xd2\x58\x85\x3a\xa1\x1a\xc7\x3f\x9c\x76\x1a\xfd\xcf\x97\x0e\x10\x08\x2e\xaf\xde\xb7\x9a\x0d\x30\x2d\xdb\xfe\xf4\xba\x61\xdb\xa7\xfd\x53\xf8\xe5\xbc\x7f\xd1\x82\xbd\xea\x2e\xf4\xb1\xd0\x4b\x4e\x31\xed\x06\xb6\xed\xb4\x31\x7a\xc7\x4a\xc5\x87\xb6\x3d\x99\x4c\xaa\x93\xd7\x55\x91\x8c\xec\x7e\xd7\xbe\x21\x5a\x7b\x84\x9c\xff\xb4\xd4\x02\x66\xd5\x57\xbe\x79\x82\x9c\x2d\xcb\xe8\xa9\xdb\x80\x81\x8b\xd2\x6a\x26\x3e\x4b\x38\xc5\x0d\x99\x0d\x88\xb4\x44\xda\x23\xec\xb9\xd2\x41\xd5\x13\xa1\x4d\x3a\x8c\xd2\xc8\xd6\xe4\x5c\x2f\xa3\x67\x69\xd5\xac\xa9\x39\x24\x5a\xb0\x3f\x66\x70\xd1\xec\x43\x8b\x7b\x2c\x92\x0c\x5e\xe1\x60\xc7\x30\x1a\x22\xbe\x4d\xf8\x68\x8c\x71\xef\xed\xc0\xfe\xee\xde\x1b\xb8\xc8\x28\x1a\xc6\x25\x4b\x42\x2e\x25\x52\x04\x2e\x61\xcc\x12\x36\xb8\x85\x11\xf2\xc1\x95\x5b\x41\x81\x18\x03\x31\x04\xcc\xf6\xc9\x88\x55\xb0\x75\x45\xa1\x6f\x01\xbb\x57\x89\x08\x62\xa0\x5c\x1e\xd1\x32\x73\xc1\x43\x1e\x06\xce\x54\x63\x24\x23\xc5\x50\x4d\xdc\x24\xd3\xd0\x95\x52\x78\x1c\x25\xf4\xc1\x17\x5e\x1a\x62\xd6\xd4\xf9\x01\x86\x3c\xc0\x8c\xf0\x4a\xa1\xd0\x66\x2f\xc7\x30\x77\x34\x13\x9f\xb9\x81\x81\x79\x82\x9e\x4d\x1f\xe9\x26\x54\xa4\x0a\x12\x26\x55\xc2\xb5\x15\x2a\xc0\x23\x2f\x48\x7d\x92\x61\xfa\x38\xe0\x21\xcf\x39\x10\xba\x56\x5c\x1a\x48\x14\x9b\x9a\x8a\x96\xb3\x02\xa1\xf0\xf9\x90\xfe\xcf\xb4\x5a\x71\x3a\xc0\x95\x3c\xae\x80\xcf\x89\xf4\x20\x55\x08\x94\x04\xd4\x76\xac\x90\x1e\xb6\x48\x40\xb2\x20\x30\x90\x02\x47\xb9\xb5\xae\x73\xe9\xf4\x1c\x12\x3d\x26\x83\xaa\xdc\x44\x92\x20\x93\x31\x7a\xb5\xa0\x09\x97\xc6\x30\x4d\x22\x64\xc9\x34\x8e\x2f\xd0\x64\x9a\x23\x45\x33\x41\x68\xfa\x50\x04\x81\x98\x90\x6a\xd8\x09\xf8\x3c\xef\x3b\xb5\x93\xdd\x01\xf5\xde\xde\xcc\xaf\x98\x73\x51\xd4\x4c\x04\x72\x40\x3c\xf7\x6a\xfe\x48\x8e\xb1\x05\x83\x01\xcb\x0d\x86\x7c\xd1\xbc\xee\x82\x3a\x09\xb1\xa7\xd2\xa3\xb8\x1b\x40\x8c\xa9\x9b\xf8\x2d\xab\x59\x45\xfe\xe7\x0e\xf4\x3a\x67\xfd\x4f\xf5\xae\x03\xcd\x1e\x5c\x76\x3b\x1f\x9b\xa7\xce\x29\x98\xf5\x1e\x8e\xcd\x0a\x7c\x6a\xf6\xcf\x3b\x57\x7d\xc0\x19\xdd\x7a\xbb\xff\x19\x3a\x67\x50\x6f\x7f\x86\x7f\x35\xdb\xa7\x15\x70\x7e\xb9\xec\x3a\xbd\x1e\x74\xba\x46\xf3\xe2\xb2\xd5\x74\x10\xd6\x6c\x37\x5a\x57\xa7\xcd\xf6\x07\x78\x8f\x78\xed\x0e\x86\x70\x13\x63\x17\x89\xf6\x3b\x40\x0c\x73\x52\x4d\xa7\x47\xc4\x2e\x9c\x6e\xe3\x1c\x87\xf5\xf7\xcd\x56\xb3\xff\xb9\x62\x9c\x35\xfb\x6d\xa2\x79\xd6\xe9\x42\x1d\x2e\xeb\xdd\x7e\xb3\x71\xd5\xaa\x77\x71\x61\x77\x2f\x3b\x3d\x07\xd9\x9f\x22\xd9\x76\xb3\x7d\xd6\x45\x2e\xce\x85\xd3\xee\x57\x91\x2b\xc2\xc0\xf9\x88\x03\xe8\x9d\xd7\x5b\x2d\x62\x65\xd4\xaf\x50\xfa\x2e\xc9\x07\x8d\xce\xe5\xe7\x6e\xf3\xc3\x79\x1f\xce\x3b\xad\x53\x07\x81\xef\x1d\x94\xac\xfe\xbe\xe5\x64\xac\x50\xa9\x46\xab\xde\xbc\xa8\xc0\x69\xfd\xa2\xfe\xc1\xd1\x58\x1d\xa4\xd2\x35\x68\x5a\x26\x1d\x7c\x3a\x77\x08\x44\xfc\xea\xf8\x4f\xa3\xdf\xec\xb4\x49\x8d\x46\xa7\xdd\xef\xe2\xb0\x82\x5a\x76\xfb\x33\xd4\x4f\xcd\x9e\x53\x81\x7a\xb7\xd9\x23\x83\x9c\x75\x3b\x17\x15\x83\xcc\x89\x18\x1d\x4d\x04\xf1\xda\x4e\x46\x85\x4c\x0d\x05\x8f\xe0\x14\x1a\x5f\xf5\x9c\x19\x41\x38\x75\xea\x2d\xa4\xd5\x23\x64\x52\x71\x3a\xb9\x6a\x58\x16\x66\x24\x9d\x02\x6f\xc2\x20\x92\xb5\x92\xc4\xb6\x77\x70\x70\x90\xe5\x33\x73\xb3\x49\x92\x92\x5b\xcd\x1c\x8a\x48\x59\x43\x37\xe4\xc1\xed\x21\xfc\x74\xce\xb0\x32\x62\x24\xba\xd0\x66\x29\xfb\xa9\x02\x33\x00\xaa\x9a\x60\xc8\x61\xf8\x63\x72\xb3\x70\xe3\xc1\x87\x47\x30\x10\x37\x96\xe4\x7f\x50\xc9\xc7\xdf\x09\x26\x48\x0b\x41\x47\xa0\x89\xe2\x03\xdc\x2d\xed\xbd\x89\x11\x10\x62\x62\xe2\xd1\x21\xec\x1e\x51\x6e\x1d\x33\xd7\x7f\x4e\xfe\x21\x53\x2e\xd0\xc6\xa9\x86\x45\x91\x4d\x68\x15\x99\xb4\x7a\x15\x26\xbd\x9a\x39\xe1\xbe\x1a\xd7\x7c\x86\xf5\x92\x59\x7a\xf0\x7c\xc6\x02\x7b\x2a\x2e\x39\xd3\x62\xbf\xa7\xfc\xba\x66\x36\x32\x51\xad\xfe\x6d\xcc\x16\x04\xa7\x8e\xc7\x26\xe7\x1e\xe9\x4a\x20\x99\xaa\x5d\xf5\xcf\xac\x9f\x9f\x59\x7c\xbd\x4b\x7b\x3e\x77\xdf\xd5\x8b\x1c\xdb\x5a\xb8\x13\xc3\x38\xb6\x29\x28\xe9\xc7\x40\xf8\xb7\xc0\x11\x45\x62\xce\x45\x89\x4d\x3d\x50\xb7\xf4\x3b\x5f\x51\xd2\x1b\x63\x55\xd7\x2b\xca\xa1\xea\x7e\x31\x6d\xde\x9e\x54\x49\x6b\xc2\x06\xdf\x38\x32\xd2\x0f\x42\x21\xb0\xa6\x10\x52\x56\x1b\xb8\x2b\x99\x3f\x9f\x44\xb1\xa1\xb1\x2d\xd7\xff\x9a\x4a\x75\x88\x15\x27\x62\x47\xd8\x4a\x50\x65\x42\x92\xbb\xbb\xff\x77\x84\x45\x39\x62\xd6\x0c\x54\x7d\xc7\xc2\x23\xd0\x2b\x20\x9b\x00\x3f\xf0\x90\x16\x0b\x72\x40\x39\x71\x5f\x3b\x4a\x44\x1a\xf9\x96\x27\x02\x91\x1c\xc2\x8f\xc3\x77\xf4\x77\xd1\xfc\x10\xbb\xbe\xaf\xa5\xa2\x68\x18\x8c\xf4\xcc\x9a\x99\xcf\x34\xc9\xde\xca\x1d\x3c\x75\x78\x2c\xa8\xb4\xa1\x1e\xa5\xb2\x03\x1c\xab\xe4\x19\xf3\x18\x00\x49\xf0\xc4\x99\xf4\x1a\xf7\x06\x48\x24\xb0\x30\xc4\x46\x28\x89\x12\x71\xd1\x50\xd7\xfa\x01\x66\x23\x11\x9b\x27\xb8\xc0\xfc\xb9\xa0\x59\x66\x35\xdf\xed\xee\x3e\xf1\x52\x29\x15\x1a\xbb\x48\xcc\x0a\xc8\x76\x10\x08\xef\x5b\x21\xb6\x43\xf7\xc6\xca\x83\x04\x85\x8d\x6f\x0a\x0f\xbd\x80\xb9\x09\x31\x54\xe3\x02\x7c\xdd\x42\x99\x19\x07\xdc\x54\x89\xa5\x25\x51\xb0\x96\x36\x14\x9a\xca\xe7\xd7\x4f\x1d\x56\x45\x7d\x97\x8d\x73\xb7\x12\x53\xb9\xc9\xc9\x7a\x31\xe7\x7e\x26\x4b\x60\x79\xc2\x6e\x3c\x9f\x5d\x33\x77\xb3\xb1\x8c\x5d\x6f\x3a\x7e\x52\x45\xf3\x87\x89\xeb\xf3\x54\x1e\xc2\x6b\x0d\x2b\x49\x00\xc3\x61\x21\x8b\x65\x68\x48\x04\x43\x41\x8a\x80\xfb\xf0\x23\x3b\xa0\xbf\xc5\xc4\x30\x1c\x2e\xd8\x62\x1b\xb2\xc3\x5c\x92\xa7\xcb\x12\xef\xd6\x2e\xb8\x82\x75\x35\xca\x24\x2f\x35\x6f\x77\xd1\xc8\xba\x44\xe5\xf3\x71\x43\xa7\x58\x52\xe6\x2f\xfd\xef\xae\x76\xca\xaa\xdf\x9c\x77\x6f\xf7\xf7\x1b\xe5\x05\x68\x9f\xe2\xda\x84\x7c\xbd\x65\x0c\x16\xbd\x97\xe1\x96\xaf\xc8\xe9\x9f\xf9\x61\xcf\xec\x94\x07\xf4\xdb\x92\xd2\x57\x56\x3b\xb0\x87\x13\xe4\xec\x85\x07\xea\x9c\xc0\xfc\x40\x62\xcd\x81\x10\xbd\xf7\x00\x58\xe5\x9b\x1f\x4f\xd4\x0a\x87\x13\x2b\xd3\xf2\x57\x2b\x05\xe7\xcf\x72\xf0\x6c\x9c\xbc\x84\xe9\x26\xc5\x6c\x1e\x3c\x7b\x59\xf0\xdc\x15\x1b\x5b\x9f\xfb\xd6\x9a\x7d\xbb\x82\x60\xdb\x43\x01\x73\xcf\x34\x97\xdc\x15\x0e\xb9\x1a\xb8\x71\x4b\xd8\xb0\x66\x6e\xf2\xd6\xf6\x89\xe3\x61\x9a\x34\xcf\xce\xce\xf2\xe4\xeb\x33\x4f\x24\xfa\x9d\xdc\x74\x7b\x50\xd8\x10\xec\xd3\x76\xa0\x90\xb7\x07\x22\xf0\xcb\x13\xb7\x97\x26\x92\xa8\xc7\x82\x67\x80\x59\x43\xc1\x23\x4d\x34\xef\x2b\x96\x12\xfc\x5b\x12\x4c\xd3\xd3\x2f\x51\x31\x61\x86\x48\xd3\x8d\xb9\x42\xfa\x7f\xb0\xd2\xa4\xff\xfa\xcd\xcf\xcc\x77\x4b\xea\xf5\xca\x8c\x1c\xac\xad\x7c\x98\x15\xf2\x19\x70\xd6\xbd\x61\x79\xc9\xdc\x7b\xf2\x91\xb3\x09\xbd\x7f\xbb\xf7\xf5\xf8\xb1\xed\x96\xc6\xf0\x52\xe2\x2d\x4f\xbf\xd9\x9f\xfb\xce\x58\x4a\x8a\xc2\xcb\x92\xfd\x7b\x96\xac\x54\x89\x88\x46\xcf\x67\xda\x5f\xd7\x5f\x29\xf9\x2d\x3f\x60\x3b\xb6\x33\x21\xbf\x43\xd4\x95\x34\x0c\xf9\x93\xe9\xbd\x89\xe5\x93\xba\x97\x38\xfc\x67\xc4\x61\xd6\x9a\xce\x42\xed\x78\xf0\x7c\x6e\xa6\xf7\x88\x65\x36\xba\xe7\xc2\xd0\xfa\x5b\x3d\xcf\xac\xcc\xfa\x75\x97\x6b\x55\xa8\x05\xf3\xb3\xfa\xac\x12\x3c\x7b\x64\x2c\x48\xb4\x2d\xe1\x71\xaf\x45\xef\xbd\x05\xf6\x5f\x1a\x2c\x8b\x1d\xe6\xf2\xb5\xb4\x67\x6a\x28\xa7\xed\xd6\x4a\x4f\x89\x5d\x1b\x4b\xa8\xfb\x2b\x86\x53\x76\xb1\x8e\x9a\xa8\xed\xcb\x31\x8f\xab\xa6\x1b\xb6\x77\x8b\x57\x5a\x4a\xdd\xfb\xd2\x15\x6e\x4d\x35\xde\xba\xc8\x44\x99\xc6\x5b\x28\xd3\xd6\xd9\xe9\x21\x2b\xf8\xae\x8e\xf8\x65\x61\xfd\x6f\xb6\xb9\x8b\xdb\xad\xd9\xd5\xc0\xf9\x86\x6b\x0a\x7a\x86\x2d\xd7\xe2\x45\xc5\x97\x68\xfc\x67\x44\xe3\xcb\xa6\xeb\x65\xd3\xf5\xb2\xe9\xda\xf6\x60\x79\xd9\x74\x6d\x4d\xcb\xb6\xce\x51\x38\x9b\xce\xe3\x4e\x1e\x70\x14\x3a\x43\x99\x43\x9e\xfc\x26\x46\xe1\x6a\xd2\xc2\x4d\x93\xb9\xa3\x0f\x0e\x0e\xee\x3a\xe0\x2e\x9e\xec\xae\x1e\x49\x6e\x47\xd3\xb0\x4d\xed\xcb\x53\xb6\x2e\xfb\x6b\x5b\x97\xd2\x43\xb4\xfb\x5c\xbe\xd0\xdb\x2c\xdd\x6b\x28\xde\xc2\x5a\x4c\x57\xc5\x0f\x67\x9f\x2e\x20\xf6\x17\xb3\x95\xd6\x68\xe3\x54\x85\x3a\xc1\xe0\x76\xb3\x73\xb8\xd5\xdc\xb1\x72\xdf\x61\x39\x33\x1c\xdb\xb8\xcc\x4f\xb2\xff\x1a\xc5\x34\xb1\x6d\x6d\xed\x9a\xeb\x75\x99\x8a\xf3\xfc\x75\x6c\xd3\x2d\x56\x82\xd0\x75\xe0\x13\xc3\x28\xff\x32\x37\x4e\xe5\x58\x20\xc7\xef\xf0\x61\xea\x0a\xa9\xe2\x07\x4d\x7f\xc7\x67\x67\xdf\xe7\xab\xb3\xcd\x3f\x3a\xfb\x7e\xdf\x9c\x2d\xf0\xdc\xc0\x92\xf3\xaf\x4b\x1f\xf0\xdd\xd7\x7f\x02\x00\x00\xff\xff\x07\xde\x61\xfd\x76\x3f\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 16246, mode: os.FileMode(420), modTime: time.Unix(1470746752, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/default.tmpl": templateDefaultTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
