// Code generated by go-bindata DO NOT EDIT.
// sources:
// completion.sh
package completion

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

var _completionSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xfb\x6f\x1b\xb9\xf1\xff\x5d\x7f\xc5\xdc\x4a\x80\x1f\xf1\x9e\xed\xfb\x7e\x7f\xa9\x1d\x05\xb9\xe6\x51\x04\xb8\xbb\x04\x4d\xef\x87\xc2\x35\x04\x6a\x77\x56\xcb\x9a\x22\x17\x7c\x48\x51\x5d\xff\xef\xc5\x90\xdc\x97\xb4\x76\x92\xd6\x89\x93\xd6\x0b\x1c\x62\x71\x39\xc3\xe1\x3c\x3e\x33\xe4\xcd\x8e\x7f\x38\x9e\x73\x79\x3c\x67\xa6\x1c\x8d\xc6\x30\x9b\xe5\x99\x32\xb3\xa7\x25\x8a\x0a\x35\x14\x4e\x66\xcf\x68\x38\x8c\x66\x82\x83\x71\xf3\x4c\x2d\x97\x4c\xe6\xcf\x46\xa3\x38\x3d\xc7\xb9\x5b\xec\x1f\xc0\xf5\x08\x00\x80\x17\x70\x71\x01\xa9\x84\xc9\xf5\xcb\x17\x6f\xdf\xcf\x5e\xbc\xfd\xf5\xdd\xec\xe5\xab\x3f\xfe\xfe\xa7\xd9\xeb\x37\xbf\xbc\xba\x81\xcb\xcb\x73\xb0\x25\x4a\x3f\x9b\x1e\xcc\x4a\x05\xc9\xe4\xfa\xf5\xef\xbf\xbd\xf8\xed\xe7\x5f\x5f\x5d\x9c\x5e\xde\x9c\xc1\xe4\x30\x81\x67\xcf\x68\x7c\x90\x4d\xe2\xc9\x0b\x3e\xba\x19\x91\xe4\x2f\xb1\x60\x4e\x58\x98\x63\xc9\x56\x5c\x69\xb0\x0a\x16\x68\x69\x21\x90\xf8\xc1\x42\x94\x1a\x0a\xad\x96\x7e\x34\x73\x5a\xa3\x6c\x5e\xfc\x08\x6f\x64\x18\x67\x06\x41\x15\x90\xa6\xa4\x05\x20\x9a\x25\xdb\xcc\x11\x94\x2d\x51\x83\xe1\xd6\x31\xcb\x95\x34\xa3\x31\xf0\x40\x52\x38\xeb\x34\x82\x2d\x99\x05\x53\x2a\x27\x72\x40\x99\x13\xe7\x4a\x20\xcd\x3d\x02\x5b\x72\x03\x1a\xad\xd3\xd2\xc0\xe9\x51\x60\xb6\xe6\x06\xe1\x64\x34\x1e\x8d\xe1\x67\x63\xdc\x12\x0d\x30\x98\xd4\x92\xae\x98\xe6\x6c\x2e\x10\xf0\x03\x37\xd6\xd4\x8b\x65\x4c\x08\x2e\x17\x90\x29\x69\x69\x63\x7e\xd5\x35\x17\xc2\x8f\xb0\x5a\x24\xe5\x64\xde\xb1\xd6\xd6\x12\x9f\xc5\x9c\x9b\xae\xc2\x46\x63\x58\x2b\x9d\xc3\x1c\x69\x22\xae\x98\x70\xcc\x62\x47\xb1\x5c\x56\xce\x82\xb1\x9a\xde\xef\xfb\x8d\x73\x03\x39\x16\x5c\x62\x4e\x0b\x45\xb7\x59\x92\xac\x46\x45\x8d\xc9\x3d\x32\x1e\x30\xa8\xb4\x9a\x0b\x5c\x1e\xb4\xde\xe5\x2d\x3b\x8b\x1b\x99\x55\x4c\x1b\x6c\xbc\x6d\x5d\x72\x81\x70\x01\xc9\x24\x4b\x20\x15\x96\xfe\x20\xf1\x12\xb8\x3c\x87\x5c\x35\x3e\xc6\xa7\xc9\xe4\x9a\x5e\x98\x8b\xec\x32\x3a\x0f\x3d\xde\xda\xc9\x84\x27\xc0\x5b\x87\xa4\x27\x98\xff\x9f\x69\x79\xd0\x1b\xa6\x67\x4c\x3e\xee\x9d\x83\x93\x36\x0b\xc1\x16\x47\x90\xa9\xb9\x66\x60\xac\xaa\x0c\x90\x88\xb4\x79\xbe\x5c\x62\xce\x99\x45\xb1\xa1\x8d\xae\x31\x7a\x00\x94\xa8\x71\x87\x6d\x7c\x77\xba\xf3\xe2\xfc\xbc\x2f\xd9\xe1\xc1\xf6\xd0\xe1\xae\x90\x51\x5d\x53\xda\xdb\xf9\xc0\x16\x26\x19\x49\x8f\x45\x81\x99\xe5\x2b\x92\x70\x21\xd4\x9c\x09\x12\x34\xb8\x14\xdb\x00\xb7\x90\x31\x09\x99\x66\x6b\x01\xb6\xd4\xca\x2d\x4a\x6f\x63\xa6\x17\x6e\x89\xd2\x1a\x60\xc1\x39\x2a\xad\x16\x9a\x2d\x07\x16\x92\x6c\xc5\x17\xcc\xa2\x89\xa1\x22\x33\x8a\x08\xb0\x1a\xd1\x07\x97\x51\x24\x8b\x77\x60\x26\xd6\x6c\x63\xc8\x0d\x9a\xb0\x75\x52\x23\xcb\xbd\xc7\x0d\xf0\x66\x85\x25\x98\xe2\x32\x27\x75\x6f\xc7\xfa\x11\x69\x9c\xcb\x4c\x23\x89\x4a\xab\xcc\xb1\x50\x1a\x61\xae\x91\x5d\x11\x85\x72\x96\x42\x9d\x08\x85\x52\xd5\xce\x0a\xfb\xfb\xd9\x93\x27\x07\xbb\xca\xf5\x0c\xee\xb2\x13\x1a\x96\x8d\x86\xb8\xe4\x4a\x06\xcb\x47\x6b\x9f\x10\x7e\x45\x4f\x2f\x99\xcc\x05\x92\xa3\x57\x1a\x2b\xb1\x69\x9c\xbc\x8b\xb3\x90\xb4\xb0\x62\x3c\x4a\x8e\xfc\x1c\xa1\x32\x26\x80\x4f\x4f\xc2\xcf\x42\xe9\x0e\xfe\x50\xd4\x25\x93\xe7\x49\x2f\x24\xc6\x50\x70\x41\xfa\x23\x2d\x74\x98\x52\x4c\x65\x25\x49\xba\x67\x61\xc9\x6c\x56\xf6\xf0\xb2\x67\x8a\x00\xf6\x93\xce\x4a\xd3\x29\x85\xa0\xd3\xc9\xe1\x2e\xd0\x87\x55\xb3\x12\xb3\x2b\xa2\xf4\x5c\x5b\x4a\x55\x14\xa8\x09\x21\xea\xa0\x0a\x7e\xc8\xb2\x0c\x2b\x72\x34\xd9\x78\x5d\x94\x90\x1b\x58\x32\x7d\x85\x39\xcc\x37\xf4\x7a\xba\xb5\x10\x2f\xc0\xa8\x23\x60\x60\x2a\x96\x21\x4d\x97\x8a\x76\x24\x1d\x13\x62\x03\x2c\xcf\x31\x07\xc3\x65\x16\xfc\xcd\x19\xd4\x3e\x26\x3e\x54\x98\x11\xa2\x59\x45\x73\xa0\x70\xda\x23\xbf\xc7\xb5\xde\x12\x11\x3c\xda\x3d\xec\xa0\x08\x78\x24\x39\x9c\x0e\x04\x28\x3d\x94\xd4\xfe\xfc\xea\xdd\x2f\x7f\xbd\xe0\x4f\x9e\x5c\x4e\x7b\xac\x06\x09\xce\x77\x43\xf9\xf3\x59\xc3\x27\xf1\xee\xb9\x70\xc1\x5b\xef\xbd\xa1\x7c\xfb\x82\x09\x61\x42\xb4\xb5\x49\x86\xec\xa0\x9c\xee\x66\x91\x15\x4a\xef\x54\x3f\x8e\xc6\xf0\x97\xb7\x2f\xdf\x9e\xb5\x46\xf4\xee\xee\xa7\x31\xe9\x51\x9f\xcd\xc5\x86\x82\x9f\x56\x81\x25\x05\x2a\x7e\xa8\x04\xcf\xb8\x15\x1b\x22\xa7\x4c\xc2\x62\x12\x0b\x39\x4e\x08\xb5\x26\x0e\x75\x36\x33\x21\x9d\x6d\x67\x33\x93\xa9\x2a\x80\x0d\xd3\xe4\x72\x5a\x63\x66\xcf\x46\xe3\x1a\x28\x0c\x49\xa5\xd9\x86\xb0\xa0\xdd\x8d\x09\x49\xbb\x41\xac\x52\x89\xdc\xb4\x44\x67\xb5\xff\xfa\x9d\x5b\xe5\x57\x03\x2e\xad\xda\x8a\xe7\x96\xa3\x0f\xe8\x51\xe3\x9c\xac\xa1\x5e\x33\x03\x0b\xbe\x42\x79\x14\xa3\x23\xe4\x5d\x9f\x7a\xc9\xef\x33\xeb\x98\x68\x66\xd3\x7f\x7e\x31\x8f\xc4\xc6\xa8\x8c\x87\x0c\x1c\x25\x6d\x01\x60\x99\x87\xc8\xbf\xae\xb7\x74\xf1\xfc\xf2\xa6\x8f\x02\x75\x04\x2f\xf3\x18\xba\x61\x66\x32\x1c\xbd\x7d\x24\xda\x2e\x32\xce\x3c\x9f\x08\x48\xf5\x13\x80\xa9\xe3\x24\xe4\x34\x75\x36\xdf\x8a\x58\xb3\x66\x15\xa4\x24\x73\x63\x03\xc9\xc8\xdc\xb4\x9b\x59\x6f\x72\xcb\x70\xda\x6c\xef\xf8\x38\x3d\x9e\xdd\x8c\xee\x90\xb8\xf6\x08\x32\x53\x6b\xd9\xc9\xb5\x60\xa6\x91\xe9\x66\x36\xb9\x6e\xb9\xdf\xf4\x83\xa5\x2b\xfc\xf4\x4e\xba\xbe\x12\x3a\xf3\xa6\x93\x9e\x06\x7a\xf3\x78\x01\x39\x66\x82\xdc\x34\x2d\x20\xe9\xcd\x4c\xe0\x19\x1c\xe7\xb8\x3a\x96\x4e\x88\x01\xdb\xd0\x33\xb9\x55\xb9\x28\xcc\x6e\xc9\xd1\x57\xce\xe4\xba\x4b\x7d\x03\xb9\xc2\x00\x99\x3e\xa8\xfa\x7a\x88\x80\x50\x3f\x21\x9b\xdd\x06\x17\x75\x40\x18\xe5\x74\x86\xb3\x4a\xb8\x05\x97\xb3\x4e\xca\x69\x32\x1d\xd9\x39\xe7\xfa\x96\x7c\x65\xae\x78\x45\x2a\x0a\x0c\xbc\x7c\x94\xa3\x4a\xb6\xc2\x10\x4c\x35\xc0\xe5\x9c\x22\x5c\xe9\xcd\x96\x9b\xa7\x39\x4c\x88\xfd\xa0\x6f\xd3\xda\x05\x55\x93\x7e\xf1\x53\x7f\x28\x3a\x4e\x0e\x7b\x32\x0c\x2b\xce\xef\x2b\x00\x5e\x9b\x45\x7d\x3d\x3c\x21\x86\xbb\x70\x1b\x13\x87\x7f\x39\x94\x32\xe8\x39\xfc\xd1\x0c\xd4\x9e\x1d\x65\x94\x28\x44\x40\x8b\x9c\x1b\x42\xbf\xe9\xfb\x17\xa7\x27\x7f\x38\xb9\x83\x26\x4e\x34\x5d\xe2\x35\xd3\x32\x54\x4f\x01\x72\x32\x46\x3a\x0d\xd8\x1a\x10\x30\x98\xed\x56\xb6\xe1\x75\xbd\x9d\x5b\xa7\x0d\xa4\x2e\xbf\xcd\x9d\x62\x16\xb6\x93\x0f\x74\xeb\x26\x18\x72\x2f\x6f\x0c\xe6\x6c\xd9\x38\x52\x2c\x87\xa0\x17\x65\xbc\x80\x1f\xe0\xae\x63\xc5\x96\x57\x74\xbc\xba\xe0\xdd\x32\xab\x46\xa7\xe9\x7e\x22\xb8\xb1\x69\xa5\xd5\x8a\xe7\xa8\x4d\x02\x89\x50\x0b\x2e\xc3\xbf\xca\xd9\xe4\xa0\x43\x46\x45\x0d\xd1\x84\xb3\x45\x72\xd0\x48\x75\x01\xe9\x3f\x7a\xd8\xbb\x25\x48\x5d\x68\x38\x3d\x70\x4e\x19\x28\x00\x6e\x29\x28\x29\xc4\xbd\x0c\x3e\x0d\x7c\xec\x9c\xf1\x79\x7c\x7b\x19\xe6\x93\x4b\xe3\x01\x0d\xdf\x96\x3c\x7d\xa1\xdc\x18\x7a\xe6\xd5\xfc\x95\xcc\x1d\xed\x06\xb5\xe1\xfc\xab\x24\x4d\x2b\x66\x0c\x95\xc3\xd3\x9d\x91\x94\x62\xa1\x33\xac\xf9\x8a\x59\x4c\xaf\x70\xd3\x1d\x0c\x4e\xd3\x8e\x50\x05\x4a\xf9\x2e\x8e\xdc\x97\x7f\x0c\x55\x9f\xe3\x58\xdf\x1b\x57\x55\x4a\xdb\x50\x6f\x37\xa5\x59\x07\x4a\x37\x68\x6f\x39\xa9\x7e\xf4\x9c\xfa\x05\x3d\xf3\x53\x1d\x6a\xcb\x67\xb8\xb1\xb3\x26\x56\xbf\xae\xf3\xd4\xbe\x43\x76\xfe\xbb\x51\xf2\xde\xa2\xff\x7f\xd7\xba\x3b\x80\xa0\x9c\x7d\x18\xa3\x3e\xda\xf2\xde\x6c\x99\x09\x67\x2c\xea\xaf\x9f\xc8\x99\xb5\x2c\x2b\x13\x48\x62\x98\x12\x58\x24\x90\x68\x5c\xaa\x15\xfa\x3f\x08\x9a\x13\x48\x0c\x5a\x57\x3d\xe6\xf5\x6d\xd6\xff\x79\x5e\x8f\xa6\x9f\x05\x4b\x3c\x46\x72\xef\xf9\x6e\x23\xd9\x67\xdd\x87\xcb\xb5\xc1\x99\x30\x7f\x4c\xbc\x5f\xd0\xc4\x01\x24\x1f\xd0\xc8\x42\xf8\x7f\x9d\x64\x2b\xc6\x05\x9d\x71\x1f\xcd\xfc\x25\xcc\x4c\x29\xf0\xfb\x06\xe6\x6f\x4e\xa9\xbe\x9c\x78\x18\x9d\xfa\x37\x49\x9a\x66\x2c\xcd\x50\x5b\xd3\x9e\x44\xb9\x34\x98\x39\x8d\xcd\x40\xe7\x58\x4a\xbf\x54\xea\x2f\x8f\xba\x03\xe1\x52\xee\xf1\x6c\xfc\xfd\x45\xb7\x92\x05\x5f\x7c\xfd\x82\xdb\x20\xd5\xd7\xa6\x54\xeb\x04\x12\x27\xe9\xe7\x63\x51\x0d\xf7\x5e\x54\x7b\xeb\x12\xca\x3c\xe2\xf6\xa0\xdd\xfe\xbd\x70\x99\x91\xdf\x3e\x6a\xf4\x3e\x35\xea\x21\xe0\x51\xa5\xf7\xa2\xd2\x90\x8c\x1f\xe0\x12\x25\xcf\x77\xae\x4e\x1e\x41\x1d\xee\x1b\xd4\xe3\xff\x41\x65\x79\xfe\x70\x67\x2e\x57\xe5\xcc\xde\xdf\x31\xeb\x1b\x0b\x9d\x07\xbe\xb5\xb8\xd7\x8b\x8a\x6f\x4d\xb5\x0f\x79\x5b\xf0\xdf\xa6\xd3\x5d\x2d\x4e\x4f\xbf\x02\xcc\x3b\x5b\x26\x90\xc4\x23\x2c\xfd\xe5\x53\x78\x7b\x79\x1e\x0f\x83\x77\x40\x3f\x79\xf9\x0a\xb5\xe1\xad\xa3\x87\x79\x6d\x77\x48\x68\x00\x6a\xcf\x5b\xb3\x9c\x6b\x13\xa6\x86\xb6\xdf\x37\xaf\xdf\x4f\x27\x7b\x7f\x93\x7b\xe0\x9b\x44\x53\x0d\x82\x4b\x3c\x87\x5c\x6d\x73\x79\x32\xdd\x4f\x26\xf4\x32\x39\x38\x0f\xfd\x65\x4f\xe1\xe9\x3e\xe9\xa4\x6e\x26\xa1\x78\x87\x34\x6d\x96\xfd\xa9\xd3\x6c\x13\xb6\xd1\xe3\x75\xbd\xb5\x80\xb7\xee\x41\x17\xbe\x63\x6f\xc8\x6b\xdf\x23\x15\x17\xa9\x67\x27\x70\x0b\x87\x4f\xdc\xdd\x96\x4e\x3e\x67\x77\x91\x2e\x25\xba\x81\x4d\xf6\x85\x7f\xd7\x88\xbd\xdd\x59\xc3\x31\x6c\x62\x4b\x92\xd6\xc9\x3f\xd6\xf2\x73\x2b\xb1\xa7\x1e\x83\x54\xdd\x9e\xb1\xba\x39\x2e\x10\x20\x20\xf7\x2d\x99\xde\xa7\x40\xe9\x6e\xd7\xde\xfd\xc4\xf7\xf0\xbd\x40\xe8\x60\x1c\xb8\x1d\xe8\x68\x48\xf3\x45\x69\x41\xaa\xf5\x00\xbd\x6f\xa7\xf1\x6d\x55\x02\xd9\x0a\x7d\xe7\xad\xef\x9e\x56\x16\xa5\xe5\x14\x25\x4a\x43\x8e\x16\x33\xcb\xe5\x22\xee\xcf\x77\xe6\x58\x76\x85\xb0\x62\xc2\xa1\x81\xb9\xb3\xbe\x93\xd1\x60\xc5\xb4\x6f\xfd\x13\xfc\x6a\xb7\x41\x67\x0c\x69\xea\xa5\xf3\x64\xc0\xa5\xb1\xe4\x48\xfe\xdb\x06\x1a\x9f\xfa\xf1\x01\xb2\x35\xee\x69\xf4\x9d\x5f\x6b\xa5\xf5\x86\x24\x61\x73\x92\xb5\xbe\x09\xd9\xba\x04\x01\x5b\xfa\xce\x70\xa3\x80\xdb\x3d\x03\x86\x15\x08\x56\x01\x5f\x48\x15\xbf\x8b\xd8\x59\xe5\xd3\xee\x49\x08\x2a\x76\x0d\xf1\xf0\x77\x2c\x03\x4a\x93\x6a\xb7\x9d\x93\x34\xe2\x43\x8f\x94\x57\x29\x63\xf8\x5c\xe0\x8e\xb3\xf6\xf9\x70\x49\xf1\x2e\xc0\x19\xb6\xc0\xa3\xf6\x73\x90\xd8\x21\x6f\x94\xff\xb6\xc4\x55\xf1\x03\x8c\x6e\x37\x7e\x5c\xd4\xaa\x6e\x6b\xea\x91\x37\xa4\x71\x1a\x07\xef\xab\xc6\x50\xaa\x35\xac\x11\xd6\x4c\x5a\x22\x0d\xaa\x19\x36\xda\x83\x57\xbf\xed\x97\x23\xb7\x34\xc7\x4f\x9e\x37\x08\x32\xd8\x20\xf7\xd3\xc9\xff\xfd\x7f\x7c\x1f\x07\xc1\x49\x67\xb0\xf3\xb9\x4d\x6c\x88\x3b\x82\x4a\xe3\x0a\xb8\xa9\x27\x34\x1f\x6d\xc4\x74\xe8\x74\x98\xe1\xbf\x2a\x01\xff\xd5\x49\xaf\xcf\x33\x2e\x43\x36\xe9\x8d\x7b\xe3\x68\xa5\x2c\x91\x17\xfc\x43\xe0\xd9\x6d\x10\x4d\xfc\x96\xe2\x46\x9a\x46\xee\xe9\xfe\x41\xbd\xb5\x7e\x63\x72\xff\x13\x9e\x39\x33\x65\x07\xe8\xa1\x62\xd9\x15\x5b\x60\xec\x99\x7f\x03\x73\x14\x1c\x57\x08\x4b\x67\x6c\x64\x37\x0f\xe0\xc0\x84\xc0\xbc\x01\x15\xb1\x09\x2d\x8c\xfe\xdb\x30\xaf\xe8\x05\x7a\x11\xab\x99\xdf\xf1\x6c\xbe\x99\x69\x2c\x20\x95\x90\x4c\xcf\x92\x41\x7d\x8c\x06\x2c\xf4\xde\x32\x6d\xfb\xcd\x92\xa0\x24\x10\x10\xc7\x9e\xc9\xf8\x9d\xce\x61\x93\x0d\x3c\x3d\xd9\xbf\xc1\xff\x54\x79\xb9\x62\x3d\x43\x3f\x3b\x7f\x4a\x15\x1a\xfe\xd3\xd7\xbd\x4f\x8d\x3c\x93\x7f\x05\x00\x00\xff\xff\x5e\xa4\x1b\x67\xf1\x36\x00\x00")

func completionShBytes() ([]byte, error) {
	return bindataRead(
		_completionSh,
		"completion.sh",
	)
}

func completionSh() (*asset, error) {
	bytes, err := completionShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "completion.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"completion.sh": completionSh,
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
	"completion.sh": &bintree{completionSh, map[string]*bintree{}},
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

