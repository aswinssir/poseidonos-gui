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

var _ResourcesEventsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x7d\xdb\x72\x1b\xb7\xb2\xf6\x75\xfc\x14\x28\x5d\xac\x4a\xaa\xa4\xfc\x16\x0f\xb6\xa4\x3b\x59\x07\xff\x5a\xdb\xb2\xb4\x45\xd9\x6b\xaf\xab\x14\x38\x03\x92\x88\x66\x80\x09\x80\xa1\xc4\xa4\xf2\x2e\x7a\x16\x3d\xd9\x2e\x1c\xe6\x40\x02\x43\x36\x31\xae\xb5\x53\x95\x2a\x4b\x9a\xc1\xd7\xdd\x68\x34\xba\x1b\x8d\x9e\xa3\xa3\xa3\x77\x9f\x39\xce\xce\xd0\xd5\x92\x30\x25\xd1\x37\x46\x67\x34\xc1\x8a\x72\xf6\xee\x3b\x11\x92\x72\x76\x86\x0e\x96\xbf\x7e\xfc\x70\xf0\x2e\xe7\x69\x99\x11\x79\xf6\x0e\xa1\x23\xc4\x70\x4e\xce\xd0\xc1\xc5\xdd\xed\xed\xdd\xd7\x83\x77\x08\x21\x94\xf0\x92\xa9\x33\x74\x7a\x7a\x6a\x7e\xa4\xe9\x44\x61\xa1\xce\xd0\x7b\xf7\xe3\x15\x4b\xcf\x10\x6a\xfe\xce\x66\xfc\xcc\xfc\x4b\x8f\x97\xf0\x94\x54\x8f\xea\xff\x32\xb2\x24\xd9\x19\x3a\xb8\xf9\x7a\x7d\x77\x50\xff\x36\x27\x52\xe2\xb9\x06\x9e\x94\x49\x42\xa4\x6c\xfe\x54\x08\x3e\xcd\x48\x7e\x86\x0e\x9a\xdf\x49\x9e\x95\xca\xb2\x70\xb0\x46\xf5\x97\x9b\x35\x92\x8f\xdf\xbf\x5f\x27\xf9\xf8\xfd\xfb\x0d\xaa\x8f\xdf\x6f\x21\xbb\x7e\xbc\x45\xb9\x26\xdc\xa7\x9b\x4e\xf9\x8c\x4b\x44\x25\x92\x1a\x89\xa4\x3e\xfd\x01\xea\x7d\xb8\xe3\x7d\xe1\x14\x11\x39\x65\x38\x1e\x71\x00\x43\x4c\x32\x8a\x24\x11\x4b\x22\x34\x2a\x65\x54\x51\x9c\xd1\x3f\x03\xb0\x3e\xaa\x0f\x3a\x84\x81\x32\xf2\xac\x81\x09\x53\x1a\x34\xe1\x8c\x91\x24\xc4\x29\x04\x72\x04\xe6\xd3\xc1\xa5\x54\xf6\x43\x1c\xc3\x10\x05\xf9\xa3\x24\x52\xa1\x5c\xce\x35\xac\x20\x09\xa1\xcb\x48\xc8\x0f\x50\x48\x59\x70\x26\x49\x85\x29\x09\x53\x31\x78\xc7\xc0\xd5\xd1\x52\x9e\x42\x90\x02\x0b\xca\xe6\x88\xbc\xd0\x38\x50\xe0\x1a\x69\x81\xaa\x85\x20\x38\x45\xbf\x73\xca\xe2\x04\x7b\x0c\x5c\x25\x84\xe1\x69\x46\x90\x20\xa5\x24\x47\x38\x4d\x05\x8a\x42\xf3\x96\xc7\xbf\xce\x1f\xbe\x42\xd0\x66\x98\x66\x24\x45\x3f\x45\xa1\x7a\x2b\xe4\xea\xe1\xe1\xee\xc1\x87\x95\x3c\x79\x22\x0a\x25\x82\x98\xfd\xc4\x81\xc6\x61\x7a\x6b\x64\x3b\xe6\x94\xb2\x54\xeb\x4e\x1f\x48\x6f\x8d\x6c\x87\xcc\xa8\x54\xa4\x1f\x93\x1f\x81\x88\xa4\xe0\x59\xf6\x63\xe4\x7a\xb2\x1f\x93\x38\x49\x48\xa1\x7a\x21\x9e\x02\x11\x73\xfc\x52\x19\x59\x22\x04\x17\x51\x68\x03\xcf\xf0\x74\xa1\xd9\x7f\x68\x0b\xd7\x57\x71\x06\x9e\xdd\xd9\x81\x69\x2d\x79\x5f\x54\xcf\xf2\x84\x6d\xc1\xe3\xcd\xed\xd5\x25\xba\xfb\xf6\x18\x05\xe2\x19\x9c\x0e\xd6\x6e\xbe\x7e\x3f\xff\x72\x73\x89\xee\xcf\x1f\xce\x6f\x63\x90\x86\xc0\x1d\xe3\xfe\x6e\x82\x6e\x26\xe8\xd3\xb7\xc9\xbf\x61\x30\xb5\x0f\x78\xf3\x89\x5f\xdf\x4d\xd0\x44\x61\x45\xd0\x2d\x66\x78\x4e\xc4\x9a\x53\x38\xf0\x9c\xc2\x63\xcf\x29\x1c\x6c\x73\x0a\x8f\x7d\x2f\xed\xf2\xea\xd3\xb7\xcf\x81\xd5\x65\x88\x48\x38\x53\xe4\x45\x21\x9c\xa6\x51\xbb\xcf\x31\xd4\x47\x73\x70\x0b\xcc\xe6\x51\x40\x83\xa1\xb7\x05\x3c\x5c\xdd\xdf\x3d\x3c\xfe\xf6\xf8\x70\x7e\x71\x15\x70\x43\x9d\xac\x57\x52\x91\x1c\x3d\x90\x84\x2f\x89\x58\xa1\x1b\x56\x08\x3e\x17\x44\xca\x3d\xe7\xee\x3b\xcf\xca\x9c\x84\x26\x6d\xb8\x39\x69\x03\xcf\x93\x1f\x6c\x9b\xb4\x01\xd8\x93\xb7\x34\x58\x2b\x1c\x23\xc5\x01\xd8\x89\x77\x48\x29\xc9\x48\x2c\x12\x50\x31\x1c\x52\xae\x65\x19\x89\x04\xf4\xd8\x1d\x52\xc9\xfa\x60\x01\x5d\x75\x87\x65\x96\x15\x52\x1c\xa9\x05\x31\xbb\x75\x14\x26\xd0\x59\xb7\x98\x6f\xaf\x82\xe4\x7c\x49\x52\x34\x13\x3c\xd7\xc0\x6f\xaf\xb1\xc8\xbe\x0f\xdd\x61\xe0\x17\xa4\x0a\x15\x48\x8a\x96\x4e\x77\x38\x91\x88\x71\xa5\x9d\xe9\x00\xbc\x79\xc9\x3d\xfa\x4c\xd5\xc2\x88\xc8\x1b\x44\x2f\x3e\xc4\x45\xf5\xe3\xcd\x65\xd7\xb0\x4d\x18\x79\xc5\x94\x71\xab\xb5\x71\x13\x82\x24\xaa\x7b\x2c\x3c\xd3\x8f\x26\x0b\x92\x3c\xe9\xad\x50\x35\x14\xb5\x24\xd6\x12\x87\xb7\x78\xc2\xe2\x70\x93\x2f\xe9\x9f\xc4\x84\xc1\x9c\xa3\x0c\x8b\x39\xf1\x65\x50\x2d\x68\xcc\x0c\x4b\xd3\x7a\x6d\xa3\x29\x49\x70\x29\x89\xa1\x29\xc7\x2f\x34\x2f\x73\xc4\xca\x7c\x4a\x04\xe2\x33\x47\xa5\x44\x6a\x81\x95\x79\xbb\xf5\x26\x95\x88\xbc\x24\x84\xb4\x0d\x7a\x23\x9d\x07\xa2\xc4\xca\x31\x6e\x14\x45\x33\x5e\xea\x50\x52\x53\x2f\x56\x95\x00\x70\xce\x6d\x1c\x24\x95\x7e\xa2\x1a\xdc\x01\xfb\xa2\x01\xba\x02\x57\x86\x32\x3b\x9c\x09\x10\x96\x8d\xac\xc2\x2a\xe2\xa9\x84\x11\xab\xe5\x50\xa2\x73\x21\xf0\x0a\x25\xb8\xc0\xc9\x2a\xc0\xed\x85\x9e\x5a\x23\x43\x69\xf7\x01\xf3\x24\x55\x2b\x84\x59\x8a\x8c\x24\xe6\x98\x32\x8f\x1d\xdf\x87\xdb\x3a\xd3\xd5\xfc\x2d\xf0\x92\x20\x8c\xa6\x19\x66\x4f\x46\xdd\x7c\x96\x3e\x99\xbf\x51\xab\xc2\x38\xcb\xf8\xb3\x5e\xa9\x8d\x56\xae\xbd\x05\x55\xea\xd6\xfb\x95\x87\xee\xb3\xb4\x9f\xf2\xea\x71\xdf\x5e\x6b\xf5\xe5\x6c\xde\xa9\xbd\x86\x04\xa2\x09\xb4\xda\xe7\xbd\xb0\xc9\x06\x5e\x23\xde\xe8\x70\x6b\x55\x1b\x05\x79\x79\xd1\xbe\x82\xc0\x89\x22\xc2\xd7\x36\xa8\xe3\xd9\x26\x90\x4a\x94\x96\x45\x46\x93\x70\x3e\xea\x7c\xcd\x16\xe1\xe6\x59\xfb\x36\xce\x74\x94\xbe\xb2\x0b\x42\x6e\xe1\x2c\xa5\xb3\x19\x11\x3a\x82\xf0\x27\xb4\xc5\x00\x30\x8a\xfe\xc6\x6c\xf4\xd3\x5e\x2a\x61\xbd\x0a\xcd\x85\x76\xee\x30\x65\x12\x61\x24\x95\xb0\x16\x0e\xab\x0d\xe5\xdb\xaa\x6c\xde\x3c\xe5\x84\x28\xe9\xfd\x49\x94\x59\xc0\x26\xf8\xfe\x74\x98\xc9\x87\x7a\x85\x9b\xa5\xbd\xc5\x5c\x76\xdb\x83\xa0\x99\x0d\x5b\xbd\x35\x73\x5f\x70\x29\x69\xd8\x0e\xb5\x18\x01\x2e\x9d\x00\x23\x32\xc7\x59\xb6\x3f\x23\x6f\xaf\xeb\x2f\x86\x2c\x5a\x4e\x99\xd9\x15\xf4\x3c\x26\xbe\x35\x35\x16\x4e\x68\xbe\x7d\x7e\x80\xcb\x67\x83\x1f\xab\x34\x74\x1e\xca\x53\x05\x39\x7a\x7b\x35\xef\xd5\x6b\xdb\xbc\x6c\x17\xd8\x34\xe3\xc9\xd3\xba\xd5\x6f\x78\xbc\x61\x45\xa9\xd6\x78\x51\x5c\xef\x70\x79\x99\x29\x5a\x64\x44\xef\x80\xde\x00\x0d\x7b\xa3\xfd\x8c\x77\xb5\xb4\xbb\xbc\x42\x74\xae\x14\xc9\x0b\xa5\x89\x30\xcf\x34\xf6\xab\x5a\x4e\x5d\x43\x34\x2c\x7d\xe5\x6a\x61\x74\x8e\xa3\x94\xfb\x14\xef\x67\x9b\x2b\xb8\x6e\x4f\xb6\x4d\xb3\x7b\xaa\x9b\x6a\x7f\x18\x28\xdd\x63\x3f\x6e\x61\x33\xee\xd3\x7d\xe1\xf6\x47\x1b\x4e\x54\x62\x72\xf4\x74\x12\x6f\x3d\x67\xf7\x0e\xee\x7a\xab\xa1\xb5\xc5\xb3\x7b\xa9\xf2\x65\xcc\xba\x77\x6c\xae\x3b\x7a\x1e\x4b\x1f\xfc\xec\x8d\x10\x5c\xf8\x3c\x5d\x63\x9a\x69\x28\xeb\x19\x55\x50\x39\x51\x38\xc6\xd3\xfe\xe0\x27\x70\xb6\xc3\xf2\x82\xb0\xde\xa0\x9e\x1d\xd8\x01\x6a\x32\xd5\x7d\x41\xfd\x7c\xce\x76\xd0\x67\x41\x7f\x80\x7c\xfd\x04\x72\x18\xf5\x7b\x83\xf3\xf6\x6a\xb2\x23\x4c\xa1\xa9\xe0\x4f\x84\xc5\xe0\x7e\x84\xaa\xd3\x2d\xc9\xb9\xde\xa2\xac\x31\xa7\x9c\xbd\xbd\xce\x30\xcd\x4a\x11\x5a\x1f\x28\xc1\xd2\xad\x63\xb9\xe0\x65\x96\x22\x46\x96\x3a\x32\x48\x92\x52\xa0\x23\xb4\x20\xb8\x68\x0d\x85\x36\x47\x6a\xd6\xcc\x63\xa7\x0b\xfc\x11\xaa\x91\x37\x6c\x89\x33\x9a\x22\xca\x52\xf2\x62\xf3\xa6\x11\x24\x9b\xb7\x7f\x76\xb3\x4c\xd3\x5f\x10\xd5\x4e\x08\xc3\x59\xb6\x42\x73\x81\x99\x8b\x6c\xa8\x05\x0b\x6e\x1a\xf6\x79\x94\xf1\x39\x4d\xde\x5e\xdb\x84\xb4\xb8\xda\x57\xe5\x8d\x14\x7f\x7d\x7b\x65\xe4\xf9\xed\xb5\x8e\x1c\x23\x18\xfc\x66\x0f\x44\x14\x47\x73\xba\x24\x4d\x10\x7a\x88\x52\x22\x0b\xad\xe2\x2d\xaf\xca\x64\x96\x2a\x47\x2d\xc7\x2f\xf1\xfc\x42\x57\x9b\xde\xbf\x25\xb6\x61\xb1\x23\x62\xc3\xd7\x85\xb3\x7a\xee\xc5\xf5\xd5\xc8\x3b\xdd\x68\x18\x57\x27\xfe\x0e\x8f\x05\xa3\xad\x90\xa3\xe6\xeb\xae\x54\xda\x5d\xf8\x83\x4b\x24\x30\x0b\x39\x95\xe7\x68\x89\xb3\x92\xa0\x8c\x48\x13\x50\xb3\x75\xef\xaa\x30\x71\x80\x9e\x3a\x3d\x86\x7d\xf4\x19\xcb\xca\xc9\x06\xb9\x68\xcd\x9b\xed\x80\xbd\x72\xd3\xd7\x42\xd1\x5f\x3d\x66\x4f\xe1\xcc\xda\x68\xd8\xe8\x4c\x97\x63\xd0\x44\xac\x1b\x09\x07\x2e\x50\xc6\x71\x4a\x52\x33\x6b\xbc\x54\xd5\x41\x7e\xb7\x73\x50\x1b\x0f\xb7\xc3\x5a\x3f\xc3\xbe\xe6\xb3\xe1\xfb\x38\x5d\x6c\xb8\x9c\xed\x35\x2e\xb3\x40\xde\xa8\xe2\x80\xe7\xb9\x96\x5c\xc3\x49\x41\xc4\x8c\x8b\x5c\x1b\x0a\x3b\x87\x93\xc7\xbb\x7b\x9b\x78\x06\x58\xea\x53\xdf\x29\xee\xa2\xef\x92\x33\xa7\xdb\x1d\xd6\x6e\xc2\xf5\xf2\xd1\x7f\x93\x28\xc7\x2b\xb7\x30\xd2\x52\xd4\x61\x87\xe0\x09\x91\x52\xff\xc8\x67\xed\xcc\xd7\xa1\xd5\x06\xbd\x64\xca\xa9\xd4\xbf\x63\x4a\xef\xf5\xc2\x1a\xf2\xdc\x4d\xef\x33\x17\x4f\xe8\x99\x64\xd9\xaf\xa1\xf0\x4d\x03\x9b\x9c\x80\x21\x01\x2d\x30\x4b\x33\x0d\x85\x33\x3d\xb1\xf3\x05\xa2\xaa\x12\x9b\xa5\xcc\xf0\x52\x4a\x22\x90\x85\x4c\x3c\xbf\xe8\xd8\x4f\x51\x77\x8a\x47\x53\x6e\x67\x50\x23\x74\x25\xab\x7d\xb7\xba\x6d\x34\x18\x6f\x0d\xb3\x39\x44\xc3\xea\xa4\x7a\x46\xa2\xbc\x94\x1b\x09\xb4\x19\x17\xce\xd7\xd4\xcc\x37\x99\xab\x3a\xa5\x7f\xfb\xe9\x21\x78\x0a\xe3\x95\xe6\x0c\x86\x5e\x42\x7f\xb8\x2d\xa1\x3f\x84\x26\xf4\xcf\xed\x1e\xad\xfd\x0d\xb3\xfd\x4b\x5b\x61\x14\xe1\x69\x0c\xa1\x99\xfd\x7c\x2a\x50\x8a\x15\xae\x1c\x57\xad\x6f\xc6\xc5\x8a\x02\x05\x26\xf9\x6b\x50\x9c\xa6\x3d\x11\x81\xc9\xfe\x02\x0b\xaa\x56\x36\xd8\xef\x25\x56\x60\xc2\xdf\xa9\x6a\x59\xd2\xb4\x3f\x28\x34\xef\x9e\x92\x25\x4d\x6c\x60\x3e\xe3\x25\x8b\x39\xd1\x18\x42\x93\xda\x6b\x02\xd5\x0e\x65\x14\x18\x30\xf3\xe0\x4b\x33\x1a\x11\x98\x69\xab\x35\xb4\x9f\x30\x3d\x6d\xe9\x46\x7b\x22\xab\x7e\x60\x7e\x7e\xad\xe3\x64\x7c\xd3\xca\xc4\xca\xd2\xcf\x83\x75\x55\x19\x54\xc2\xb4\xb9\xfa\xe0\x26\x09\xc1\xf3\x9d\xf5\x6e\x3c\xbd\xf5\x63\xd5\x03\x0c\x5a\x67\x30\x27\x6a\x4d\x3f\xe1\xd2\xac\x37\x9e\x4b\xb3\x6e\x61\x5b\xcf\xc8\xdb\x7a\x46\xdb\xb6\x9e\x11\x74\xeb\xb9\xb1\x65\x92\x48\x26\x38\x26\xb2\x1d\x41\xf7\x9b\x07\x12\x8d\x00\xdc\x5c\xbe\x50\xa9\x9c\x29\x8c\x82\x01\xee\x28\x06\x06\x6b\x0f\xfb\xed\xb5\x07\x5a\x70\x3f\xf9\x7f\x61\x33\x71\x4d\x59\xaa\x39\xfb\xb9\xa4\xe9\x2f\x51\x68\xe0\x1a\x36\x1d\x93\xc4\xae\x9d\x91\x5f\xdb\xd9\x01\xe3\x36\x2c\x93\xbc\x8a\x47\x83\xd6\xac\x39\xb4\x24\xe3\x32\xde\x0a\x8d\xde\x7b\xf5\x6a\x5b\x37\xe3\xb7\xd7\x6e\xa3\x8e\xec\xc2\x6f\x25\xca\x6d\x62\xbb\x3a\xb1\xd7\x71\xa5\xff\xf6\x16\xd2\xbc\xc2\xb6\x30\x69\x5f\x39\x33\x10\x26\x8c\xbd\x0c\xeb\x2e\xba\x63\xd9\x0a\xa5\x44\xe1\x64\x41\xd2\x5a\x76\x36\x54\x75\x04\x02\x48\x82\x96\xdd\xa6\x98\xe4\x9c\xd9\x8a\xf4\x98\x59\x81\x56\xda\xd6\x38\xbc\x88\x82\x01\xda\x20\x27\x2e\xac\xb4\xf8\x72\x1d\xb4\xa5\x44\x45\x96\x67\x0f\x46\xbe\xd3\xb2\x0d\x55\x9b\xa3\x1f\x03\x0b\xf4\x75\x2d\x97\x3d\x4c\xae\x5f\x5a\xdb\xc5\x5f\x5f\x20\xcf\x30\x85\x17\xc8\x1a\x47\xf1\xa6\xc2\xaf\xa6\xed\x32\x15\x3f\x06\x0e\x68\x99\xca\x3a\xf5\x68\x96\x1b\xca\x39\xa3\x8a\x0b\x1a\x3a\xd3\xc7\x59\xd6\xfa\xbb\x5b\x3e\x12\x61\x51\x1f\x3b\xbd\xbd\x8a\x92\xe9\xc8\xff\x10\x71\xa1\x03\x75\xf7\xb8\xec\xaa\xc8\xf1\xe9\x06\x9a\xad\x36\xdd\xbc\xd8\x45\xf6\xdb\xab\x4f\xf7\xdb\x6b\x8b\x70\x33\x4a\x41\xd2\x68\xba\xc1\x95\xbd\x2d\xa3\x53\x90\x14\x95\x8c\xbc\x14\x66\x55\x66\x2b\x9f\x74\xc8\xc3\x5b\x68\x82\xd6\x7b\xda\xcb\x0d\xa4\xc7\x62\xf2\xab\x2e\x3a\xa0\x04\xc9\x08\x96\xfb\x41\xd5\xbe\xb1\xc9\x61\xee\xa8\xaf\x1c\x7b\x3e\xf1\xc7\x6d\x3e\xf1\xd8\xf7\x89\xbb\x22\x25\x93\x40\x75\x67\xa3\x29\xae\x4f\x74\x7f\x8d\x90\xd7\xd8\xf7\x90\x21\xa8\xe6\x44\xb6\x17\x2c\x34\x6a\xb2\xb0\x1d\x29\xf9\x7d\x00\xa1\x91\x93\x05\x5c\x2f\xe0\x8b\x43\x84\x5e\x07\xb9\xec\x9f\x20\x19\x83\x7d\x68\xb5\x20\xad\x12\x3d\x29\xd3\x26\x29\x39\x25\xea\x99\x10\x86\xfe\xfa\xdb\x64\xbf\xfe\xfa\x3b\x8a\x10\xa8\x97\xcd\x96\x79\xb5\xb7\xf4\x0a\xc0\xc7\x60\x4f\x7b\x22\x5d\x8d\x48\xc5\xf0\xdc\x24\x17\x85\x3d\x4f\xf9\xeb\x6f\x34\x5d\x29\x12\x37\xd5\xd0\xdb\x22\x93\xb6\xbc\xeb\x33\x27\x4d\x54\x1c\x2e\xf4\xce\xc8\xed\xa7\x87\xb7\x57\x73\x30\x1e\x2d\x66\xdf\x67\xee\xc6\x72\xc7\xe1\xf1\x58\x50\xa3\x64\x4e\x5f\x8e\x14\xcf\x88\xc0\x2c\x21\xc6\xb2\xa2\x9e\x7c\x42\x2d\xd3\xbf\x04\x67\x73\x8f\x82\x9c\xa8\x05\x4f\x91\x5a\x15\x31\xdb\xd7\xd8\x77\xaa\x3b\xd0\x0f\xfe\xfa\x1b\xdd\x63\xa1\xa8\x39\x72\xa9\xaf\x3c\x19\xb6\xfd\xbb\xbf\x10\x64\xa8\xb9\x6a\x90\x39\x33\xa7\x71\x7d\x40\xa1\x66\xcb\x08\xfb\xed\xd5\x1a\x67\xb2\x34\x37\x4c\xa3\x0c\x25\xf8\xf6\x9a\x3b\xc4\xb7\x67\xbc\x38\x43\x38\x4d\x05\x91\xb2\x87\x62\x41\x8d\x84\x55\x2c\xad\x41\xe6\x7c\x0c\xdb\x0d\xa9\x23\x20\x7e\x5c\x10\xfb\xe8\xcf\xfa\xd9\x69\x39\x9b\x69\xcb\x6e\x4f\xd6\x52\xac\xf0\x91\x54\x5c\xe0\x39\xf9\x05\xc9\x82\x24\x74\x46\x49\x8a\xa6\x2b\x63\x7b\xda\x03\xd7\x27\x7a\x38\x51\x25\xce\xaa\xdf\x9a\x91\x8d\x43\x56\x55\x8c\x86\xce\xf2\x9a\xb3\x64\xfb\x7c\x57\x4d\xdf\xd8\xf7\x4d\xc3\x21\x95\x39\x0c\xac\x5c\x33\xd4\x04\x8e\x31\x62\x87\x96\x14\x33\x6d\x38\x72\x4c\x75\xd0\x80\x64\x3f\x2f\x74\xec\x7b\xa1\x1d\x73\xed\xee\x9b\xea\xe8\xe1\x19\x17\x0e\xd7\x1c\x2e\xea\xc9\x8b\x81\x86\x5e\x10\xf3\x34\x60\x81\x25\x9a\xea\xfd\x3f\xf2\x9a\xd5\x60\xec\xa7\xde\xf7\x85\xde\x4c\xe1\xec\x03\x0e\x3c\xb5\x69\x24\x8e\xd3\xd4\x23\xc5\x5f\x60\x1e\xb1\x71\xb5\xd7\x3e\xbd\xc0\x33\x9f\x86\x5e\x2b\x1c\x00\xc9\xda\x26\xb4\x89\x5d\xbf\xb4\x52\x15\x9a\x6b\x7f\xd3\xd5\x20\x64\x54\x9a\x02\x94\xf6\xb0\x81\x85\xee\xf3\x00\xcd\xc5\xb4\xc6\xb5\x61\x7a\x92\x11\x2c\x22\x27\x1a\xba\x67\xac\xcf\x34\x17\x95\x00\xbb\xc4\x86\xeb\xc2\x94\xce\xb3\x7c\x9f\x18\xe8\x76\xb2\x2f\x31\x25\x7b\x62\xfc\x99\xad\xcd\xe3\xb3\xd9\x18\xc2\x97\x2f\x7c\xca\xa0\x6e\xf1\x3a\x65\xdb\x96\x41\xa3\x43\x8e\xaa\x56\x34\xe8\x34\xc9\xc8\x10\x40\x5d\xa0\x28\x77\x8b\xe6\x48\x85\x55\x29\x51\x59\xa4\x91\x57\x0a\x3f\xf8\xc1\xf5\xc5\xc3\xcd\xe3\xcd\xc5\xf9\x97\x80\x93\xfe\xef\xc9\xe3\xd5\xad\x29\xcc\xb9\xbf\x0a\x54\xd1\xcd\xb4\xcb\x87\x1a\x97\xaf\xba\xc8\x84\xd2\x92\xd8\xba\x5e\x23\x9c\xee\x7a\x46\x41\x8a\x0c\x9b\x27\xf4\x40\xf5\xb2\x30\xdb\xe6\x51\x55\xf1\xe0\x04\x59\xa7\x1c\x1e\xc8\xb4\xa4\x59\xba\xe3\x20\xee\xc4\x4b\x3a\x9c\x6c\x4b\x3a\x9c\xc0\x7b\x6c\x18\x74\xb4\xdf\xbd\xd4\x0d\x2c\x60\x73\x0d\xe7\xa8\x24\xe6\x3e\x48\x61\xb4\x2c\xe3\xc9\x53\xd4\xc4\x9f\xf8\xd1\x21\x08\xf4\xed\x95\x4a\x54\xb2\x78\x5c\xe8\x11\x43\x25\xd8\x84\xe7\x85\x29\x07\x77\x15\x19\xb3\x32\x0b\xe4\xe1\x20\xc0\xd0\xd8\xe9\xa0\x82\x16\x44\x96\x99\x6a\x48\x71\x8a\x1b\xe3\xd0\x9f\x40\x6b\x27\x3a\xc1\x5d\x86\x31\x0e\x1c\xb8\xa5\x1a\x57\xcf\xc2\x29\x2c\xe6\x24\x50\xbf\xa7\x16\x1b\xee\xa7\xad\xec\xb5\x45\x7c\x19\xb7\x97\x0d\x30\x5b\xa1\xa2\x8a\xc1\x20\xf4\x01\xb7\x4b\xc6\x11\x23\xd5\x9e\x6f\xe8\xf4\x09\xb4\x57\xd5\x1d\x49\x29\x99\x0b\x1c\xbc\x4a\xe9\xd3\x00\xdd\x3d\xaf\x1b\xcf\xd4\xe4\xe3\x1d\x25\x87\x55\x59\x81\xe9\xd4\xc4\x8b\x8d\xd2\x45\x58\xfa\xf4\x82\xb3\x19\x9d\xc3\x4a\x0b\x4e\x3d\x8b\xb6\xad\x4f\x56\xfd\xf8\x4e\x21\x27\x86\x06\x34\xa3\xa6\x01\x0c\x4e\x51\xca\x59\x8c\xaf\x7f\x0a\xad\x31\x98\x13\xe5\xea\x6b\xa3\x91\x80\xe7\x7c\x36\x70\xb5\x58\xf1\x05\x5e\xa7\xd0\x9a\x83\x36\x5c\x64\xd5\xce\x69\xa0\x98\xac\xab\x72\xb3\x6a\xf7\xe4\x26\xd0\x84\x9c\xb1\x41\xf9\x69\xa0\xb0\x6c\x17\xee\x13\x59\xc5\xe3\x05\x4e\x67\xc2\xa5\xee\x6d\xf5\x34\x95\x08\xb1\x92\x0d\x74\x5d\xd9\x89\xe8\xee\xc4\xc5\x73\x09\xbd\xc0\xb0\x86\xd9\x2b\x6f\x77\x1a\x68\xc1\xb2\x5b\xb2\x3d\xeb\xb1\x4e\x07\xd0\x1b\x3b\xbf\x4b\xce\x50\xca\x93\xca\x62\xf3\xe9\xef\x24\x09\xec\x3b\x95\x92\x35\xc6\xe2\x1f\x76\x7d\x99\xeb\x5b\xe6\x17\xb6\x8c\x7d\x5d\x31\x0e\x5b\x49\x4f\xf4\x8f\x30\x5b\x41\x4b\x7c\x8b\x8b\x62\x57\x7b\x97\xa1\xd7\x29\x64\x78\xbc\x61\x82\x9b\x83\x31\x5b\xc5\xc7\x77\x0d\x39\xf0\x86\xdc\xac\x55\xae\x87\xfc\x7c\xb1\x7d\x87\x18\x7a\xc5\x67\xc3\x6d\xc5\x67\x43\x70\xf1\xd9\xe7\x0b\xa4\x04\x9d\xcf\x49\x54\xa0\x3a\x04\x17\x9f\x7d\xbe\xe8\xec\x79\x08\x41\x01\x6e\x0a\x9f\x2f\xde\x5e\xe3\x36\x9e\x21\xb8\xfa\xec\xf3\x45\xdd\x65\x20\xb2\x66\x66\x08\xae\xcd\xf9\x4e\x13\x45\xf3\xda\x5d\x4f\x38\x93\x4a\x94\x89\x8a\x59\xc8\x43\x70\xa5\xce\x17\x8e\xb5\xcf\xba\x24\x42\x12\x94\xe3\x88\x72\x9d\x21\xb8\x5c\xc7\x60\xd9\xcd\xd5\x5c\x24\xde\xb7\x19\xd0\x2d\x51\xf8\x7a\xb2\xb6\x78\xc6\x9b\x8b\x67\xe4\xad\xed\xd1\xb6\xc5\x33\x0a\x74\x01\xea\x38\x5f\xbf\xb5\xff\x68\xf2\xd8\x29\x99\x96\xf3\x79\xa8\x30\x62\xa7\xcc\x46\x81\x9b\x1d\x61\x99\x79\xa8\x9a\x0b\x6d\x0c\xdb\x1e\x3a\x1c\x77\x10\xb8\x51\x12\x0c\x29\x6e\xea\x7e\x9e\xe8\xf6\x7a\xb2\x76\x5b\x71\x3f\x3c\x60\xb6\xfa\xd6\x5c\x46\xe9\x07\x05\x0c\xd5\xbe\xb9\xab\xd9\x7d\xc0\x02\x77\x4d\xba\x4e\x0f\xaf\x27\x7a\x8f\x8c\xee\x7e\x34\x0a\xdc\x31\xd9\x0e\xd5\xa3\xd5\xd2\x28\x70\xb7\xa4\x0b\x6c\x53\x31\x23\x7d\x8e\x51\xe0\x72\x49\x57\xaf\x39\x77\x82\x55\x60\x81\x73\xd2\xea\x51\xb2\x1f\x1c\xf4\x5c\x50\x5b\x7d\x93\x0b\x8c\x42\x81\x86\xa5\x55\xbe\x31\x1e\x69\xff\x8c\x6d\xd3\xac\x37\x0a\x70\xff\x44\xec\xd4\x94\x87\x95\x11\xbb\xcb\x68\x08\xae\x45\x60\x1c\xe5\x24\xa5\xd8\x68\xe3\xed\xf5\x24\x0a\x0c\x5a\x80\xe0\xba\x74\xac\xf5\x41\x99\xd2\x98\xad\x20\x70\xef\x68\xa7\x3c\xb5\x9f\x1c\x05\x05\x35\x25\xb4\xbe\x5e\xde\x67\xb7\x09\xdc\x3c\xda\xc9\x9a\x4b\x17\xeb\x70\x20\x0a\x11\x6a\x4c\x1a\xc4\x19\x75\x47\xc5\xd1\x98\x50\x8b\x52\x81\xd4\x41\x53\x41\xa2\xda\x0f\x8f\x86\xe0\xac\xd7\x8f\x64\x13\x6a\x68\x94\xa0\x8d\x9a\x36\x4c\x47\xef\xb5\xe0\x76\xb9\x2d\x2d\x32\xb7\x0f\x7a\x71\x0b\x35\x3b\x2d\x09\xdb\x28\xb5\x17\x2a\xd4\xfe\x34\xa8\xae\xeb\x4a\x1f\x54\x70\x89\x6d\x5b\x81\xab\x8d\x2b\xf6\x48\x61\x34\x04\xf7\xcf\x6d\x1f\x19\xe3\x7e\x4a\x1c\x51\xc8\x60\x6b\xbf\x5c\x3b\xc9\x68\x5c\xa8\x5d\x7a\x16\x5c\x35\x5d\xa9\x6c\x27\x0a\x65\xd2\xf5\xe6\xc8\x54\xb4\xf6\x1d\x8f\x0e\xb3\xf3\x09\x9c\xa3\x94\xca\x27\x08\x51\x50\xc3\xd5\xaa\xea\x20\xe4\xa9\x5b\x10\x7b\x13\x00\x35\x63\x66\xde\xff\x43\x42\x81\x9a\xb9\xff\xe8\x4c\x41\x4d\xa0\x0d\x6d\xec\x7c\xf5\x47\x85\xda\xc0\xaa\xdb\xd5\x8f\xc2\x85\x5a\xc1\xca\xe5\xeb\x8d\x08\xbe\xd5\xeb\x39\x2a\x39\x51\x18\xb9\x9a\xb4\x18\xc3\x00\xbe\xde\xbb\xb9\xb9\xf5\x06\x06\xdf\xf3\x6d\x01\x55\x9a\x1e\x1d\xad\x80\x2f\xfc\x86\x05\xdd\xd1\xbc\x0c\x02\xbc\xbf\xad\x33\xee\x4b\x5f\xd8\xfd\x1d\xb5\xd6\xf4\xf6\xc0\x05\x97\xad\x73\x85\x08\x33\x0d\x4a\x64\x81\x13\xd2\x48\x3c\x0a\x16\x6a\xa7\x70\x96\x73\xa9\xd0\xac\x0c\x34\x6a\x84\xe0\x40\x2d\x53\xe5\xa5\x18\xa1\x46\x79\x29\x43\xa8\x31\x32\x7d\x62\xb1\xc2\x19\x9f\x87\x9a\x66\xed\x01\xe9\xf7\x52\xec\xfa\x3e\x04\xfb\xa3\x24\xee\x94\x32\x8e\x3b\xbf\x09\xe2\x76\xfd\x34\x5b\xdb\x12\x0b\xca\x4b\xa9\xad\x80\x8c\x0b\x0c\x47\x11\xf5\xa4\x78\x49\x10\x65\x3c\x8d\x52\xcc\xd1\xfe\x76\xa7\xe0\x85\xed\xe6\x8d\x91\x91\x72\x14\xec\xfe\x56\xa7\x28\xe5\xc2\x14\x63\x44\xa3\x06\x0e\x84\xba\x8b\xc3\x2e\x1a\x8d\xdd\xa7\xe1\x5e\x9d\x94\x9f\xac\xe4\x91\xa9\x5f\x5f\xcb\xcb\x9f\x6c\xa6\xe5\xbd\xcb\x63\xa3\xf1\xc7\x2d\x69\xf9\xc0\xe5\xb1\x8e\xb4\xbc\x35\x91\xbf\x19\x12\x22\x64\x15\xb8\x06\xd2\x01\x64\xab\x67\x7e\xc3\x69\x3c\x96\x1f\x67\x6d\xfb\x4c\xc4\x6f\xee\xbb\x0d\xfb\xe0\xd5\xd3\xf2\x85\xcf\xfd\xcf\x5c\x6c\xce\xc9\x89\x37\x27\xdb\x8e\x4a\xc6\x81\x06\x70\xe1\x0a\x5f\x9a\x29\x22\x90\x5c\x31\x85\xbb\x7a\x13\x02\xa4\x75\x02\x2f\x74\x77\x88\x36\x34\xb5\x29\x87\x02\xab\x45\x14\x2a\xf0\x14\xa1\xa9\x1d\xcd\xf8\xfc\xc8\x3c\x6a\xe3\x23\xb5\xb7\xf1\x6f\xbe\x54\x72\x77\x8f\xd5\xe2\x88\x2d\xf3\xd9\xf6\x23\xe2\xb1\x77\xca\x35\xde\xfc\x6a\xdd\xe6\xa0\x33\x61\xda\x6a\xee\xa8\xb7\x1c\x7b\x5f\x3e\x19\x77\x1e\x8d\xbb\x81\xa7\x38\x79\xda\x3d\xae\x77\x3e\x3e\xde\xfc\x38\x47\x3d\xae\xad\xc7\xda\x1c\x6f\x63\xc0\x53\x4f\x02\xa7\xdb\xca\xa8\x4e\xc1\x5f\xfb\x98\xdc\x9f\x5f\x5c\xad\xff\x0d\xac\x3d\xa7\x81\x2f\x7d\xec\xb4\xf7\x55\xec\xb4\xd6\xc1\xaf\x15\xc9\xb8\x86\x80\xa1\xfe\x82\x5b\xe8\xd8\x7f\x7f\xdd\x45\xc7\x82\xce\x17\xe6\x73\x6d\x94\x9b\x9e\x54\xbf\xf3\x29\x92\x65\xb2\x40\x58\x56\xb5\x74\x94\x99\x7d\xa4\xaa\xa9\xa5\x6c\x1e\xba\x5d\xa3\x36\x1a\x18\x62\xf4\xbc\x68\x92\x09\x2d\x1e\xa0\x5b\x76\xab\x0b\x76\xab\x31\x21\x79\x21\x49\x69\xdb\x68\x9b\x9b\x74\xf5\x07\x12\x3b\xbb\x34\xba\x47\x2a\x97\xb1\xbe\xc4\xd1\xdd\x8b\x71\x93\x15\x4f\x86\xad\xba\x14\x86\xe7\xc4\x34\x8b\x98\x28\x9c\x3c\xfd\x7c\x7b\xfa\x5f\xbf\x78\x3a\x1e\xf8\x38\xe5\x66\xb1\x60\xf3\xab\xe0\xe7\x29\x8f\xc1\xa7\x09\xdf\x18\x2e\xd5\x82\x0b\xfa\x27\x49\xd1\x37\x49\x02\x1d\x15\xed\x3d\xa8\x73\xf7\x98\xbd\x90\xf7\xff\x09\x4e\x89\xf8\xc9\xff\x6e\xe4\x31\x38\x8b\xf8\x89\xa7\x2b\xf4\x4f\xc9\x99\x2d\x2e\xf2\xbf\x82\xe5\x80\xcd\x73\xf7\x78\x65\xee\xe4\xd9\xe4\xea\x4f\x3e\x2c\x38\x74\xff\x9f\x23\xa7\x24\x47\x37\xa9\x63\x63\x07\x01\x6b\x6f\xfc\x5c\x96\x34\xfd\x05\x7d\xc7\x59\x49\x3c\x32\x42\xdf\x05\xed\xba\x32\xeb\x3e\x3c\xd6\xd1\xc5\xd2\x7d\x61\x52\x8b\xda\xfd\xae\xbe\xc7\x7c\x7f\x37\x31\x45\xf5\x9b\xba\x14\x50\x4c\x41\x6c\x91\xeb\xfd\xdd\xc4\x27\x15\xac\x22\xcd\x1c\xf9\x74\xde\x62\xd7\xf6\xf3\x9f\x32\x54\x2b\xac\x35\xca\xd5\x7a\x98\x51\xec\xf4\xf9\xb4\x80\x95\xe6\xa2\x91\xcb\xff\x95\xe4\x42\x9f\xe6\xe9\xa0\x96\x33\xd3\x2f\x9f\x7e\xe2\xd7\xe6\x93\x93\xee\x72\x49\x4a\x70\xea\x8f\x0a\x96\x41\x62\xdb\xfe\x72\x65\x0f\x73\x6c\xc5\x37\x6d\x9a\xbf\xd7\x63\x0e\xe1\xea\xf8\x50\x7d\xf9\x53\xd1\x9c\xf0\x52\xad\x19\xad\xf3\xf9\xa6\x63\x1f\xb0\x52\x03\xdf\x4a\x0d\xb7\x5a\xa9\xc1\x7b\x70\x83\xc0\xcb\xed\xcd\x01\xdb\x37\x78\xf4\x93\x2e\x3b\x55\x08\xb2\xf6\x01\xd3\x66\x76\x6d\x5b\x43\x6a\xbb\xbf\x56\xaf\x34\x83\x98\xeb\xa9\xac\xf9\x00\x0f\x61\x8a\x2a\xef\xda\xe7\x1e\x8b\xfd\xbf\x4b\x22\x56\x71\xab\xfd\x86\xcd\xb2\xf2\xe5\xf2\x93\x51\x5c\x33\x11\x5b\xb4\xb5\x7a\xd8\x27\x15\xac\xb2\x8f\x34\x27\x4d\x69\x45\x17\xc1\x55\x0d\x86\x32\x4f\x13\x41\x79\xda\x5c\xc4\x0d\x10\x58\x3a\xcd\xb2\x91\xf4\x71\x7e\x38\xce\x0f\x8f\xf5\xff\x8b\xc3\x0f\x8b\xc3\xe3\xc1\xe2\x70\x30\x5a\x1c\x7e\x4c\x0f\x87\xef\xd3\x77\xff\x1b\x00\x00\xff\xff\x3f\x47\x1e\x88\xa5\x7a\x00\x00")

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

	info := bindataFileInfo{name: "../resources/events.yaml", size: 31397, mode: os.FileMode(420), modTime: time.Unix(1598943413, 0)}
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
