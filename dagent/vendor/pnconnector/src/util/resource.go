// Code generated for package util by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../resources/events.yaml
package util

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ResourcesEventsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5d\xdd\x52\x1c\x39\xb2\xbe\x9f\xa7\x50\x70\xb1\x3b\x13\x01\x3e\xa6\x7f\x6c\xe0\x0e\xf3\xe3\xc3\x86\x31\x5e\xda\xf6\x9c\x73\x35\xa1\xae\x52\x77\x6b\xa8\x92\x6a\x24\x15\xd0\xb3\xb1\xef\xc2\xb3\xf0\x64\x27\xf4\x53\x3f\x5d\x52\x55\x67\xab\xbd\x73\x26\x62\x22\x0c\xa8\xf4\x65\xa6\xa4\x54\x2a\x95\x99\x3a\x3a\x3a\xfa\xe9\x23\xc7\xd9\x19\xba\x7a\x24\x4c\x49\xf4\x8d\xd1\x05\x4d\xb0\xa2\x9c\xfd\xf4\x9d\x08\x49\x39\x3b\x43\x07\x8f\x6f\x4e\x47\x07\x3f\xe5\x3c\x2d\x33\x22\xcf\x7e\x42\xe8\x08\x31\x9c\x93\x33\x74\x70\x71\x77\x7b\x7b\xf7\xf9\xe0\x27\x84\x10\x4a\x78\xc9\xd4\x19\x3a\x3d\x3d\x35\x3f\xd2\x74\xa6\xb0\x50\x67\xe8\xad\xfb\xf1\x8a\xa5\x67\x08\x35\x7f\x67\x0b\x7e\x66\xfe\xa5\xfb\x4b\x78\x4a\xaa\xa6\xfa\xbf\x8c\x3c\x92\xec\x0c\x1d\xdc\x7c\xbe\xbe\x3b\xa8\x7f\x9b\x13\x29\xf1\x52\x03\xcf\xca\x24\x21\x52\x36\x7f\x2a\x04\x9f\x67\x24\x3f\x43\x07\xcd\xef\x24\xcf\x4a\x65\x59\x38\xd8\xa0\xfa\xd3\xcd\x06\xc9\xc7\x6f\xdf\x6e\x92\x7c\xfc\xf6\x6d\x87\xea\xe3\xb7\x03\x64\xd7\xcd\x5b\x94\x6b\xc2\x7d\xba\xe9\x9c\x2f\xb8\x44\x54\x22\xa9\x91\x48\xea\xd1\xef\x13\xef\x83\x1d\xef\x0a\xa6\x88\xc8\x29\xc3\xb1\x78\x23\x18\x5e\x92\x51\x24\x89\x78\x24\x42\x63\x52\x46\x15\xc5\x19\xfd\x33\x12\x74\x0c\x03\x65\xe4\x49\x03\x13\xa6\x34\x68\xc2\x19\x23\x49\x2c\x9f\x13\x30\x9f\x0e\x2e\xa5\x72\x3f\xc4\x29\x0c\x51\x90\x3f\x4a\x22\x15\xca\xe5\x52\xc3\x0a\x92\x10\xfa\x18\x09\xf9\x0e\x0a\x29\x0b\xce\x24\xa9\x30\x25\x61\x2a\x06\xef\x18\xb8\x32\x5a\x93\xa7\x10\xa4\xc0\x82\xb2\x25\x22\xcf\x34\x0e\x14\xb8\x42\x5a\xa0\x6a\x25\x08\x4e\xd1\xef\x9c\xb2\x38\xc1\x1e\x03\x57\x09\x61\x78\x9e\x11\x24\x48\x29\xc9\x11\x4e\x53\x11\x05\xe6\xad\x8e\x5f\xcf\xef\x3f\x03\xc0\xd0\x02\xd3\x2c\x92\x41\x6f\x79\x5c\xdd\xdf\xdf\xdd\xfb\xa0\x92\x27\x0f\x44\xa1\x44\x10\xb3\x91\xec\x03\xe9\xad\x8f\x61\xc8\x39\x65\xa9\x9e\x37\x7b\x20\x7a\xcb\x63\x18\x31\xa3\x52\x91\xbd\x58\x7c\x0f\x04\x24\x05\xcf\xb2\x1f\x22\xd4\x93\xdd\x58\xc4\x49\x42\x0a\xb5\x0f\xe0\x29\x10\x30\xc7\xcf\x95\x72\x25\x42\xf0\xa8\x95\x31\xf2\xf4\x4d\x1f\x98\xfd\x87\x56\x6c\x7b\xce\x99\x91\xa7\x6d\xb6\x40\x5a\xfd\xbd\x27\xa8\xa7\x6e\xc2\x1a\xe0\xeb\xcd\xed\xd5\x25\xba\xfb\xf6\x35\x0a\xc4\x53\x33\x3d\x9c\xdd\x7c\xfe\x7e\xfe\xe9\xe6\x12\x7d\x39\xbf\x3f\xbf\x8d\x41\x1a\x03\xb7\x89\x2f\x77\x33\x74\x33\x43\x1f\xbe\xcd\xfe\x17\x06\x53\x1b\x7d\x37\x1f\xf8\xf5\xdd\x0c\xcd\x14\x56\x04\xdd\x62\x86\x97\x44\x6c\x58\x81\x23\xcf\x0a\x3c\xf6\xac\xc0\xd1\x90\x15\x78\xec\x1b\x66\x97\x57\x1f\xbe\x7d\x0c\xac\x2c\x43\x44\xc2\x99\x22\xcf\x0a\xe1\x34\x8d\x9a\x03\xc7\x50\xc3\xcc\xc1\xad\x30\x5b\x46\x01\x8d\xc6\x9e\xea\xbf\xbf\xfa\x72\x77\xff\xf5\xb7\xaf\xf7\xe7\x17\x57\x01\xcb\xd3\xc9\x7a\x2d\x15\xc9\xd1\x3d\x49\xf8\x23\x11\x6b\x74\xc3\x0a\xc1\x97\x82\x48\xb9\xe3\xd8\x7d\xe7\x59\x99\x93\xd0\xa0\x8d\xbb\x83\x36\xf2\x4c\xf7\xd1\xd0\xa0\x8d\xc0\xa6\xbb\xa5\xc1\x2a\xe0\x18\x29\x8e\xc0\x76\xbb\x43\x4a\x49\x46\x62\x91\x80\x13\xc3\x21\xe5\x5a\x96\x91\x48\x40\x33\xdd\x21\x95\x6c\x1f\x2c\xa0\x7d\xee\xb0\xcc\xb2\x42\x8a\x23\xb5\x22\x66\x9f\x8e\xc2\x04\x5a\xe8\x16\xf3\xf5\x45\x90\x9c\x3f\x92\x14\x2d\x04\xcf\x35\xf0\xeb\x4b\x2c\xb2\x6f\x38\xf7\x28\xf8\x15\xa9\xce\x07\x24\x45\x8f\x6e\xee\x70\x22\x11\xe3\x4a\x5b\xd0\x01\x78\xf3\x91\x6b\xfa\x44\xd5\xca\x88\xc8\xeb\x44\x2f\x3e\xc4\x45\xf5\xe3\xcd\x65\x5f\xb7\xcd\x39\xfb\x8a\x29\x63\x4b\x6b\xe5\x26\x04\x49\x54\x7f\x5f\x78\xa1\x9b\x26\x2b\x92\x3c\xe8\x9d\x50\x35\x14\xb5\x24\xd6\x12\x87\xb7\x78\x7a\xc4\xc1\x39\xca\x31\x5b\xbb\xce\x7c\x4d\x53\x2f\x64\xcc\x0c\x2b\xf3\x7a\x4d\xa3\x39\x49\x70\x29\x89\xa1\x25\xc7\xcf\x34\x2f\x73\xc4\xca\x7c\x4e\x04\xe2\x8b\xaa\x43\xa4\x56\x58\x99\xaf\x5b\x5f\x52\x89\xc8\x73\x42\x48\x5b\x91\x37\x52\xb9\x27\x4a\xac\x1d\xc3\x66\x82\x68\x86\x4b\x7d\x6e\xd4\x54\x8b\x8a\x56\x84\x73\x6e\x0f\x3d\x52\xe9\x16\x55\xe7\x9b\x9c\xb4\x44\x02\x34\x01\xae\x0c\x65\xb6\x3b\x73\x1c\x70\x70\x92\xfe\x49\xc2\x53\xc3\x9b\x0a\xba\xa9\xe3\x50\xa2\x73\x21\xf0\x1a\x25\xb8\xc0\x09\x55\xeb\x00\xbf\x17\x7a\x50\x8d\x14\xa5\xdd\x01\xaa\xb6\x08\xb3\x14\x19\x59\x2c\x31\x65\x1e\x43\xbe\xed\x16\x66\xe8\x7b\x6b\x4e\x51\x89\x14\xe7\x48\xae\xb8\x18\x9e\xe7\xa6\x35\xd1\xf3\xd3\x8e\x97\xea\x7e\xd4\x9d\xc5\x78\xe3\xcb\x39\x51\x4f\x84\x30\x34\x32\x3c\x8c\xa6\x53\xbd\x9f\x0a\x9c\x28\x22\xfc\x91\xf1\x2d\xc2\xad\x8c\xbc\xbe\x54\xac\x64\x9c\x2d\x7b\x67\xad\xcf\x45\xe7\x83\x61\x2e\xcc\xdc\x6d\xad\x62\x33\x31\xb6\x31\x03\x9c\x66\x9d\x51\x49\xcb\x22\xa3\x49\x70\xaf\x44\xe7\x1b\xca\x07\x37\x6d\xed\xd7\x38\xd3\x67\xf1\xb5\x5d\x09\x72\x80\xb5\x94\x2e\x16\x44\xe8\xf3\x42\x8b\x49\x9f\x01\xe0\x61\xf9\x1b\xb3\x47\x9d\xf6\x1a\x69\x75\xb8\x65\x30\xb4\x35\x87\x29\x93\x08\x23\xa9\x84\x55\x69\xd8\x78\x89\xb4\xa8\x71\x96\xf1\xa7\xa0\x72\x68\x54\xa6\x37\x50\x39\x21\x4a\x7a\x7f\x12\x65\x16\x50\x06\xbe\x01\x1d\x66\xf2\xbe\x5e\xda\x66\x4d\x9b\xf9\x83\xc5\x72\x27\x45\x50\xcd\xbb\x8d\xef\xc2\xea\x6e\x43\xbf\x17\x5c\x4a\x1a\x56\x40\x2d\x46\x80\x6b\x27\xc0\x88\xcc\x71\x96\xed\xce\xc8\xeb\xcb\xe6\x87\x21\x45\x96\x53\x66\xb6\x03\x3d\x8e\x89\xaf\x46\x8d\x52\x10\x9a\x6f\x9f\x1f\xe0\xf2\xe9\xf0\x63\x27\x0d\x5d\x86\xbc\x51\x41\x8e\x5e\x5f\xcc\x77\xf5\xe2\x36\x1f\xdb\x05\x36\xcf\x78\xf2\xb0\xa9\xee\x1b\x1e\x6f\x58\x51\xaa\x0d\x5e\x14\xd7\x5b\x5b\x5e\x66\x8a\x16\x19\xd1\x5b\x9f\xd7\x41\xc3\xde\x64\x37\x9d\x5d\x2d\xed\x3e\x33\x10\x9d\x2b\x45\xf2\x42\x69\x22\x4c\x9b\x46\x81\x55\xcb\xa9\xaf\x8b\x86\xa5\xcf\x5c\xad\xcc\x9c\xe3\x28\xe5\x3e\xc5\xbb\x29\xe7\x0a\xae\xdf\x74\x6d\xd3\xec\x5a\xf5\x53\xed\x77\x03\xa5\x7b\xea\x1f\x54\xd8\x82\xfb\x74\x5f\x38\xc3\xc6\x9e\x1f\x2a\x31\x39\x7a\x7a\x89\xb7\xa6\xb2\xfb\x06\xf7\x7d\xd5\xd0\xda\xe2\xd9\x7d\x54\x19\x31\x66\xdd\x3b\x36\x37\x2d\x3b\x8f\xa5\x77\xbe\xb3\x66\xc3\xef\x53\xf3\x74\x8d\x69\xa6\xa1\xac\x49\x54\x41\xe5\x44\xe1\x18\xd3\xfa\x9d\xef\xb0\x19\x86\xe5\x05\x61\x7b\x83\x7a\x7a\x60\x0b\xa8\xf1\x47\xef\x0b\xea\x3b\x70\x86\x41\x9f\x04\xfd\x01\xf2\xf5\x3d\xc5\x61\xd4\xef\x0d\xce\xeb\x8b\x71\x87\x30\x85\xe6\x82\x3f\x10\x16\x83\xfb\x1e\x3a\x9d\x6e\x49\xce\xf5\x16\x65\x95\x39\xe5\xec\xf5\x65\x81\x69\x56\x8a\xd0\xfa\x40\x09\x96\x6e\x1d\xcb\x15\x2f\xb3\x14\x31\xf2\xa8\x8f\x04\x49\x52\x0a\x74\x84\x56\x04\x17\xad\xae\x50\xb7\xa7\x66\xcd\x7c\xed\xb5\x7c\xdf\x43\x67\xe4\x0d\x7b\xc4\x19\x4d\x11\x65\x29\x79\xee\xf1\x92\x6e\x27\xd9\x7c\xfd\xb3\x1b\x65\x9a\xfe\x82\xa8\x36\x42\x18\xce\xb2\x35\x5a\x0a\xcc\xdc\x91\x86\x5a\xb0\xe0\xa6\x61\xdb\xa3\x8c\x2f\x69\xf2\xfa\xd2\x26\xa4\xc5\xd5\xae\x53\xde\x48\xf1\xcd\xeb\x0b\x23\x4f\xaf\x2f\xf5\x51\x31\x82\xc1\x6f\xf6\xde\x43\x71\xb4\xa4\x8f\xa4\x39\x75\x1e\xa2\x94\xc8\x42\x4f\xf1\x96\x55\x65\x5c\x49\x95\xa1\x96\xe3\xe7\x78\x7e\xa1\xab\x4d\xef\xdf\x12\xdb\x73\xb0\x23\xa2\x63\xeb\xc2\x59\x3d\xf7\x0e\xf2\x55\xcf\x5b\xcd\x68\x18\x57\x27\xfe\x0e\x8f\x05\xa3\xad\x33\x47\xcd\xd7\x5d\xa9\xb4\xb9\xf0\x07\x97\x48\x60\x16\x32\x2a\xcf\xd1\x23\xce\x4a\x82\x32\x22\xcd\x49\x9a\x6d\x5a\x57\x85\x39\x07\xe8\xa1\xd3\x7d\xd8\xa6\x4f\x58\x56\x46\x36\xc8\x44\x6b\xbe\x6c\x9f\xd4\x2b\x33\x7d\xe3\x04\xfa\xc6\x63\xf6\x14\xce\xac\x3d\x06\x9b\x39\xd3\x67\x18\x34\xae\x86\x8e\xa7\x81\x0b\x94\x71\x9c\x92\xd4\x8c\x1a\x2f\x55\x75\x59\xdf\x6f\x1c\xd4\xca\xc3\xed\xb0\xd6\xce\xb0\x9f\xf9\x6c\xf8\x36\x4e\x1f\x1b\xce\x49\x7b\x8d\xcb\x2c\x70\x80\xae\x38\xe0\x79\xae\x25\xd7\x70\x52\x10\xb1\xe0\x22\xd7\x8a\xc2\x8e\xe1\xec\xeb\xdd\x17\xeb\x69\x06\x68\xea\x53\xdf\x28\xee\xa3\xef\x92\x33\x37\xb7\x7b\xb4\xdd\x8c\xeb\xe5\xa3\xff\x26\x51\x8e\xd7\x6e\x61\xa4\xa5\xa8\x8f\x1d\x82\x27\x44\x4a\xfd\x23\x5f\xb4\x5d\x5d\x87\x76\x36\xe8\x25\x53\xce\xa5\xfe\x1d\x53\x7a\xaf\x17\x56\x91\xe7\x6e\x78\x9f\xb8\x78\x40\x4f\x24\xcb\xde\x84\x8e\x6f\x1a\x18\x2d\xb8\xb0\x24\xa0\x15\x66\x69\xa6\xa1\x70\xa6\x07\x76\xb9\x42\x54\x55\x62\xb3\x94\x19\x5e\x4a\x49\x04\xb2\x90\x89\x67\x17\x1d\xfb\x3e\xe9\x5e\xf1\x68\xca\xed\x08\x6a\x84\x3e\xef\xb4\x6f\x56\xb7\x95\x06\xe3\xad\x6e\xba\x5d\x34\xac\xce\xaa\x36\x12\xe5\xa5\xec\x78\xce\x16\x5c\x38\x5b\x53\x33\xdf\xe3\xb2\x0a\xdc\x90\xf4\x32\xd6\x31\xbf\x57\x58\x22\x49\x54\x4d\x28\x43\xec\x0f\xf6\x1f\x22\x12\x3e\x39\xeb\x8d\x4b\x29\x9c\xac\xcc\xf9\x5c\x16\x38\x31\x3b\x4f\x2d\xd2\x5e\xad\x45\x17\xe6\xca\xb3\xfe\x4a\x1a\x5d\x27\x0b\x92\xd0\x05\x25\x69\x35\x87\x3b\x63\xa3\xa7\xe6\xcf\xe4\xf9\x0d\x3a\xca\xd1\x68\xfa\xee\x97\xf6\x45\xc9\xed\x87\xfb\xe0\xdd\x96\x17\xe1\x34\x1a\x7b\xd7\x24\xe3\xa1\x6b\x92\x31\xf4\x9a\xe4\xdc\x1a\x42\xda\xa8\x33\x36\x96\xb4\x81\x5a\x11\xe6\xdc\x18\x7a\x5f\x92\xcf\x05\x4a\xb1\xc2\xd5\xe9\x40\x2f\x6a\x63\xc7\x46\x81\x02\xaf\x4e\x6a\x50\x9c\xa6\x7b\x22\x02\xaf\x50\x0a\x2c\xa8\x5a\x5b\x8f\xca\x5e\x62\x05\x5e\xa3\xb8\x39\x57\x96\x34\xdd\x1f\x14\x7a\x9b\x91\x92\x47\x9a\x58\xef\xc7\x82\x97\x2c\xe6\x9e\x68\x0c\xbd\x2a\xd8\x10\xa8\xb6\xda\xa3\xc0\x80\xee\x1d\x5f\x9a\xd1\x88\x40\x77\x66\x3d\x43\xf7\x13\xa6\x37\x5b\xfa\xd1\x1e\xc8\x7a\x3f\x30\xdf\x89\xd9\x13\x6f\xd0\xd5\x32\xb1\xb2\xf4\x9d\x8d\x7d\xa1\x1b\x95\x30\xed\x4d\x48\x5c\x74\xca\x68\xec\x3b\x03\x07\xf0\xb4\x7d\x85\xa3\x43\x61\x46\xe3\x31\x34\x7a\x63\xa9\xf7\xd5\xd6\xfc\x84\x4b\xb3\xde\x78\x2e\xcd\xba\x85\x6d\x3d\x13\x6f\xeb\x99\x0c\x6d\x3d\x13\xe8\xd6\x73\x63\x23\x4e\x91\x4c\x70\x8c\xfb\x60\x02\xdd\x6f\xee\x49\x34\x02\x70\x73\xf9\x44\xa5\x72\xaa\x30\x0a\x06\xb8\xa3\x18\x18\xac\x8f\x31\xaf\x2f\x7b\xa0\x05\xf7\x93\xff\x0a\xab\x89\x6b\xca\x52\xcd\xd9\xcf\x25\x4d\x7f\x89\x42\x03\x87\x04\xea\x83\x5f\xec\xda\x99\xf8\x61\xb2\x3d\x30\x6e\xc3\x32\x1e\xc2\x78\x34\x68\x10\xa0\x43\x4b\x32\x2e\xe3\xb5\xd0\xe4\xad\x17\x01\x38\xb8\x19\xbf\xbe\xf4\x2b\x75\x64\x17\x7e\xeb\x36\xc2\xde\x1e\x54\x71\x10\xda\xa0\xf5\xbf\x1e\x20\xcd\x8b\x15\x0c\x93\xf6\x99\x33\x03\x61\x7c\x05\x97\xe1\xb9\x8b\xee\x58\xb6\x46\x29\xd1\xa6\x39\x49\x6b\xd9\x59\x7f\x80\x23\x10\x40\x12\x34\x82\x39\xc5\x24\xe7\xcc\x06\xf6\xc7\x8c\x0a\x34\x68\xb9\xc6\xe1\x45\x14\x0c\x50\x07\x39\x71\xd9\x93\x4d\xae\x4f\xc6\x29\x51\x91\x91\xee\xa3\x89\x6f\xb4\x0c\xa1\x6a\x75\xf4\x63\x60\x81\xb6\xae\x3b\xbf\xc5\x2b\x41\x3f\x52\xb9\x8f\xbf\x7d\x81\x3c\xc5\x14\x5e\x20\x1b\x1c\xc5\xab\x0a\x3f\x3c\xb9\x4f\x55\xfc\x18\x38\xa0\x66\x2a\x6b\xff\xae\x59\x6e\x28\xe7\x8c\x2a\x2e\x68\x28\x72\x02\x67\x59\xeb\xef\x6e\xf9\x48\x84\x45\xed\x5c\x78\x7d\x11\x25\xd3\x07\xfc\x43\xc4\x05\x62\xbc\x6a\x2e\xfb\xe2\x9c\x7c\xba\x81\x6a\xab\x4d\x37\x2f\xb6\x91\xfd\xfa\xe2\xd3\xfd\xfa\xd2\x22\xdc\xf4\x52\x90\x34\x9a\x6e\x70\xb4\x74\x4b\xe9\x14\x24\x45\x25\x23\xcf\x85\x59\x95\xd9\xda\x27\x1d\xd2\x78\x80\x26\x68\x14\xad\xcd\x13\x21\x7b\x2c\x26\x3f\xb4\xa5\x07\x4a\x90\x8c\x60\xb9\x1b\x54\x6d\x1b\x1b\x47\xf1\x96\xa8\xd5\xa9\x67\x13\xbf\x1f\xb2\x89\xa7\xbe\x4d\xdc\x77\x52\x32\x5e\x6a\x77\x01\x9d\xe2\xfa\xda\xfc\x4d\x84\xbc\xa6\xbe\x85\x0c\x41\x35\xd7\xde\x7b\xc1\x42\x4f\x4d\x16\xb6\xe7\xde\x63\x17\x40\xe8\xc9\xc9\x02\x6e\x86\x45\xc6\x21\x42\x93\x6b\x2e\xf7\x77\x90\x4c\xc1\x36\xb4\x5a\x91\x56\x00\xa4\x94\x69\xe3\x54\xad\xe2\xe1\xc6\xc6\xf9\xd5\x3a\xd7\xee\x42\x07\xd4\xc8\x66\x8f\x79\xb5\xb5\xec\x75\xfe\x9e\x82\x0d\xed\x99\x74\x71\x38\x15\xbf\x4b\xe3\x5b\x14\xf6\xce\xea\x5f\xff\x46\xf3\xb5\x0a\x04\x97\x42\x48\x80\xa6\xdf\xcc\xda\xe2\xae\xef\xf5\x34\x51\x71\xb8\xd0\x2c\x9c\xdb\x0f\xf7\xaf\x2f\x26\xf8\x20\x5a\xcc\xbe\xc9\xdc\x8f\xe5\x42\x0e\xe2\xb1\xa0\x3a\xc9\xdc\x70\x1d\x29\x9e\x11\x81\x59\x42\x8c\x62\x45\x7b\xf2\x09\x55\x4c\xbf\x0a\xce\x96\x1e\x05\x39\x51\x2b\x9e\x22\xb5\x2e\x62\x76\xaf\xa9\x6f\x53\xf7\xa0\x1f\xfc\xeb\xdf\xe8\x0b\x16\x8a\x9a\xbb\x83\xfa\x12\xc1\xb0\xed\x67\x50\x43\x90\xa1\xda\xaa\x41\xe6\xcc\xdc\x78\xee\x03\x0a\xd5\x5a\x46\xd8\xaf\x2f\x56\x37\x93\x47\x93\xab\x1b\xa5\x27\xc1\xc9\x80\x2e\x50\xc2\xde\xa3\xe3\x0c\xe1\x34\x15\x44\xca\x3d\x26\x16\x54\x49\xd8\x89\xa5\x67\x90\xb9\x83\xc4\x76\x3f\xea\x39\x0f\x7f\x5d\x11\xdb\xf4\x67\xdd\x76\x5e\x2e\x16\x5a\xb1\xdb\xdb\xcb\x14\x2b\x7c\x24\x15\x17\x78\x49\x7e\x69\xdd\x41\xcd\xd7\x46\xf7\xb4\x3b\xae\x6f\x4d\x71\xa2\x4a\x9c\x55\xbf\x35\x3d\x1b\x7b\xac\x8a\xca\x0d\xdd\x97\x36\xf7\xf5\xb6\x7d\x5f\xdc\xe4\xd4\x37\x4d\xc3\x27\x2a\x73\xe1\x5a\x59\x66\xa8\x39\x37\xc6\x88\x1d\x1a\xb7\xcd\xb4\xe2\xc8\x31\x65\xe6\x7a\x6e\x3f\x23\x74\x0a\x8d\xaf\x66\xdc\x8d\x42\x38\xa7\x03\x82\x04\xbc\x39\x60\x84\xa4\xee\x9a\x78\x41\x85\x34\x19\x6a\x7a\x9e\x58\x9f\xc9\x1e\xac\x42\xb3\xfc\xbc\x19\xb7\xc2\x12\xcd\xb5\xb9\x11\x99\x2b\x37\x9a\xfa\x9e\xfe\x5d\xa1\xbb\x1e\xa3\x5d\xc0\x81\x43\x6c\xd3\x41\xcd\x85\x72\x9a\x7a\xa4\xf8\x0b\xda\x23\x36\x2e\x9e\xde\xa7\x17\x38\x51\x1a\x7a\xdd\xd4\xd8\x4e\xb2\xd6\x41\x6d\x62\x37\x33\x8f\xaa\xe4\x01\x13\xc0\x1c\x5e\x59\x03\x54\x43\x9d\x3d\x2d\x02\xad\x1f\x20\xc9\x08\x16\x91\x43\x0b\xdd\x95\x36\xc7\x76\xeb\x6a\x72\x54\x86\x82\x3a\x06\x88\x81\x6e\x58\xbb\x12\x53\xb2\x07\xc6\x9f\xd8\xc6\xc8\x3d\x99\xad\x67\x23\x41\x62\x80\x32\xa8\xe1\xbd\x49\xd9\xd0\xc4\x6f\x66\x8d\xa3\xaa\x75\xdc\x74\x31\x49\x46\x86\x00\xea\x02\xa1\xd5\x03\x33\x47\x2a\xac\x4a\x89\xca\x22\x8d\xcc\x04\x9d\x4e\x81\xca\xc8\xba\x0e\x5c\xce\xf2\x19\xba\xd2\xeb\x17\x7d\xe6\x22\xc7\xd9\x81\xdf\x29\xd0\xb1\x1b\xec\xf4\x92\x2c\x05\x4e\x49\x1a\xe8\x16\xe8\xb9\x0d\x76\x7b\x4b\x4d\xdc\x55\xa0\x57\xe0\x5a\x0d\xf6\xfa\xc1\x44\x08\x07\x3a\x05\x3a\x61\x3b\x9d\xf6\x0a\x14\x58\x2a\xa5\xd3\xdd\x80\x28\xbd\x45\x00\xea\xf0\x9e\xcc\x4b\x9a\xa5\x61\x39\x7a\x86\x22\xa8\xcb\x99\xe2\x85\xdf\x99\x1f\xe9\xd9\xd3\xd9\xad\xcb\x0e\xb0\x9d\x9a\x41\x09\x56\x5d\xfa\xe6\xb2\x21\x6a\x0f\x58\xed\x13\x73\x3c\x6d\xb9\x29\x3e\xf1\xbc\x62\x27\x43\x5e\xb1\x13\xe8\x4d\x71\x23\xd1\x4b\x32\x2f\x97\x19\x5f\xfa\x5d\x01\x57\x53\xd3\x55\xe5\xef\xf4\xbb\x02\xae\xa0\x56\x57\x1b\x75\x21\x5a\x3d\x0d\xe6\xe1\x23\xbf\x4b\x61\xbb\x44\x55\xde\x3d\x12\xa4\x68\x72\x1e\x5b\x1d\x83\x0b\x11\x6d\xf6\x17\xa1\xfc\x4e\xa0\x15\x88\xdc\x19\x24\x31\xe9\x74\x85\x51\xef\x19\x4f\x1e\xa2\x34\xee\x89\xef\xf8\x01\x81\xbe\xbe\x50\x89\x4a\x16\x8f\x0b\xbd\x3c\xac\x04\x9b\xf0\xbc\x30\xd9\x34\x2e\xd6\x6a\x51\x66\x01\x0f\x3b\x04\x18\xea\x16\x39\xa8\xa0\x05\x91\x65\xa6\x1a\x52\x5c\x1e\x43\xcc\x59\xfd\x04\x1a\x15\xd5\x0b\xee\xd6\x52\x1c\x38\xf4\x98\xc3\x6b\x38\x85\xc5\x92\x04\xc2\x9f\xd5\xaa\x73\xb2\xb4\x89\x11\x36\x06\x3a\xe3\x36\x57\x0b\xb3\x35\x2a\x2a\xf7\x0a\x84\x3e\xe0\xde\xc7\xb8\x3d\x89\xa9\x9a\x4e\x9f\x40\x5b\xda\xc3\x91\x94\xba\x9d\x07\x42\x03\xd4\x6c\xbd\xae\xed\x31\x7b\xd3\xe6\x28\x39\xac\x02\x86\x4c\x29\x3b\x5e\x74\x22\xbf\x61\x17\x23\x17\x9c\x2d\xe8\x12\x16\x34\x74\xea\x6d\x05\x43\x85\x04\xeb\xe6\x5b\x85\x9c\x18\x1a\xd0\x82\x9a\x32\x59\x38\x45\x29\x67\x31\x67\xdb\x53\x68\xf4\xd0\x92\x28\x97\x9e\x10\x8d\x04\xdc\x9a\xac\x4f\xca\x62\xc5\x87\x6e\x9e\x42\xb7\xaf\x36\x5c\x64\x3c\xde\x69\x20\x4c\xb4\x2f\xf4\xba\xaa\x89\xe7\x06\xd0\x78\x93\x62\xfd\x6d\xa7\x81\x90\xd1\x6d\xb8\x0f\x64\x1d\x8f\x17\xb8\x77\x0d\x67\x0a\xb5\xa7\xa7\x89\x31\x8a\x95\x6c\xa0\x48\xd5\x56\x44\x97\x52\x1c\xcf\x25\x34\xff\x6b\x03\x73\x2f\x97\xfc\x69\xa0\x64\xd5\x76\xc9\xee\x19\x69\x79\x3a\x82\x26\x3c\xfe\x2e\x39\x43\x29\x4f\x2a\x8d\xcd\xe7\xbf\x93\x24\xb0\xef\x54\x93\xac\x51\x16\x7f\xb3\xeb\xcb\x64\xbf\x9a\x5f\xd8\x2c\xa0\xcd\x89\x71\xd8\xba\xcf\x40\x7f\x0b\xb3\x15\xd4\xc4\xb7\xb8\x28\xb6\x95\xc3\x1a\x7b\x95\x95\xc6\xc7\x1d\x15\xdc\x5c\x79\xdb\xf8\x5c\xbe\xad\xcb\x91\xd7\x65\x37\x0b\xa1\xee\xf2\xe3\xc5\xf0\x0e\x31\xf6\xc2\x4a\xc7\x43\x61\xa5\x63\x70\x58\xe9\xc7\x0b\xa4\x04\x5d\x2e\x49\x94\x87\x68\x0c\x0e\x2b\xfd\x78\x11\x5f\x14\x76\x0c\x0e\x2d\xfd\x78\xf1\xfa\x12\xb7\xf1\x8c\xc1\x71\xa5\x1f\x2f\xea\xea\x3a\x91\xd1\x70\x63\x70\xd4\xdd\x77\x9a\x28\x9a\xd7\xe6\x7a\xc2\x99\x54\xa2\x4c\x54\xcc\x42\x1e\x83\x63\xf0\x3e\x71\xac\x6d\xd6\x47\x22\x24\x41\x39\x8e\x08\xc4\x1b\x83\x03\xf1\x0c\x96\xdd\x5c\x4d\x1d\x86\x5d\x8b\xa7\xdd\x12\x85\xaf\x67\x1b\x8b\x67\xda\x5d\x3c\x13\x6f\x6d\x4f\x86\x16\xcf\x24\x50\x35\xad\x27\x72\xe6\xd6\xfe\xa3\xb9\xa2\x4a\xf5\x89\x7b\x19\x0a\x79\xda\x2a\xb3\x49\x20\x31\x2e\x2c\x33\x0f\x55\x73\xa1\x95\x61\xdb\x42\x87\xe3\x8e\x02\x09\x79\xc1\x23\xc5\x4d\x5d\xf4\x18\xdd\x5e\xcf\x36\x92\xbd\x77\xc3\x03\x5e\x44\x19\x47\xcc\x9e\x50\xc0\xa3\x5a\xe5\xcb\xd9\x0f\x0c\x78\x34\x6b\x4e\x1d\x84\xfd\x51\x92\x92\x98\x3a\xcf\xb9\x0c\xc4\xc9\xf5\x36\xd5\xbf\xba\xbd\x9e\x1d\xda\x3c\x4e\x77\x90\x22\xcf\x05\x66\x29\x5a\x08\x62\x2b\x1b\xff\x13\x42\x34\x30\x05\xe6\x82\xe7\x05\x4e\xd4\x40\x75\xd6\x76\x93\x84\x97\x59\xca\xfe\x6e\xa2\x3b\xb4\x42\x46\x69\x69\x02\x03\x4d\x1c\x11\x33\x99\xa3\x86\x4a\x93\x19\xf8\x9f\xa1\x32\x86\x84\x1a\x30\x90\x11\xd8\x17\xe4\x71\x3d\xd3\xf6\x4e\x74\xe5\xbf\x49\x20\x13\x70\x18\x6a\x8f\x32\x83\x93\x40\x06\x60\x1f\x58\x57\xc9\x44\xda\x8f\x93\x40\x0a\x60\x5f\x9d\x55\x17\x68\x50\x60\x81\x73\xd2\x2a\xd7\xb5\x1b\x1c\x34\x7c\x43\x4f\x04\x73\xa1\x12\x85\x02\x75\x31\x54\x97\x36\xf1\x48\xbb\x5f\x7b\x35\xd5\xe9\xa3\x00\x77\xbf\xcd\x9a\x9b\x20\xde\x32\xc2\x52\x98\x8c\xc1\x21\x63\x8c\xa3\x9c\xa4\x14\x9b\xd9\x78\x7b\x3d\x8b\x02\x83\xc6\x89\xb9\x82\x55\x1b\x25\xc1\xe6\x34\x66\x5b\x0f\x64\x87\x6e\x95\xa7\x3e\xf3\x44\x41\x41\x55\x09\xad\x2b\xad\xec\x63\x39\x04\xf2\x43\xb7\xb2\xe6\xd2\x98\xf5\xd1\x2e\x0a\x11\xaa\x4c\x1a\xc4\x05\x75\x11\x3d\xd1\x98\x50\x8d\x52\x81\xd4\x07\xe0\x82\x44\xd5\xdb\x9f\x8c\xc1\x1e\xcc\x1f\xc9\x26\x54\xd1\x28\x41\x9b\x69\xda\x30\x1d\x6b\x37\x8d\xc1\x55\xe2\x5b\xb3\xc8\xe4\x88\xed\xc5\x2d\x54\xed\xb4\x24\x6c\x3d\x0e\x7b\xa1\x42\xf5\x4f\x83\xea\x0a\x90\xed\x83\x0a\x4e\x84\x68\x4f\xe0\x6a\xe3\x8a\xbd\x1e\x9a\x8c\xc1\xa5\xe3\xdb\x91\x36\x78\xbf\x49\x1c\xf0\xca\x6d\xc3\xb4\x21\xba\xae\x94\x72\x34\x2e\x54\x2f\x3d\x09\xae\x9a\x02\x8d\xb6\x28\x93\x32\x57\x2f\x26\xee\x44\xb4\xf6\x1d\xdf\xc2\xd6\x3b\x9f\xc0\x39\x4a\xa9\x7c\x80\x10\x05\x55\x5c\x8d\x30\x24\x21\x0f\xfd\x82\xd8\x99\x00\xa8\x1a\x33\xe3\xfe\x17\x09\x05\xaa\xe6\xfe\xd2\x91\x82\xaa\x40\x17\xb8\xd8\x73\x02\xdb\x15\x15\xaa\x03\xab\xc2\x8f\x3f\x0a\x17\xaa\x05\x2b\x93\x6f\x6f\x44\x70\xed\x05\xcf\x50\xc9\x89\xc2\xc8\x85\x0e\xc7\x28\x06\x70\x11\x86\xee\xe6\xb6\x37\x30\xb8\x1a\x43\x0b\xa8\x9a\xe9\xd1\xa7\x15\x70\x59\x86\xb0\xa0\x7b\xea\x78\x42\x80\x77\xd7\x75\xc6\x7c\xd9\x17\x76\x77\x43\xad\x35\xbc\x7b\xe0\x82\xb3\x8b\x1a\x77\x47\x5d\xbc\xc9\x4a\x3c\x0a\x16\xaa\xa7\x70\x96\x73\xa9\xd0\xa2\x0c\xd4\x2c\x86\xe0\x40\x35\x53\x65\xa5\x18\xa1\x46\x59\x29\x63\xa8\x32\x32\x85\xff\xb1\xc2\x19\x5f\x86\xea\x47\xee\x00\xe9\x97\x15\xee\x7b\x16\xc9\x39\xfd\x62\x5f\xd2\x99\x8c\xfd\x7a\xc0\xc3\xf3\xd3\x6c\x6d\x8f\x58\x50\x5e\x4a\xad\x05\x64\xdc\xc1\x70\xb2\xbb\x19\x26\xf1\x23\x41\x94\xf1\x34\x6a\x62\x4e\x76\xd7\x3b\x05\x2f\xec\x4b\x16\x18\x19\x29\x47\xc1\xee\xae\x75\x8a\x52\xae\x4c\x60\x4d\x34\x6a\xe0\x72\xef\xe2\xfe\xe6\xeb\xcd\xc5\xf9\x27\x1f\xf8\xa2\x99\xb1\xbb\xd4\x9e\xad\x2f\x58\x66\x6b\x79\x64\xd2\x8c\x36\xee\x58\x4e\xba\x57\x2c\x5e\x8a\xef\x64\xfa\x7e\xe0\x8a\x25\x90\xe2\xdb\x73\xc5\x62\x55\xe4\x6f\x86\x84\x08\x59\x05\xb2\xf5\x7a\x80\x6c\x24\xd4\x6f\x38\x8d\xc7\xf2\xcf\x59\x43\x4f\x24\xfd\xe6\xde\x2c\xda\x05\xaf\x1e\x96\x4f\x7c\xe9\x3f\xf1\xd4\x1d\x93\x13\x6f\x4c\x86\xae\xbd\x02\x11\xb2\x3d\x89\x11\x34\x53\x44\x20\xb9\x66\x0a\xf7\x95\xe9\x05\x48\xeb\x04\x9e\x8f\xe4\x10\xed\xd1\xd4\xba\x1c\x0a\xac\x56\x51\xa8\xc0\x1b\xa1\x26\x00\x3f\xe3\xcb\x23\xd3\xd4\x9e\x8f\xd4\xce\xca\xbf\x1e\xb4\x7f\xf2\x59\xf8\x6d\xae\xce\xc0\xbd\xf3\x16\xd3\xbb\xa1\x81\x7b\x07\xbd\x38\xd4\xc4\xdb\x94\xc1\x82\x67\x34\x59\x6f\xbc\x20\x76\xf7\x05\xab\xd5\x11\x7b\xcc\x17\xc3\xa1\x08\x53\xef\x36\x75\xda\x7d\x3e\xb6\xdb\xe9\x42\x98\xea\xd7\x5b\x02\xa2\xa7\xde\x8b\x64\xd3\xde\x10\x0c\xd7\xf1\x1c\x27\x0f\xdb\xfb\xf5\xe2\x30\xa6\xdd\x47\xb3\xea\x7e\x6d\xdc\x5f\xb7\xbf\x4e\x87\xa7\x9e\x04\x4e\x87\xc2\xf5\x4e\xc1\xaf\x70\xcd\xbe\x9c\x5f\x5c\x6d\xfe\x0d\x3c\xb3\x4f\x03\x2f\x70\x6d\xdd\x8b\xaa\x73\xdd\x46\xa1\xdd\xd6\x29\xcb\xd5\xed\x0d\x95\x01\x1e\xa0\x63\xf7\xbd\x7f\x1b\x1d\x2b\xba\x5c\x99\xb7\x53\x29\x37\x55\x0d\x7f\xe7\x73\x24\xcb\x64\x85\xb0\xac\x62\x36\x29\x33\x7b\x5c\x15\xbb\x4d\xd9\x32\x94\xa0\xa9\x3a\x75\x86\x31\x7a\x5a\x35\x8e\x8e\x16\x0f\x50\x73\xa2\xf5\x58\x45\xab\x7e\x30\x79\x26\x49\x69\x5f\xbb\x30\xc9\xd8\xf5\x5b\xc5\xbd\xc5\x94\x5d\x93\x76\x41\x58\x93\x97\xd7\x5f\x32\xb9\xcb\x4a\x50\x86\x2d\x8e\x80\x71\xb9\x26\x8f\xd9\x85\xbf\x3e\xda\x47\xb2\xcf\x7d\x72\x07\x5a\xb5\x20\x81\x71\xf7\x55\x75\xbe\xfa\x55\x05\x1f\xaf\xe0\x32\xd8\xac\x29\xd1\x77\x74\xbe\xec\x1a\x29\x7a\xe5\x06\xde\xbe\xee\x3e\x7b\x38\xb8\x78\xf5\x07\x50\x1b\xef\x1b\x4b\xc9\x82\x32\x92\x6e\x04\xc8\xb5\x7b\x82\x1e\x16\xff\x21\x39\x43\x5f\xd7\x05\xe9\x84\xda\x75\x33\x8d\xab\x5b\xaa\x0f\x3c\x5d\x9b\xf6\x3e\xa6\x9f\xfd\xd5\x7b\xe3\x55\xbd\xdb\x69\x23\xec\x8c\x69\xfc\xe5\x6e\xe6\x75\x79\x0c\xbe\xd0\xfa\xc6\x70\xa9\x56\x5c\xd0\x3f\x49\x8a\xbe\x49\xd2\xcf\xc8\xb9\x6b\x67\x73\xf7\xff\x9b\xe0\x94\xf8\xf2\x3b\x06\xfb\xb1\x8d\x3c\x8c\x10\x87\xe5\x67\xda\x7d\xc1\x6b\x33\x9d\xad\x7b\xdf\x47\x05\xfb\x8e\xfe\xe7\xc8\x69\x82\xa3\x9b\xd4\xf1\xb0\x05\x7f\xe3\x8b\x9f\xcb\x92\xa6\xbf\xa0\xef\x38\x2b\xfd\x71\x0c\x44\xb2\xf4\xe5\xd9\xd8\x29\x71\x6e\x5e\x7b\x72\x35\xe1\x0b\x2c\xa5\xd5\xb4\x81\x01\x0d\x3d\xf0\xde\x57\xb5\xc3\x3d\x27\xdb\x53\xac\xdc\x3d\x17\xae\x87\xd0\xfd\xae\xae\xa4\xf2\xe5\x6e\x66\x28\xb1\xf6\x8f\x29\x7d\x36\x53\x38\x09\xf8\xee\x90\x20\x36\x18\x3f\x48\x2a\x78\xee\x05\x46\xbf\xa6\xf3\x16\x67\xae\xbc\xfb\x3f\x64\x28\xa9\x41\xcf\x55\x17\x94\x66\xba\x09\xcf\x8c\x1d\x9e\xe3\xa5\x1f\xf8\x02\x5d\x34\xd2\xe9\xa1\xeb\x2f\x90\x1f\xbc\x08\x2d\xe3\x6a\x65\x9e\xf4\xb3\xf3\xa9\x95\x7c\xda\x9b\x1f\x85\x70\xe7\x23\x63\x29\xd7\xed\xdf\xa0\x5f\x31\x55\x36\x62\x4d\xfd\x5d\x56\xe9\x40\xcd\xdd\x73\x8b\x4e\xb0\x8b\xe4\x5c\x4a\x9e\x50\x53\x02\x5d\xcb\x28\xc1\x59\xd6\xeb\xa6\xae\x1a\x68\xa3\x41\x95\x42\x6b\x6a\xab\xeb\x4c\xc1\x08\x59\x25\xbf\x06\x64\xda\x7d\x9f\x40\xf2\x9c\x20\x45\xbd\x57\x7f\x8e\x8f\xe1\x0a\xf7\x7b\xfb\x81\x8e\xea\xf9\xc1\x8c\xe6\x81\x57\xe1\xed\xf0\x67\x19\x7f\x92\x88\xb3\x6c\x8d\x46\xd3\x77\xcd\x93\x8d\x95\x43\xef\x0d\x2a\x6c\x21\x2f\x5b\x82\x39\xf4\x1c\x61\x97\x27\x97\x52\x6f\x9e\x08\x23\xa2\xca\x6f\xe0\xc2\x52\xd5\x65\x2e\xf4\x6e\x67\x0f\x73\x9c\x99\xb7\xb5\xe8\x07\x7e\x6d\x1e\xa1\x77\x29\xcc\x29\xc1\xdd\xc4\xbf\xe3\xd0\xd3\x88\x3d\xbd\x26\xf6\x89\x10\xae\xec\x6d\xb7\x4d\x6f\xa2\xcd\x43\x51\xad\x3e\xfd\xb7\x2e\x7a\xfa\x34\xa2\xad\x2a\x7a\x01\x1e\xbc\x90\xdd\xe7\x31\xcd\x80\xd8\x69\xa1\xfb\x0a\x58\x79\x8d\xcc\xad\x89\x16\x58\x99\x63\xb8\x12\xbe\x27\xb2\xe0\x4c\xda\xf9\xc7\x4b\xb5\x11\x0e\x0b\x33\x83\x46\xbe\x19\x34\x1a\x34\x83\x46\x6f\xc1\x95\xb9\x2f\x87\xab\x72\xb7\x33\xdb\x75\x4b\x27\xf4\x42\x10\xd9\xf6\x82\x34\x22\x4b\xaa\xd7\x04\x4c\xba\x9a\xfb\xa4\xe9\xc4\x14\x86\x61\xcd\xbb\xa2\x84\xa9\xe6\xdd\xce\x9a\xfc\x1d\xf6\xb8\x7f\x96\x44\xac\xe3\x36\xb9\x1b\xb6\xc8\xca\xe7\xcb\x0f\x66\xe9\x99\x81\x18\x50\xcf\x55\x63\x9f\x54\xf0\x22\xfb\x4a\x73\xd2\x44\xcb\xf5\x11\x5c\x85\xd5\x29\xd3\x9a\x08\xca\xd3\xa6\x04\x4e\x80\xc0\xd2\x4d\x2d\x6b\x01\x1e\xe7\x87\xd3\xfc\xf0\x58\xff\xbf\x3a\x7c\xb7\x3a\x3c\x1e\xad\x0e\x47\x93\xd5\xe1\xfb\xf4\x70\xfc\x36\x6d\xcf\xbd\xcf\xdf\x6f\xaf\x10\x4e\x73\xca\xb6\xe4\x32\xf8\x93\x6f\xfc\xb6\xeb\x42\x68\x8b\xa4\xf9\x60\xbb\x48\x0c\x11\x89\xca\x7a\xa5\x71\x81\xd9\xdf\x6d\xea\x09\x7b\xd4\xda\x57\x89\xac\x6b\x69\x6a\x3c\xe8\x59\xfa\x13\x5f\xa2\x02\x2f\xfb\xf2\x88\x5a\x70\x99\x6b\x19\x00\x83\x1e\x98\x2f\xea\x0d\x73\x18\xce\xbd\x09\xd3\xda\x60\xa5\x37\x50\xe7\xff\xcf\x03\xf5\x83\xc7\xe1\xc7\xca\xb9\x57\x8a\xff\x17\x00\x00\xff\xff\x5a\x04\x9c\x52\x54\x8a\x00\x00")

func ResourcesEventsYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesEventsYaml,
		"../resources/events.yaml",
	)
}

func ResourcesEventsYaml() (*asset, error) {
	bytes, err := ResourcesEventsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../resources/events.yaml", size: 35412, mode: os.FileMode(420), modTime: time.Unix(1622779497, 0)}
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
	"../resources/events.yaml": ResourcesEventsYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"events.yaml": &bintree{ResourcesEventsYaml, map[string]*bintree{}},
		}},
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
