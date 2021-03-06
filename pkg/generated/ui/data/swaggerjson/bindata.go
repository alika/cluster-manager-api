// Code generated by go-bindata.
// sources:
// assets/generated/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\xfd\x6f\xdb\x38\x96\xbf\xf7\xaf\x78\xf0\x1d\x70\x29\xd0\xd8\x6d\x0f\x7b\x58\x74\xb1\xd8\xcb\x26\xdd\x8e\xd1\x36\x0d\xe2\x7e\x60\x70\x1d\x18\xb4\xf4\x6c\x73\x22\x91\x2a\x49\xd9\xe3\x39\xf4\x7f\x5f\xf0\x43\xb6\x24\x8b\xb6\x2c\x3b\xb1\xd2\x7a\x80\x41\x63\x89\x7a\x7c\x5f\x22\xdf\x17\x9f\xfe\xff\x09\x40\x47\xce\xc9\x64\x82\xa2\xf3\x0a\x3a\x2f\xbb\xcf\x3b\xcf\xf4\x35\xca\xc6\xbc\xf3\x0a\xf4\x7d\x80\x8e\xa2\x2a\x42\x7d\xff\x32\x4a\xa5\x42\x01\xef\x09\x23\x13\x14\x70\x71\xd3\x37\xe3\x01\x3a\x33\x14\x92\x72\xa6\x47\xcd\x9e\x77\x33\x40\x00\x9d\x80\x33\x45\x02\xb5\x84\x06\xd0\x61\x24\x36\xe0\x06\x24\x96\x29\x9b\xc0\xe5\xf5\xe5\x47\x37\x1c\xa0\x93\x8a\x48\xdf\x9c\x2a\x95\xc8\x57\xbd\xde\x84\xaa\x69\x3a\xea\x06\x3c\xee\x49\x3b\xfe\x3c\x60\x81\xea\x05\x16\x97\xf3\xd8\xe2\x72\x4e\x12\xba\x82\x81\x31\xa1\x06\x0a\x09\x63\xca\xfe\x37\xff\x60\x97\xf2\x8e\x19\xf6\xfd\x09\xc0\x77\x43\xad\x0c\xa6\x18\xa3\xec\xbc\x82\xff\xb3\x38\x9b\xb9\x33\x02\xf4\x0f\xfd\xc4\x6f\x66\x6c\xc0\x99\x4c\x0b\x83\x49\x92\x44\x34\x20\x8a\x72\xd6\xfb\x5d\x72\xb6\x1a\x9b\x08\x1e\xa6\x41\xcd\xb1\x44\x4d\xe5\x8a\xe5\x3d\x92\xd0\xde\xec\x45\x46\x65\x9e\x7b\x13\xcc\x33\x53\xa3\x9f\xc6\x31\x11\x0b\x4d\xee\x17\x1a\x45\x20\x50\x09\x8a\x33\x04\x35\x45\x90\x8a\xa8\x54\x02\x1f\x03\x01\x07\x0c\x08\x0b\x81\x2a\x09\x77\xe9\x08\x03\xce\xc6\x74\x02\x63\x2e\x20\xe0\x8c\x61\xa0\xe8\x8c\xaa\xc5\x92\x95\x00\x1d\x9e\xa0\x30\x28\xf7\x43\x3d\xc7\x1b\x54\x4e\x0f\xf2\x83\x04\xca\x84\x33\x89\xb2\x80\x1b\x40\xe7\xe5\xf3\xe7\xa5\x4b\x00\x9d\x10\x65\x20\x68\xa2\x9c\xc6\xe4\x00\x59\x8a\xb4\x40\xc8\xda\x63\x00\x9d\xff\x14\x38\xd6\x4f\xfc\x47\x2f\xc4\x31\x65\x54\x43\x90\x19\x97\x86\x4e\x17\x86\x24\xa1\x2b\x2c\x6f\x31\x89\x16\x9d\x02\xa0\xef\x4f\xaa\xfe\xfe\x9e\x23\x27\x21\x82\xc4\xa8\x50\xac\x84\x67\xff\x2b\x11\x92\xa9\xb2\xf9\xf7\xd9\x46\x22\xaf\x49\x8c\x5a\x0e\x5a\x2a\x99\x24\x14\x87\x11\x42\xc4\xf9\x1d\x86\x90\x26\xdd\x32\x08\x6a\x9e\xfc\x96\xa2\x58\x94\x6f\x09\xfc\x96\x52\x81\x5a\x24\x63\x12\x49\x2c\xdd\x56\x8b\xc4\x20\x26\x95\xa0\x6c\x92\x27\xff\xfb\xb3\xed\xe4\x90\xb9\xec\x4a\x0c\x04\xaa\xe1\x1d\x2e\x86\x34\xdc\x42\xdb\xc7\x29\xc2\xc0\x8c\x7f\x8b\x8b\x7e\x68\xd4\xe9\xe2\xa6\x0f\x17\x41\x80\x52\xb6\x91\x2c\x62\x30\xd3\xd4\xd5\x26\xcd\x12\xf3\x16\x17\x4b\xf2\x48\xfb\xc8\x13\x38\xd1\x88\x6f\xa7\xe9\xd6\x0c\x6c\x2d\x29\x7f\xa6\x02\xbb\x24\x49\xea\xe9\xde\x45\x92\xb4\x58\xeb\x0c\x2d\x0a\x19\x61\xaa\x06\x2d\x1f\xcd\xc0\x76\x0b\x26\x21\x52\xce\xb9\xa8\x23\x9a\x1b\x37\xb4\xdd\x04\xc9\x74\xb4\xc4\xbc\xe6\x72\x97\x7b\xa2\xad\xb4\x25\x82\xcf\x68\x58\xd8\xa8\xab\xc8\xc9\xef\x4c\xd9\x23\x12\xce\xc8\x5c\xf6\x0c\x77\x7a\xb3\x78\x4e\x04\xf6\x50\x05\x4f\xef\x8f\xb6\xd2\x5d\x64\x69\x5c\xda\x7b\xcd\xf5\x94\x99\x9d\x1f\xcb\x32\x02\xb3\xfe\x55\x5c\xd4\x14\xac\x5f\xb6\x24\x15\xcd\x82\xdf\xd6\xd8\x34\x26\x69\xa4\xad\xad\xdc\xac\x95\xa6\x43\xee\xc9\x8e\x22\x93\xb2\xd1\x90\x19\xce\xab\x87\x7f\x73\x7f\x2d\xa5\xd8\x09\x31\x42\x85\x9b\x2d\x3b\x3b\x66\x65\xc9\x6d\xb0\xd2\xae\xcc\xd0\x47\x60\xa8\x15\x10\x6d\x8b\xad\xf6\x65\x4a\x14\x50\x99\xb7\xd5\xfe\x4b\x82\x7e\x50\x9b\x6c\x21\x4a\x25\xf8\xa2\x35\xaf\xf9\xc9\x5a\x3b\x59\x6b\xc7\x25\xe5\x64\xad\xb5\x59\x30\x27\x6b\xed\x64\xad\x1d\x9a\xb6\x47\x68\xad\xd5\x10\x41\x40\xa2\x68\x44\x82\xbb\x6e\x2a\xa2\x1a\x5a\xf5\xe9\xf6\x9d\xb6\x07\xf4\x53\xa0\x1f\x03\xc5\x5b\xa3\x4d\x4b\x52\xf4\x2c\x28\x55\xbd\xf7\xa4\x7f\x95\xa9\x96\x7b\xec\x61\xe8\x39\xa4\x21\x9d\x70\xb9\x25\x40\x6a\xde\x1a\xa9\x57\x82\x3a\x96\xf4\xa5\x40\xf2\x28\x2c\xe9\x02\xa2\x0f\x62\x49\x8f\x78\xb8\xa6\x04\x56\x3f\xaa\xee\xe4\xd4\x43\x89\xb4\xac\x1d\x87\x66\xc0\x7b\x39\xa9\x43\xfe\x21\x14\x2e\xdd\xa2\x6f\x24\xfc\x3d\x95\x0a\xc8\x8e\x8a\x77\x61\x1e\x73\x08\x5c\xf3\x10\x65\x9b\xb5\xaf\x80\xed\xcf\xa8\x7d\x05\x06\xdc\xbb\xf6\x3d\xc9\x71\xaf\x9c\x2d\xea\x45\xb4\xb0\x04\xee\x90\x32\x22\xa0\x9f\xd5\x7b\x80\x83\x25\x6b\x65\x82\xde\xe9\x09\x5b\xac\x9c\x45\x4c\x1f\x44\x3b\x4f\x4e\xf9\xc9\x29\x3f\x2e\x29\x27\xa7\xbc\xcd\x82\x39\x39\xe5\x27\xa7\xfc\xd0\xb4\x3d\x42\xa7\xfc\xbe\x4d\xa1\x34\x99\x08\x12\xe2\xae\xd6\x50\x2a\x18\xb8\x47\x81\x1b\x29\x4b\xa3\xa1\x04\x26\x74\x86\xac\x86\xf5\xfe\x06\xd5\x27\x0b\xc0\x61\xde\x67\x63\x2e\x62\x33\xa2\xe5\xa6\x92\x17\xef\x07\x31\x9c\x4e\x6f\xdc\xd1\xc3\x60\x27\xdb\xf5\x64\xbb\x1e\x97\x94\x93\xed\xda\x66\xc1\x9c\x6c\xd7\x47\x61\xbb\x36\x2b\x74\x00\xa5\xaf\xcd\x11\x88\x40\x08\x38\x93\x7a\x67\xa5\xcc\xd6\x08\x3b\x9b\xe8\x11\x66\x05\xb6\x06\x69\x95\xc2\x38\x51\xa0\xf8\xd2\xf0\xab\x13\xa4\x2d\xda\x4a\x6d\x36\xec\x8a\x98\xfe\x8c\x11\xda\x22\x07\x8e\x14\xa2\x15\x18\x22\x53\x94\x44\xb2\xa7\x4d\xbf\x9c\x5f\xb2\x55\x45\xd3\x24\x24\x0a\x81\xcc\x25\xe4\xc0\x40\x2a\x31\x74\xce\x49\x1d\x7d\xd5\x30\x2e\xbe\x0c\x2e\x57\x10\xda\xad\xb5\xeb\xf8\xfe\x9c\xba\xbb\xce\x87\x16\x68\xb0\xf1\x53\x9a\xe8\xb0\x7e\x70\x6f\x2d\xd6\x40\x1e\x97\x1e\x97\x30\xfe\x89\x35\xb9\xc4\x89\xe3\xe8\xf2\x14\xa3\x38\xaf\xbe\x3b\x14\xe1\x86\x98\x44\x7c\x81\x21\x68\x18\x10\x4c\x89\x50\x1b\x14\xd6\x56\xb9\xfe\x82\x51\x7c\x59\x1e\xd9\x36\x4d\x2d\xa1\xfa\x20\x2a\xba\xfe\xc6\x57\x11\x9b\x1d\x3f\x54\x34\x8a\x50\x80\x9c\xf2\x34\x0a\x61\x84\x40\x99\x54\x24\x8a\x30\x04\xce\x5a\x63\x7d\x27\x24\xb8\x23\x13\xcc\x58\x5b\xc7\x10\x5f\x23\x4c\x2f\x8a\xad\x21\xa8\xac\xe3\x95\x22\xd2\x63\xe0\x9a\xc4\x0f\xe4\x22\x9c\x82\x59\xa7\x60\xd6\x29\x98\xf5\xd0\xb4\x9c\x82\x59\x6d\x26\xe8\xc7\x0c\x66\x9d\xd2\x42\x47\x4f\x0b\xd5\xae\x24\xbe\x80\x94\xd1\x6f\x29\x02\x0d\x41\x71\xa0\x2c\xa4\x81\xf6\xfa\xd4\x94\xca\x87\x2d\x2c\xae\x63\xd6\x9c\x6a\xbe\x7f\x82\x9a\x6f\xe7\x22\xc0\x19\x17\x20\xd0\xfd\x7a\x9a\x73\xde\xbe\xb2\x8f\x5a\x3b\xe7\x7a\xf0\x08\xed\x19\xbf\x34\x50\x74\x86\x40\x4d\xf3\x0c\x63\xd9\x4e\x89\x04\x12\x09\x24\xe1\x02\x46\x88\x6c\xe5\x06\xce\xa9\x9a\xda\x8e\x1b\x7a\x05\x2a\x85\xde\xcb\xee\x60\xdf\xce\xff\x28\xfc\xc1\x32\xae\x3f\x63\xcc\xa2\xcc\x83\xe3\x44\x2b\xac\x7b\x58\x08\xb7\x35\xd5\x7b\x07\xca\xaf\xa1\x37\x59\xa9\xfa\xc7\xb5\x91\x6d\x53\xd0\x12\xaa\x3f\xa3\x7e\x96\x58\x70\x1c\xf5\x5c\xb5\x80\xda\xb9\xd4\xca\x3d\x0a\x74\x55\x6b\x04\x64\xc4\x53\x05\x24\xa1\x20\x51\xcc\xb6\xd5\x5a\x7d\xb6\x10\x1e\x4f\x91\x95\x43\xb8\x91\xb6\x36\x11\xd6\xb2\xdb\x55\x0e\xb5\x55\xbf\xa9\xf2\xe1\x85\xc2\xef\xcf\xef\xbf\x10\x81\x83\x04\x83\xbc\x6c\xb3\xad\x9d\x8f\x7e\xc7\x60\xb5\x83\x69\x43\x3d\x41\xa1\x68\x89\xd9\x1d\x12\x86\x43\x66\x4e\xb1\x94\x64\x90\x01\x22\x42\x90\xe2\xab\xd3\xa1\x0a\xe3\xf2\xf8\x5d\xf8\x6c\x31\x7f\x4f\x82\x29\x65\x96\x00\x9f\x45\xb5\x6a\x72\xe6\x46\x4b\x98\x4f\x69\x30\x85\x39\xc2\x5c\xfb\xb0\x8a\x03\x09\x8d\x39\x9b\xcb\x95\x77\x2a\xc5\x23\x30\xe6\x33\xbc\x5f\x62\xcb\x02\xb3\x94\xde\x9a\x99\x0f\x45\xaf\xa5\x03\xc6\x82\xc7\x1e\xa2\x2b\xd7\x83\xba\xa8\xed\xa1\x4b\xd3\xf2\xbe\xb7\xd9\x05\x5b\x11\xab\xed\x5d\xfd\xb0\xf1\x72\x07\x83\x5f\x9c\x97\xbb\x8d\xa4\xc2\xc9\xb2\x8b\xb7\x03\x8d\xff\xc5\xdb\x81\x31\x0b\x58\x80\x6f\x04\x4f\x93\x7d\xe8\x71\xdb\x4a\x33\x7a\x58\xce\xe5\x9d\x18\x4c\xaa\x17\x0d\x0b\x6f\xf7\x39\x32\x2a\x41\x8f\x86\xb3\x81\x22\x2c\x24\x22\x1c\x5e\xbd\x1c\xce\x5e\x3e\x03\x54\x41\xf7\x69\xf5\x94\x31\x65\xc3\x6f\x29\x61\x8a\xaa\x85\x6f\x6a\xca\x14\x96\xc3\xe3\x1d\xbb\x88\xbb\xdb\xff\xfd\xd2\x83\xd8\x7b\xca\x68\x9c\xc6\xc0\xd2\x78\x84\x42\xb3\x80\x3a\x54\x25\x9c\x39\x87\x57\x6a\x45\xfe\x13\x05\xf7\xa1\x48\xfe\xb8\x57\x14\xc9\x1f\xcd\x50\x5c\x73\x71\xd6\xc5\x61\x64\x2d\xc1\x3a\xf3\x40\xac\x78\x08\x0b\xab\x26\xeb\x6c\x50\xe7\x2f\x56\x9d\xbf\x0c\xae\x88\x22\x97\xc8\x4a\xfd\x06\x77\xd5\x65\x17\xb5\x6d\xa2\x69\x5f\xcc\x0a\x64\x01\xc0\x59\x2a\xcf\x91\x48\x75\xfe\x62\xa3\x8e\x91\x19\xa1\x11\x19\xd1\x88\xaa\xc5\xf0\x4f\xce\x0e\xb1\xe4\xd6\xf4\xc2\xcb\x88\xe7\x51\x01\x83\x4a\x8e\x88\xd1\x33\xc8\xfe\x7e\x19\x98\xbf\xe7\xa8\xff\x0e\xd7\xa9\xf3\xcb\xde\xce\xa3\xe5\x04\x4e\x50\x75\xe4\x7a\xb0\x65\xea\x50\x4b\x48\xfc\x97\x6e\x44\xc4\x04\x4f\x8b\xc7\x8f\xb3\x78\xdc\x08\xd7\x4b\x34\x15\x18\xf6\x4b\xef\xd5\xce\x9a\x36\x4b\x82\x21\x0d\x1b\x6f\x89\x9f\x6f\x2e\x81\x86\xcf\x60\x14\x11\x76\x67\xf6\x7a\xfd\xff\xd7\x4e\x60\x70\x07\xce\xd0\x5c\x58\xf0\xf4\x6b\xe7\x19\x8c\xa9\xc9\x1d\xd3\xb1\xbe\x60\xca\x1e\xff\xf9\xeb\x07\x0d\xa3\x5a\xea\x12\x83\x54\xe8\xd5\xc6\xf0\xb0\x29\x9a\x03\x07\x65\xd3\x96\x4d\x49\x3c\x14\x3c\xc2\x21\x11\xcd\x16\x54\x13\xde\xbb\x78\x0f\x1a\x88\xa1\x38\x5f\xe7\x79\x46\x04\x7b\x9a\xc9\x51\x4a\x1e\x50\xe3\x37\x85\x61\x2d\x4d\xfa\x17\x17\x30\x9f\x22\x03\xc9\x63\x13\xd9\x65\x13\x69\x78\x97\x85\xc7\x2c\xaf\xc3\xa2\xf2\xbc\x41\x86\x82\x06\xcb\x40\x4a\x16\x51\xe2\x92\x2a\x2e\x16\xfb\xa8\x8c\xed\x1b\xdc\x64\xd3\x59\x95\xc0\x7e\xba\x7d\xb7\x62\x93\x09\xf8\x09\x4c\x78\xb5\x70\x1a\xdb\x6c\xf9\x09\x35\x78\x1b\x2c\xdc\x62\x81\x96\x3c\xc6\x0a\x77\x77\x0f\xd6\x4d\xa8\x1a\xae\x7b\xef\xbb\xa9\x99\x22\x13\xe0\xcc\x1a\xa1\xd4\xf2\xcd\xc9\xb4\x92\x7b\x7a\xca\x80\xc7\x31\xdd\xc3\x8e\x27\x72\xba\xb4\x7b\xa9\x02\x07\xce\x3b\x9d\x12\x88\x43\xa9\x88\x6a\x2a\x36\x54\x53\xbd\x32\x0a\x60\x5c\x99\x59\x35\x44\x98\x13\x09\x41\x84\x84\xd9\xd7\x61\x94\xd2\xc8\x83\x84\xbe\x15\x0e\xc3\xa6\x08\x5c\x99\xa5\x6b\x6c\x66\x08\x3d\x64\xf2\xbd\xe4\xe8\xb4\x4a\x4f\x32\xe1\xb6\x22\x4f\x71\xcd\xd7\x84\x46\x58\x3d\xa3\xbb\x29\x1a\xcd\x77\xe9\x1e\x36\x53\x55\xc3\x4f\x22\xa2\xb4\x8e\x37\x82\x7f\xe3\x1e\x06\xaa\xac\x98\xec\x7c\x36\x7f\xdb\x03\x91\x32\x46\x99\x56\xdb\x6d\x6f\x5f\x55\xbb\x87\x62\x09\xed\x1e\x6f\x5f\xb1\x38\xa5\xe9\xdb\xe0\x2f\x59\xf1\xee\x62\xa5\xda\x91\xfd\x66\xf6\x55\x94\xf8\x42\x24\x8d\x9d\x04\x6f\xc1\x47\xad\x6d\x4b\x3f\x9d\x2f\x3b\x55\x46\xcf\x6d\xff\x73\xbd\x61\x69\x7d\x58\x8b\x75\x6c\x52\x83\x72\xd7\x8f\x63\x44\x01\x1a\x9e\xe2\xf0\xbc\x71\x59\xa2\xbd\x84\xc9\x6e\xc1\xe8\x72\xa6\x7e\x85\xec\x2e\x19\x7a\x8f\xcb\x37\x5f\xf3\xf1\xea\x77\x69\xf1\x55\xbe\xe7\x51\xbc\xf8\x32\x80\xfc\xa8\x6a\x2c\x4a\xb5\xcf\x3b\xe2\xe1\xaf\x5d\x2e\x60\x62\xea\xa4\xb7\xe2\xe2\x2a\x00\x6a\x21\x53\x37\xd4\xeb\x7f\xfb\x34\x7b\x64\x82\x01\x1d\xbb\x6f\x19\x7c\x65\x05\x20\xce\x29\x30\x07\x15\xfe\x0e\x7f\xf9\x5b\xf9\xf6\x5b\x77\xdb\xd0\xf6\x77\xf8\x9f\xbf\x79\xb6\x16\x97\xcd\x6e\xcc\xe3\xcb\x0c\x80\x67\xf7\x71\xb7\xf3\x59\x87\x26\x1b\xc1\x7a\xe3\xa3\x3d\xd6\x00\xbe\x46\x6e\x06\x61\xc4\xb9\xb6\x32\x7c\x8e\x64\xe5\xed\xd5\x0a\x41\xa4\xad\xbe\x20\x20\x53\xb3\x5a\x8e\xd3\x28\xcb\xf7\x37\x21\xba\xac\xbe\xfb\x24\x06\x92\xc6\x4e\x94\xaf\x5e\xce\x13\x01\xb5\x65\x6b\x4d\x27\xaa\x2c\x66\xf3\xac\xa1\x59\x4d\x59\xd3\xb9\x3c\x95\x66\x9e\xad\xbc\x54\xf0\xd5\x78\x23\xf7\x97\x81\x1d\x65\x6f\xbd\x5c\x5f\x00\x1e\xcc\x2b\xf4\x94\xfa\xf8\xac\x99\x65\xb9\x4d\x63\x2f\xbd\x5c\x84\xd3\xe0\xa5\x74\x6b\xd0\x15\x2a\x42\xa3\xbe\xc2\x78\x1f\xc6\x35\xa4\x65\x45\xc7\xc6\x14\x59\x63\x53\xa7\xe2\x2b\x2a\x9e\x97\xc2\x7c\xf8\x66\x18\xa3\x94\x64\xd2\x6c\xae\x8b\x30\x34\x9b\x0b\x89\x2a\x12\xd3\xc5\x8f\xeb\x6c\x45\x67\xf5\xad\x9d\xbd\x2d\xbc\xdc\x67\x7b\x8c\x77\x66\xbe\xda\x53\x2b\x31\x69\xf1\x6d\xbe\x9f\xda\x4b\x03\x0b\xc5\xaf\xcb\xdb\xd8\xb2\xa3\x3e\x9f\x34\xb9\x55\x9a\xfc\x38\x95\x68\x50\xc6\xda\xc7\xb8\xf5\x62\xdd\xce\xe0\xe3\xc5\xc7\x4f\x83\xe1\xa7\xeb\xc1\xcd\xeb\xcb\xfe\xbf\xfa\xaf\xaf\xf2\xb5\x25\x37\xb7\x1f\x3e\xf7\x07\xfd\x0f\xd7\xfd\xeb\x37\xf9\xeb\xb7\x9f\xae\xd7\x2e\xbd\xbe\xfc\x70\x7d\xd9\x7f\x57\xba\x3c\xf8\xf8\xe1\xe6\xa6\x74\xed\xf5\xed\xed\x87\xdb\xfc\x85\xab\xd7\x6f\x6e\x2f\xae\x5e\x5f\x65\x2c\x58\x16\xef\xe4\x4b\x7b\x37\x60\xba\xe2\xee\x39\xac\x0f\x7b\x05\xd7\x5c\x81\x44\xf5\x95\xc1\x39\xe4\x49\x7a\x05\xc6\x1e\xc9\x5d\x31\xa2\xc1\x65\x55\x6f\xd1\xed\xa4\x12\x46\xa8\x37\x7a\x17\xfe\xed\x1a\x80\x8e\x17\x16\x96\xfb\xb1\x11\xcc\x94\x48\x5b\x5f\xe9\xc0\xd8\x6f\x96\x49\x18\xa7\x51\xb4\x80\x54\x92\x51\x84\x0e\xf4\x8a\xa7\x0e\xfc\xea\x42\xc5\x14\x44\xd9\x80\xf5\x9c\x8b\x3b\x0d\x90\x98\x12\xcf\x68\xe1\xb0\x0e\x39\xc3\x2c\x8c\xe9\x70\x79\xa6\x2d\xe6\x29\x10\xe9\xdc\xe6\xcc\x8a\x89\x89\xc1\xd4\x84\x03\x43\x04\xc9\xc7\x4a\x3b\x61\x16\xab\x4c\xa4\x16\xa5\xec\x57\x3d\xce\xd9\x03\x86\xa1\x81\x63\xd4\xc0\x02\x31\x7f\x6e\x84\x10\x93\x85\x39\x34\xc6\x2c\x7f\x0c\x80\x4c\x6d\x2c\x8c\xec\xd7\x46\x30\xae\x5a\xcd\x1c\xc9\x15\x86\x41\x9a\x21\x1c\x04\x4a\xc5\x05\x1a\x19\xc0\x38\x65\x81\x5d\x51\xa8\x5a\x6c\xb7\xe4\x2a\xea\x27\xf6\x59\xd2\x23\x1e\x94\x83\xde\xf5\x17\x43\xe3\x34\x18\xdf\x73\x2d\x91\x5a\x5e\xeb\x82\x4a\x0f\x67\xb7\x05\xaf\xa6\xa7\x7f\x59\x34\x9c\x4d\x90\x77\xfb\x62\x9c\xe5\xe8\x6c\x2e\xea\xde\xaa\x8d\x6a\xd5\xbf\xe4\x00\x79\xd2\xe4\xa5\xec\x63\x93\xd5\xbc\x22\xf9\xb8\x8f\x22\x85\x44\x91\x61\x50\xae\x79\xa8\xcd\x89\xaa\xd2\x89\xcd\x71\x93\x87\x51\xb9\x3a\x41\xae\x66\x0a\x27\x50\xf2\x54\x04\xeb\x55\x16\x3b\xb1\xab\x22\x59\x5c\x8d\xe4\x3f\x7f\xfd\x00\x56\x53\xdb\xa3\xfe\x9e\xba\x8a\x07\x57\xff\x63\x45\x9b\xfd\x1f\xb2\x5c\xf6\x37\xf7\x66\x73\xf6\x8d\x2d\x17\x18\x90\x05\x9a\xb7\x84\x2b\xb3\x49\x8b\x31\xcb\x47\x1d\x73\xf4\xb3\xe1\xd8\xe9\x87\x25\xb3\xcf\x21\x48\x85\x40\xa6\xa2\x85\x8d\x3c\x6a\x63\x6b\x2e\xb5\xbd\x14\x13\xe2\x79\x9f\xef\xfe\x2a\xf7\x4e\x44\xcf\x56\x49\xcc\xb7\xe9\x08\x05\x43\x85\xbe\x10\xfe\x1e\x89\x84\xca\x9d\x68\x87\x98\xf9\x7d\x24\x15\x2a\xcd\xac\xca\x79\x76\x49\x18\x6c\x9b\x28\x97\x33\xa8\x9c\x6b\x4a\x27\xd3\x61\xbe\x36\xee\x7e\x22\xdc\xeb\xe9\xf9\x9c\x65\xfd\xcb\x85\xc7\x7b\x47\xa5\x3d\x81\xe1\x98\x8c\x04\x0d\x1a\xeb\x9c\x7d\xdc\xad\x82\xa5\x64\x76\xa3\x77\xba\xb5\x79\x04\x3f\x97\xe7\x44\xe6\x37\x00\x18\x2d\x0a\xe7\x3d\x3d\xeb\xad\xb3\x35\xf6\x8c\x26\x98\x10\x91\x5f\x3c\xa1\x89\x88\x96\x83\x09\x19\x66\x90\x9d\x0e\xd9\x57\x68\x87\x39\x27\x11\x70\xa6\x04\x8f\x86\x49\x44\xd8\x3d\x1f\x22\xb8\x9f\x13\x13\x01\x8f\x13\x41\x25\x6e\x37\x26\xf5\xab\x87\xe2\xc7\xa6\x91\x24\x74\x88\x2c\x4c\x38\x6d\x9c\x7a\xa2\x32\xd7\x12\x86\xc0\x8c\x44\x29\x42\x44\xef\x10\x68\xf2\x2a\xe1\x42\xb9\x64\xbf\x3b\xba\x4a\x60\x46\x85\x4a\x49\x04\xfd\x9b\x9e\xbe\xfd\x95\xdd\x10\x29\xf5\x2b\x69\x93\x18\x80\x7f\x28\x14\x8c\x44\x10\xa4\x52\xf1\x18\x85\x74\x4b\x17\x19\x45\xe8\x4a\x7e\xe2\x94\xb9\x43\xdb\xbe\xc0\x72\xdd\x77\xa4\xe2\x0b\x9f\xed\x5b\xd8\x2e\x0d\x7f\xf3\xab\x43\x76\x08\x98\x2f\x7c\xd6\x6c\x75\x24\xb4\x5e\xf5\x67\x21\xba\xd9\x3c\xe5\x53\xd9\xac\xa9\x7d\xdc\xb5\xe9\x67\x84\xe5\xb9\xbd\x5c\x16\xda\x53\x5a\xbd\x47\xb4\xdb\xd8\xa3\x0e\x80\x84\x39\x0a\xb4\x5d\xdb\x1b\xf0\xb7\x5c\xaf\x7a\x74\xfb\x3a\x7f\xd8\xc7\x9e\x3d\x8f\x91\x79\x36\x57\x3d\x56\x26\x24\x38\xcc\xac\x06\x92\xfd\xfc\xad\x9e\xd6\x5c\x27\x49\x12\x39\x7b\x76\x43\x72\x34\xe1\xf5\xf6\xf7\x6d\xc5\xc1\xdb\xf1\x5c\xd5\xec\x6e\xaa\x3d\x0d\x4a\x92\x6c\xc6\x13\x3b\x57\xa9\x76\x37\x6f\x62\xef\xe1\xc7\xac\xcf\x93\x41\xab\x9e\x4a\xef\x08\xcd\x16\x22\x33\x13\x11\xe8\x5e\x4f\xfb\xa5\x83\x33\x86\x52\x61\x08\x0b\x12\x47\x70\x6e\x6e\x7d\x36\x53\x74\xcd\x15\x6d\xa1\x20\x53\xd2\x53\x9e\x65\x10\x1e\x26\x64\x11\x71\xb2\x53\xea\x2d\xb7\x9a\x2c\x14\x6e\xf2\xe8\x2d\x6c\xfd\x1e\x64\x8d\x19\xce\x4c\xeb\x48\x2a\x26\x24\x49\x30\x04\xc9\xa3\xd4\xd0\x62\x8e\x98\x54\x9d\x8d\xd8\xfe\xe6\xab\xb5\x4f\x46\xb5\x6f\x61\xed\x97\x72\x08\x36\x73\x21\x17\x52\x1b\xc4\x9b\x2c\xee\x87\xb4\xb7\xf2\x46\xfa\x76\x4b\xeb\x5d\xf9\x03\x64\xfb\xc8\xee\x87\x93\xdb\xbe\x9e\x52\xae\x38\xa4\x11\x5f\xb7\x7c\x1b\xa4\x7d\x9c\xbe\x24\xac\x6c\xd6\xb9\xf2\x57\x8f\x55\xe7\x56\xda\xe3\x1c\x23\x24\xcb\x95\xde\x1e\xa6\x49\xb8\x94\x54\x1b\xe5\x82\x4e\xa6\x0a\x18\x9f\x37\x13\x5b\xa1\xdb\x40\xfb\x84\xd4\x1f\xe7\x4c\xc3\x39\x91\xf0\xe1\xed\x46\xe1\x0c\x69\xe5\xc1\x17\xd8\x64\x5b\x6c\x3b\x3f\x53\x8d\x59\x76\x24\xa2\xbf\x31\x5e\xbb\x9e\x64\x37\xb3\x98\xd2\xb7\x2c\x1e\x59\x30\xf2\x37\x48\xab\xaa\xd7\xcc\x3e\xe1\x84\xea\x75\xa3\xde\xf9\x8c\x2c\xba\xbf\xa9\xe1\xaa\xaf\x02\x72\xd9\xfd\x74\xd2\x70\x76\x7f\x4b\xd4\x1d\xcc\xba\x1d\x5a\x72\x94\x2c\x7d\x0f\x4b\x8c\xb5\x31\x40\xa5\x28\x9b\xdc\x43\x58\xf9\xc7\xac\x4f\x6f\xff\x99\x82\x47\x92\xf8\xa9\x6e\xc6\xd5\xbe\x05\xfd\xd1\x3a\xfc\xab\x94\xd1\x3b\x32\xc2\xe8\x28\xfe\x7e\xbe\xa9\x07\x81\xc8\xe0\xe1\xf7\xfb\x9a\x27\xca\x4c\x1c\xb1\x72\x8e\xba\xdc\xba\x59\x7f\xaf\x77\xa8\xab\xab\x6a\x7f\x59\x6a\x7c\xb9\xd6\xf2\xb2\xd4\xec\xb2\xb2\x0a\xae\xd4\xe0\x72\x1b\xfe\xc5\xbe\x59\xc7\xce\xa8\xbb\x0d\xaf\x07\x6e\x07\x85\xac\x7f\xf8\x2e\x5e\xc1\x7d\xef\xee\xfb\x45\x96\xae\x57\xd1\x24\xff\xd4\xd4\x33\x35\x8d\x9b\xae\x12\xe6\x49\x38\xa3\xe3\x65\xea\xca\xe9\xcc\x53\x7b\xe6\xdb\xe2\x72\xf6\x2d\x25\x8b\x2e\xe5\x3d\xc9\x63\x9c\xa4\x0b\xd7\x75\xcf\xb3\x6d\x1c\xe8\xb4\xad\x9d\xc3\x7c\x8e\xc0\x36\x4c\xb5\x4d\xfb\x7a\xcb\x4f\x11\x29\x0e\x67\xb3\x97\xdd\x17\x2f\xba\xcf\xd7\x22\x1a\xeb\xea\x30\x9c\xd3\xd0\xcb\xa3\x43\xf8\xad\x8e\x57\xcb\x4f\x75\x9c\xeb\x09\xdd\xd5\x7f\x7c\x65\x03\x2b\x4f\xaa\x60\x4a\x66\x58\x1c\x44\xc2\x98\x32\x48\x04\x9d\xd1\x08\x27\x28\xff\xe1\xb1\x5e\xf4\xb0\xe1\x52\xcf\x8e\xe1\x8d\x2d\xd5\xd4\xd5\x8d\x52\x95\x4f\xc3\xb8\x4c\x89\x25\xc7\x9b\xce\x3f\x59\x82\x27\x4b\xb0\x84\xe1\x81\x2c\xc1\xca\xae\x97\x27\x43\xf0\x60\x86\xa0\xef\x83\x4b\xc7\xb0\x0f\xb2\x9d\xda\x9f\x6c\x68\x63\xd5\xe8\x3e\x9c\x6e\xb7\x3a\x1f\xee\x1c\xad\xff\x63\x48\x3f\xb0\x9e\x1d\xbc\x20\x7e\x3f\x6e\xff\x3c\xba\x56\xfe\xfc\xe1\x1e\x34\xb7\x7f\x1f\x3f\xd9\x5e\x65\x5c\xda\xd1\x5f\x64\xef\xbc\x78\x16\x50\xd7\x33\xcf\x31\x7f\x2a\x8b\x7b\x9c\x89\x47\x62\xd2\x55\x7d\xa1\xf5\xc7\x5f\x96\xd6\xab\xf1\xf6\xa0\x39\x95\x28\xf6\x8a\xbb\x65\x00\xfc\x0d\x82\x73\x7a\x75\x8f\x6d\x88\xf3\x2e\x13\xf7\x57\xab\xec\xd1\x5d\xd2\x14\x52\x70\x51\x13\x87\x06\x0d\x2d\x2a\x3e\x35\x72\xf7\x57\xb9\x7c\x7d\x97\x7d\xf6\x6c\xbd\x2b\x98\x7a\xd7\x2e\x98\x72\x47\x73\xa2\x92\x4a\xe0\x2c\x5a\x00\xc9\x8e\x5a\xd2\xb1\x55\xbc\x31\xc5\xc8\x9c\x06\x75\x61\xbe\x6e\x35\xce\x26\xa8\xf9\x90\x75\x0e\xe5\xd8\x71\x0e\x8c\x27\xc4\xa0\x79\x62\xd1\x5c\xb2\x23\x76\x85\xa6\xb5\x3b\x6e\x14\x4e\x15\xb8\x2f\xa8\x66\xd7\x0c\x23\x5d\x14\x74\xd9\x00\x3e\xab\xfe\xbc\xe2\x41\xae\x03\x7c\x49\x5a\xef\xb9\x40\x77\x3e\x3c\xb3\x0a\xdf\x5b\x4a\xe1\xe2\xa6\xef\xf8\xe6\xba\x6b\x74\xa6\x4a\x25\xf2\x55\xaf\x37\xa1\x6a\x9a\x8e\xba\x01\x8f\x7b\x92\xc4\x32\x65\x93\xf3\x80\x05\x2a\xe3\xd5\xb9\xe3\xd5\x39\x49\xa8\x46\xea\xfb\x93\xef\x4f\xfe\x1d\x00\x00\xff\xff\xe8\x04\xa0\xf9\x83\xaa\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 43651, mode: os.FileMode(420), modTime: time.Unix(1539060021, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": &bintree{apiSwaggerJson, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
