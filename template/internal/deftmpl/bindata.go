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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3b\x79\x6f\xdb\xc6\xf2\xff\xf3\x53\x4c\x59\xfc\xd0\x18\x10\x45\xdb\x39\x50\x1f\xf2\x0f\x8a\x4c\xc7\xc2\x93\x25\x43\x92\x93\x06\x45\x11\x50\xe4\x4a\xda\x84\xe4\xb2\xdc\xa5\x65\xb7\xaf\xdf\xfd\xcd\x2c\xa9\x83\x12\x65\xcb\x46\x6a\xeb\xbd\x3a\xe9\xa1\x1d\xee\xdc\xb3\x33\xb3\xdc\xe5\x9f\x7f\x82\xcf\x86\x3c\x62\x60\x7e\xf9\xe2\x06\x2c\x51\xa1\x1b\xb9\x23\x96\x98\xf0\xd7\x5f\x75\x1a\x5f\x64\x63\x9c\xc8\x22\x1f\x81\xc6\x5a\x94\xab\x6e\x8b\xb0\xf0\x79\xd5\xb9\x51\x2c\x89\xdc\x00\x41\x08\xb1\x7f\xb4\xf5\x3c\xf9\xff\x09\xf3\x18\xbf\x66\x49\x8d\x26\x75\xf3\x41\x86\x93\x53\x2f\x92\x97\xe9\xe0\x2b\xf3\x14\x91\xfd\x95\x50\x7a\xca\x55\xa9\x84\x7f\x83\x12\x57\x71\x3c\x45\xe5\x43\x60\xbf\xcf\x1e\x9a\x43\x9e\xf0\x68\x44\x38\x87\x84\xa3\xb5\x90\xd5\x33\x0d\x45\xd4\x80\x45\x8b\x1c\x7f\x03\x9a\xf4\x21\x11\x69\xdc\x72\x07\x2c\x90\xd5\x9e\x48\x14\xf3\x2f\x5d\x9e\xc8\xea\x47\x37\x48\x19\x31\xfc\x2a\x78\x04\x26\x10\x55\xc8\x58\x8e\x14\xbc\x22\x5a\xd5\x86\x08\x43\x11\x65\xc8\x3b\x39\x6c\x81\xde\x0e\xa2\xbc\x42\x94\x09\x57\xe3\xe2\x64\xb4\x40\x28\xae\x59\x91\x7b\xdb\x0d\x91\x61\x66\xc6\x32\xee\x33\xc1\x77\x66\xbf\xd6\xf8\xc6\x67\xd2\x4b\x78\xac\xb8\x88\xcc\x3b\x6c\xac\xd8\x8d\xca\xfc\xf8\x25\xe0\x52\xe5\x53\x13\x37\x1a\xa1\x64\x38\xc8\xe4\x3a\x34\xe6\xc0\x55\x3b\x91\x55\x2c\x6d\x48\x12\x9f\x46\x35\x98\x29\x90\x0b\x96\x31\xaf\x47\x91\x40\x3f\xa1\x4c\x05\x92\x0b\xe0\xc7\xd1\xed\x89\x34\xf1\xd8\x61\xe6\x4c\x16\xb1\xc4\x55\x22\xc9\xc2\xcf\x28\x31\x54\xc1\x06\x32\x70\xbd\x6f\x55\x1c\xb9\x69\xa0\xaa\x8a\xab\x80\xe5\x56\x50\x2c\x8c\x03\x57\x15\x63\xb1\xba\xce\xe4\x45\x3a\xa9\xa4\x25\x10\x96\x91\x2a\x2e\xb4\x0d\xe9\x0d\xdd\x20\x18\x20\x60\x85\x5e\xa9\xf8\x44\x14\x03\xe7\xbe\x89\x01\x8f\xbe\x6d\x2c\x41\x9c\x30\x0a\x16\x73\xb3\xd9\x0b\xf4\xef\x34\x80\x4e\x1b\x1b\x4a\xc0\x3d\x11\xe1\x9a\xf9\xca\x37\x95\x61\x45\xdc\x82\xe3\xc7\x3c\xf6\xc6\xae\x9a\x9b\x38\x11\xe1\xe3\xdd\xb5\x4c\x0d\xd7\xb1\x44\x94\xcd\x43\xa9\x20\x5b\x4c\xdc\xfc\x54\xdd\xce\xe8\xad\xae\xe7\x87\x85\xe7\x2a\x45\x2f\xe0\x2c\x52\x8f\xd7\x78\x1d\xc5\x79\x25\x78\x9c\xd3\x57\xe9\xf2\x48\x2a\x37\xf2\x98\x2c\xa1\xbb\x92\xc0\xee\xb0\xaa\x88\xe5\x88\x45\x9c\x3d\xde\x49\x77\x11\x5b\xf5\x50\x9e\xef\xd7\xa4\xb7\xd2\x04\x6f\x2c\x95\x97\x42\xfd\xda\x81\x5d\xb0\x70\x4e\x06\x84\x0c\xa8\x13\xe9\xdd\x16\x29\x16\x41\xcd\xc4\x5a\xd0\xa8\x84\x5f\x97\x49\x11\x5c\x33\x7f\x89\xe3\x14\xbc\x39\xcf\x29\xc6\x0a\x57\x6b\x13\x93\x4a\x9d\xd7\x1f\x1e\x4d\x05\xaf\x5f\x73\x0f\xab\x01\xd2\x7e\xa8\xdb\x97\x33\xe8\x43\x82\x78\x95\xe9\x23\xd2\x4b\x41\x0d\x16\xba\x3c\x98\x5b\x66\xde\x1b\x3d\x38\x72\x8b\x94\xc6\x2a\x0c\x74\xe4\x1d\xff\x70\xda\x69\xf4\x3f\x5f\x3a\x40\x20\xb8\xbc\x7a\xdf\x6a\x36\xc0\xb4\x6c\xfb\xd3\xeb\x86\x6d\x9f\xf6\x4f\xe1\x97\xf3\xfe\x45\x0b\xf6\xaa\xbb\xd0\xc7\xd2\x2d\x39\xc5\xb4\x1b\xd8\xb6\xd3\xc6\xe8\x1d\x2b\x15\x1f\xda\xf6\x64\x32\xa9\x4e\x5e\x57\x45\x32\xb2\xfb\x5d\xfb\x86\x68\xed\x11\x72\xfe\xd3\x52\x0b\x98\x55\x5f\xf9\xe6\x09\x72\xb6\x2c\xa3\xa7\x6e\x03\x06\x2e\x4a\xab\x99\xf8\x2c\xe1\x14\x37\x64\x36\x20\xd2\x12\x69\x8f\xb0\x8b\x4a\x07\x55\x4f\x84\x36\xe9\x30\x4a\x23\x5b\x93\x73\xbd\x8c\x9e\xa5\x55\xb3\xa6\xe6\x90\x68\xc1\xfe\x98\xc1\x45\xb3\x0f\x2d\xee\xb1\x48\x32\x78\x85\x83\x1d\xc3\x68\x88\xf8\x36\xe1\xa3\x31\xc6\xbd\xb7\x03\xfb\xbb\x7b\x6f\xe0\x22\xa3\x68\x18\x97\x2c\x09\xb9\x94\x48\x11\xb8\x84\x31\x4b\xd8\xe0\x16\x46\xc8\x07\x57\x6e\x05\x05\x62\x0c\xc4\x10\x30\xdb\x27\x23\x56\xc1\x66\x14\x85\xbe\x05\xec\x47\x25\x22\x88\x81\x72\x79\x44\xcb\xcc\x05\x0f\x79\x18\x38\x53\x8d\x91\x8c\x14\x43\x35\x71\x93\x4c\x43\x57\x4a\xe1\x71\x94\xd0\x07\x5f\x78\x69\x88\x59\x53\xe7\x07\x18\xf2\x00\x33\xc2\x2b\x85\x42\x9b\xbd\x1c\xc3\xdc\xd1\x4c\x7c\xe6\x06\x06\xe6\x09\x7a\x36\x7d\xa4\xdb\x4a\x91\x2a\x48\x98\x54\x09\xd7\x56\xa8\x00\x8f\xbc\x20\xf5\x49\x86\xe9\xe3\x80\x87\x3c\xe7\x40\xe8\x5a\x71\x69\x20\x51\x6c\x53\x2a\x5a\xce\x0a\x84\xc2\xe7\x43\xfa\x3f\xd3\x6a\xc5\xe9\x00\x57\xf2\xb8\x02\x3e\x27\xd2\x83\x54\x21\x50\x12\x50\xdb\xb1\x42\x7a\xd8\x22\x01\xc9\x82\xc0\x40\x0a\x1c\xe5\xd6\xba\xce\xa5\xd3\x73\x48\xf4\x98\x0c\xaa\x72\x13\x49\x82\x4c\xc6\xe8\xd5\x82\x26\x5c\x1a\xc3\x34\x89\x90\x25\xd3\x38\xbe\x40\x93\x69\x8e\x14\xcd\x04\xa1\xe9\x43\x11\x04\x62\x42\xaa\x61\x2f\xe0\xf3\xbc\x93\xd4\x4e\x76\x07\xd4\x4d\x7b\x33\xbf\x62\xce\x45\x51\x33\x11\xc8\x01\xf1\xdc\xab\xf9\x23\x39\xc6\xa6\x0a\x06\x2c\x37\x18\xf2\x45\xf3\xba\x0b\xea\x24\xc4\x9e\x4a\x8f\xe2\x6e\x00\x31\xa6\x6e\xe2\xb7\xac\x66\x15\xf9\x9f\x3b\xd0\xeb\x9c\xf5\x3f\xd5\xbb\x0e\x34\x7b\x70\xd9\xed\x7c\x6c\x9e\x3a\xa7\x60\xd6\x7b\x38\x36\x2b\xf0\xa9\xd9\x3f\xef\x5c\xf5\x01\x67\x74\xeb\xed\xfe\x67\xe8\x9c\x41\xbd\xfd\x19\xfe\xd5\x6c\x9f\x56\xc0\xf9\xe5\xb2\xeb\xf4\x7a\xd0\xe9\x1a\xcd\x8b\xcb\x56\xd3\x41\x58\xb3\xdd\x68\x5d\x9d\x36\xdb\x1f\xe0\x3d\xe2\xb5\x3b\x18\xc2\x4d\x8c\x5d\x24\xda\xef\x00\x31\xcc\x49\x35\x9d\x1e\x11\xbb\x70\xba\x8d\x73\x1c\xd6\xdf\x37\x5b\xcd\xfe\xe7\x8a\x71\xd6\xec\xb7\x89\xe6\x59\xa7\x0b\x75\xb8\xac\x77\xfb\xcd\xc6\x55\xab\xde\xc5\x85\xdd\xbd\xec\xf4\x1c\x64\x7f\x8a\x64\xdb\xcd\xf6\x59\x17\xb9\x38\x17\x4e\xbb\x5f\x45\xae\x08\x03\xe7\x23\x0e\xa0\x77\x5e\x6f\xb5\x88\x95\x51\xbf\x42\xe9\xbb\x24\x1f\x34\x3a\x97\x9f\xbb\xcd\x0f\xe7\x7d\x38\xef\xb4\x4e\x1d\x04\xbe\x77\x50\xb2\xfa\xfb\x96\x93\xb1\x42\xa5\x1a\xad\x7a\xf3\xa2\x02\xa7\xf5\x8b\xfa\x07\x47\x63\x75\x90\x4a\xd7\xa0\x69\x99\x74\xf0\xe9\xdc\x21\x10\xf1\xab\xe3\x3f\x8d\x7e\xb3\xd3\x26\x35\x1a\x9d\x76\xbf\x8b\xc3\x0a\x6a\xd9\xed\xcf\x50\x3f\x35\x7b\x4e\x05\xea\xdd\x66\x8f\x0c\x72\xd6\xed\x5c\x54\x0c\x32\x27\x62\x74\x34\x11\xc4\x6b\x3b\x19\x15\x32\x35\x14\x3c\x82\x53\x68\x7c\xd5\x73\x66\x04\xe1\xd4\xa9\xb7\x90\x56\x8f\x90\x49\xc5\xe9\xe4\xaa\x61\x59\x98\x91\x74\x0a\xbc\x09\x83\x48\xd6\x4a\x12\xdb\xde\xc1\xc1\x41\x96\xcf\xcc\xcd\x26\x49\x4a\x6e\x35\x73\x28\x22\x65\x0d\xdd\x90\x07\xb7\x87\xf0\xd3\x39\xc3\xca\x88\x91\xe8\x42\x9b\xa5\xec\xa7\x0a\xcc\x00\xa8\x6a\x82\x21\x87\xe1\x8f\xc9\xcd\xc2\xad\x04\x1f\x1e\xc1\x40\xdc\x58\x92\xff\x41\x25\x1f\x7f\x27\x98\x20\x2d\x04\x1d\x81\x26\x8a\x0f\x70\xff\xb3\xf7\x26\x46\x40\x88\x89\x89\x47\x87\xb0\x7b\x44\xb9\x75\xcc\x5c\xff\x39\xf9\x87\x4c\xb9\x40\x5b\xa1\x1a\x16\x45\x36\xa1\x55\x64\xd2\xea\x55\x98\xf4\x6a\xe6\x84\xfb\x6a\x5c\xf3\x19\xd6\x4b\x66\xe9\xc1\xf3\x19\x0b\xec\xa9\xb8\xe4\x4c\x8b\xfd\x9e\xf2\xeb\x9a\xd9\xc8\x44\xb5\xfa\xb7\x31\x5b\x10\x9c\x3a\x1e\x9b\x9c\x7b\xa4\x2b\x81\x64\xaa\x76\xd5\x3f\xb3\x7e\x7e\x66\xf1\xf5\xbe\xeb\xf9\xdc\x7d\x57\x2f\x72\x6c\x6b\xe1\x4e\x0c\xe3\xd8\xa6\xa0\xa4\x1f\x03\xe1\xdf\x02\x47\x14\x89\x39\x17\x25\x36\xf5\x40\xdd\xd2\xef\x7c\x45\x49\x6f\x8c\x55\x5d\xaf\x28\x87\xaa\xfb\xc5\xb4\x79\x7b\x52\x25\xad\x09\x1b\x7c\xe3\xc8\x48\x3f\x08\x85\xc0\x9a\x42\x48\x59\x6d\xe0\xae\x64\xfe\x7c\x12\xc5\x86\xc6\xb6\x5c\xff\x6b\x2a\xd5\x21\x56\x9c\x88\x1d\x61\x2b\x41\x95\x09\x49\xee\xee\xfe\xdf\x11\x16\xe5\x88\x59\x33\x50\xf5\x1d\x0b\x8f\x40\xaf\x80\x6c\x02\xfc\xc0\x43\x5a\x2c\xc8\x01\xe5\xc4\x7d\xed\x28\x11\x69\xe4\x5b\x9e\x08\x44\x72\x08\x3f\x0e\xdf\xd1\xdf\x45\xf3\x43\xec\xfa\xbe\x96\x8a\xa2\x61\x30\xd2\x33\x6b\x66\x3e\xd3\x24\x7b\x2b\x77\xf0\xd4\xe1\xb1\xa0\xd2\x86\x7a\x94\xca\x0e\x70\xac\x92\x67\xcc\x63\x00\x24\xc1\x13\x67\xd2\x6b\xdc\x1b\x20\x91\xc0\xc2\x10\x1b\xa1\x24\x4a\xc4\x45\x43\x5d\xeb\x07\x98\x8d\x44\x6c\x9e\xe0\x02\xf3\xe7\x82\x66\x99\xd5\x7c\xb7\xbb\xfb\xc4\x4b\xa5\x54\x68\xec\x22\x31\x2b\x20\xdb\x41\x20\xbc\x6f\x85\xd8\x0e\xdd\x1b\x2b\x0f\x12\x14\x36\xbe\x29\x3c\xf4\x02\xe6\x26\xc4\x50\x8d\x0b\xf0\x75\x0b\x65\x66\x1c\x70\x53\x25\x96\x96\x44\xc1\x5a\xda\x50\x68\x2a\x9f\x5f\x3f\x75\x58\x15\xf5\x5d\x36\xce\xdd\x4a\x4c\xe5\x26\x27\xeb\xc5\x9c\xfb\x99\x2c\x81\xe5\x09\xbb\xf1\x7c\x76\xcd\xdc\xcd\xc6\x32\x76\xbd\xe9\xf8\x49\x15\xcd\x1f\x26\xae\xcf\x53\x79\x08\xaf\x35\xac\x24\x01\x0c\x87\x85\x2c\x96\xa1\x21\x11\x0c\x05\x29\x02\xee\xc3\x8f\xec\x80\xfe\x16\x13\xc3\x70\xb8\x60\x8b\x6d\xc8\x0e\x73\x49\x9e\x2e\x4b\xbc\x5b\xbb\xe0\x0a\xd6\xd5\x28\x93\xbc\xd4\xbc\xdd\x45\x23\xeb\x12\x95\xcf\xc7\x0d\x9d\x62\x49\x99\xbf\xf4\xbf\xbb\xda\x29\xab\x7e\x73\xde\xbd\xdd\xdf\x6f\x94\x17\xa0\x7d\x8a\x6b\x13\xf2\xf5\x96\x31\x58\xf4\x5e\x86\x5b\xbe\x22\xa7\x7f\xe6\xc7\x37\xb3\x73\x1b\xd0\x6f\x4b\x4a\x5f\x59\xed\xc0\x1e\x4e\x90\xb3\x17\x1e\xa8\x73\x02\xf3\x23\x86\x35\x47\x3c\xf4\xde\x03\x60\x95\x6f\x7e\xe0\x50\x2b\x1c\x37\xac\x4c\xcb\x5f\xad\x14\x9c\x3f\xcb\xc1\xb3\x71\xf2\x12\xa6\x9b\x14\xb3\x79\xf0\xec\x65\xc1\x73\x57\x6c\x6c\x7d\xee\x5b\x6b\xf6\xed\x0a\x82\x6d\x0f\x05\xcc\x3d\xd3\x5c\x72\x57\x38\xe4\x6a\xe0\xc6\x2d\x61\xc3\x9a\xb9\xc9\x5b\xdb\x27\x8e\x87\x69\xd2\x3c\x3b\x3b\xcb\x93\xaf\xcf\x3c\x91\xe8\x77\x72\xd3\xed\x41\x61\x43\xb0\x4f\xdb\x81\x42\xde\x1e\x88\xc0\x2f\x4f\xdc\x5e\x9a\x48\xa2\x1e\x0b\x9e\x01\x66\x0d\x05\x8f\x34\xd1\xbc\xaf\x58\x4a\xf0\x6f\x49\x30\x4d\x4f\xbf\x44\xc5\x84\x19\x22\x4d\x37\xe6\x0a\xe9\xff\xc1\x4a\x93\xfe\xeb\x37\x3f\x33\xdf\x2d\xa9\xd7\x2b\x33\x72\xb0\xb6\xf2\x61\x56\xc8\x67\xc0\x59\xf7\x86\xe5\x25\x73\xef\xc9\x47\xce\x26\xf4\xfe\xed\xde\xd7\xe3\xc7\xb6\x5b\x1a\xc3\x4b\x89\xb7\x3c\xfd\x66\x7f\xee\x3b\x63\x29\x29\x0a\x2f\x4b\xf6\xef\x59\xb2\x52\x25\x22\x1a\x3d\x9f\x69\x7f\x5d\x7f\x49\xe4\xb7\xfc\x80\xed\xd8\xce\x84\xfc\x0e\x51\x57\xd2\x30\xe4\x4f\xa6\x37\x21\x96\x4f\xea\x5e\xe2\xf0\x9f\x11\x87\x59\x6b\x3a\x0b\xb5\xe3\xc1\xf3\xb9\x99\xde\x23\x96\xd9\xe8\x9e\x2b\x40\xeb\xef\xe9\x3c\xb3\x32\xeb\xd7\x5d\xae\x55\xa1\x16\xcc\xcf\xea\xb3\x4a\xf0\xec\x91\xb1\x20\xd1\xb6\x84\xc7\xbd\x16\xbd\xf7\x5e\xd7\x7f\x69\xb0\x2c\x76\x98\xcb\x17\xcd\x9e\xa9\xa1\x9c\xb6\x5b\x2b\x3d\x25\x76\x6d\x2c\xa1\xee\xaf\x18\x4e\xd9\x55\x39\x6a\xa2\xb6\x2f\xc7\x3c\xae\x9a\x6e\xd8\xde\x2d\x5e\x69\x29\x75\xef\x4b\x57\xb8\x35\xd5\x78\xeb\x22\x13\x65\x1a\x6f\xa1\x4c\x5b\x67\xa7\x87\xac\xe0\xbb\x3a\xe2\x97\x85\xf5\xbf\xd9\xe6\x2e\x6e\xb7\x66\x57\x03\xe7\x1b\xae\x29\xe8\x19\xb6\x5c\x8b\x17\x15\x5f\xa2\xf1\x9f\x11\x8d\x2f\x9b\xae\x97\x4d\xd7\xcb\xa6\x6b\xdb\x83\xe5\x65\xd3\xb5\x35\x2d\xdb\x3a\x47\xe1\x6c\x3a\x8f\x3b\x79\xc0\x51\xe8\x0c\x65\x0e\x79\xf2\x9b\x18\x85\xab\x49\x0b\x37\x4d\xe6\x8e\x3e\x38\x38\xb8\xeb\x80\xbb\x78\xb2\xbb\x7a\x24\xb9\x1d\x4d\xc3\x36\xb5\x2f\x4f\xd9\xba\xec\xaf\x6d\x5d\x4a\x0f\xd1\xee\x73\xf9\x42\x6f\xb3\x74\xaf\xa1\x78\x0b\x6b\x31\x5d\x15\x3f\x85\x7d\xba\x80\xd8\x5f\xcc\x56\x5a\xa3\x8d\x53\x15\xea\x04\x83\xdb\xcd\xce\xe1\x56\x73\xc7\xca\x7d\x87\xe5\xcc\x70\x6c\xe3\x32\x3f\xc9\xfe\x6b\x14\xd3\xc4\xb6\xb5\xb5\x6b\xae\xd7\x65\x2a\xce\xf3\xd7\xb1\x4d\xb7\x58\x09\x42\xd7\x81\x4f\x0c\xa3\xfc\x5b\xdb\x38\x95\x63\x81\x1c\xbf\xc3\xa7\xa6\x2b\xa4\x8a\x1f\x34\xfd\x1d\x9f\x9d\x7d\x9f\xaf\xce\x36\xff\xe8\xec\xfb\x7d\x73\xb6\xc0\x73\x03\x4b\xa6\x49\xf0\xf0\xcf\xcd\xfe\x13\x00\x00\xff\xff\x56\xe5\x06\x71\x48\x3f\x00\x00")

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

	info := bindataFileInfo{name: "template/default.tmpl", size: 16200, mode: os.FileMode(420), modTime: time.Unix(1469559122, 0)}
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
