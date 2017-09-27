// Code generated by go-bindata.
// sources:
// data/client.go
// DO NOT EDIT!

package proxy

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

var _dataClientGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x5b\x6f\xdb\x36\x14\x7e\x16\x7f\xc5\x99\x06\x0f\xd4\xa6\xd1\xed\x30\x0c\x83\x8b\x0c\x68\xe3\x2c\xe8\x1e\xda\x20\x19\xd0\x87\x22\x28\x14\xe9\x48\xe6\x66\x93\x0a\x49\x45\xf5\x0a\xff\xf7\xe1\x90\xa2\x2a\x5f\xd6\x1a\xcb\x43\x2c\x92\xe7\xf2\x9d\xef\x7c\xbc\xb4\x45\xf9\x77\xd1\x20\x6c\x0a\xa9\x18\x93\x9b\x56\x1b\x07\x9c\x25\xe9\xc3\xd6\xa1\x4d\x59\x92\xa2\x2a\x75\x25\x55\x33\xff\xcb\x6a\x45\x13\xf5\xc6\xd1\x8f\xd4\xf4\x7f\xad\x1b\xfa\x51\xe8\xe6\x2b\xe7\x5a\xfa\xd6\xde\xad\x2d\xdc\x6a\x5e\xcb\x35\xd2\x07\x4d\xd8\xad\x2a\x53\x96\x31\x56\x77\xaa\x84\x0a\x1f\xba\xa6\xe6\x6d\xe1\x1c\x1a\x05\xd6\x19\xa9\x9a\x1c\x0a\xd3\x58\x10\x42\x48\xe5\xd0\xd4\x45\x89\x9f\x76\x19\x7c\x62\x89\xac\x83\x07\x5c\x5c\x40\xea\x4c\x87\x29\xcd\x26\x6b\xdd\x88\x1b\x23\x95\xab\x79\xbd\x71\xe2\xae\x0d\xdf\xe9\xfb\x99\x85\x6f\x67\xd5\x3d\xa4\x39\x44\x0c\xe2\x55\x61\x91\x6b\x2b\x5e\x9a\xc6\xbe\x7f\x76\x9f\xe5\xa0\xad\xb8\x46\xd7\xca\x8a\x67\xd9\x0f\x03\x94\x80\x41\x08\x91\xb1\x64\xc7\x76\x8c\x3d\x15\x86\x08\x09\xe9\x07\xa0\x2c\xb1\x68\x9e\xd0\xc4\x61\xc6\x98\xdb\xb6\x08\x06\x1f\x3b\xb4\x8e\xa6\xbb\xd2\x11\x44\x4a\x06\xef\xef\xa3\xdb\x95\x7a\x82\xc9\x70\x29\xcd\x18\x72\x37\xc6\xb0\xad\x56\x16\x27\x41\x5e\x2f\x41\x2a\xf7\xcb\xcf\x64\xe3\xd9\xa3\x6e\x71\x4f\x4c\x07\x8b\x0b\xd8\xab\x9d\xda\xb0\x98\xcf\x67\x76\x9e\xe6\x10\x60\x66\x03\xfc\x9a\xa7\x97\x5a\x29\x2c\x9d\x54\x0d\x38\x0d\x33\x9b\xe6\xd0\xf9\xe5\x1a\x0d\x50\xec\x10\x76\xb4\xff\x5d\x2a\x69\x57\x58\x41\x6b\x74\x89\xd6\xa6\x44\x0b\xcf\x18\x4b\xfa\x2a\x07\x34\x86\xf2\x07\x1e\xfb\x8a\x67\xbe\x53\x34\xfb\xcd\x05\x28\xb9\xf6\xa1\xda\x42\xc9\x92\xa3\x31\x9e\x51\x96\x50\x81\xa3\x2b\x89\xea\x46\x5b\xc7\xbb\x3c\xb2\x47\x3e\x44\xdb\x02\x62\xb3\x9e\x2f\xee\x73\x96\x10\x79\x0b\xa0\xc9\x2b\xf5\x24\x8d\x56\x3c\xa3\xd9\xa5\x34\x0b\x80\xbe\xca\x59\xb2\x23\x5c\x5f\x07\x40\x1d\xed\x1b\x20\x41\x8a\x77\x85\x74\xd7\x46\x77\x2d\x4b\xfa\x46\xbc\xac\x2a\xfe\x13\x05\xa9\xe5\xb4\xb8\x3b\x57\x49\x25\xee\x5c\xe1\xce\xab\x70\x3e\x87\x55\xa1\xaa\x35\xf5\xb0\x92\x8a\x25\x8d\x9e\x72\x6b\x72\xe8\x29\xb2\xd4\xe2\x46\xb6\xe8\xd9\x1c\x09\xf7\xb9\x60\x55\x58\x98\x55\x10\x36\x21\x69\x58\xdc\xc9\x7f\x90\x67\x19\x4b\x28\xff\x38\x86\xdf\xe0\x99\x8f\x19\xd1\x3f\xf7\xc1\x0e\x12\x26\x43\x7f\xfb\x46\x2c\xb5\xa2\x84\x61\x2e\x4a\xa2\xdd\x92\x1e\x6a\xa3\x37\xe0\xd3\xa7\xc1\xe0\x43\x0e\xa5\x6e\xb7\x57\x81\x07\xa9\x05\x59\xf2\x3e\x1f\x29\x09\x66\xb2\x1e\xcd\x26\x9c\x4c\x12\x5c\x19\xa3\x8d\xb7\x19\xd3\x78\x5e\x16\x30\x7b\x4a\xc7\x1c\x21\x58\xd2\x8b\xcb\xb5\xb6\xf8\x4e\xba\x95\xf7\x1b\x68\xa5\x3f\x83\xae\x33\xca\x7f\xef\xd8\xc4\xf6\xa0\x1e\x2a\x71\x3f\xdb\xa4\xa8\x9d\x37\xde\x01\xae\x2d\x0e\xc4\x4d\x82\x50\xf3\x12\x8f\xed\x16\x1f\xf3\x80\x72\x28\x9f\xb6\x95\x78\x83\xfd\x6d\x10\x29\x4f\x6f\xde\xde\xfd\x49\xad\x99\xee\xbd\x99\x9d\x55\x73\xef\x45\xfb\x2a\xf7\x7b\x59\xbc\x5e\x66\x39\x98\xa1\x73\x63\xc8\x29\x55\x41\x3f\x71\x69\x04\x12\x0b\xa2\xfd\x31\x6c\xd8\x21\x38\x99\x7c\x08\x0a\x1d\x90\x2d\xb1\x2e\xba\xb5\xbb\x5c\x4b\x54\x4e\x2c\x35\x8f\x55\x0c\x79\xf1\x64\x4a\x8c\xd9\x86\x3d\xbd\x27\x5c\xdd\xb9\x43\xe5\x46\x44\xd7\xe8\x3c\xa2\x79\x30\xf3\x78\xc2\x67\x3e\x78\x4e\x59\xbb\x46\xc7\x4f\xd2\x44\xae\x7b\x3c\x7d\x26\x69\x08\x71\x92\xa5\xb0\x36\xd2\x74\xa0\xf5\x43\x59\x3b\x4d\xfd\x8f\x28\x93\x28\xe3\xa0\xe1\x09\x62\xf1\x4a\x57\x5b\x6f\x32\x19\x4f\x05\xb6\xb7\x7b\x8e\xce\xc7\x43\xc1\xc5\x84\x5e\x70\x27\xd8\x45\x63\xce\x61\x17\x8d\x89\xec\xa2\x31\xf9\xe0\x79\x16\xbb\xe4\xfa\x1f\xec\x0e\x21\x4e\xb2\x1b\xd6\xce\x65\x37\x96\x1b\x71\x1e\xf0\x3b\xc1\xbc\xc7\x6f\x1c\xff\x7f\x7e\x63\xc2\x09\xbf\xd1\x87\xce\x74\x6f\xaa\xfd\x05\x8d\xc5\xc6\x92\x0a\x6a\x1f\x8a\x9c\xfa\xc6\x9f\xfb\xde\x09\x3f\x4a\x77\xa9\x2b\xbc\x9d\xde\x4a\x5f\x20\x96\xec\x4b\x5d\xe1\x11\xb5\xe7\xdd\x3b\x31\x1d\x5d\xe8\xa3\x4f\xb8\x06\xe9\x74\x59\x22\xc5\x36\x7c\x8a\x2a\x10\x27\xc2\x12\xff\x2e\x2e\x65\x2f\xbe\x96\x8f\x6e\xca\x8f\xd2\x8d\xc1\xb2\xf1\xfd\xf0\xf9\xd6\x1d\x5f\x5e\x06\x1f\xe3\x05\x9c\x01\xff\x3e\xbe\x42\x3c\x25\xda\xf8\xee\x3f\xe8\x6a\x4b\xec\x28\xec\xb9\xbf\x99\xc4\xab\xae\xae\xfd\xf3\x62\x28\x64\xf1\xb9\x92\x2b\x15\x2a\x79\xf0\xe0\xc3\x88\x1b\x7c\x3c\x86\x1d\x8e\x75\x1a\xfa\x64\xe3\x15\x4a\x00\xa1\x80\x3f\xee\xde\xbe\x81\x4a\x97\xdd\x06\x89\x32\x73\xd4\xa7\xf8\x7c\x48\x8b\xb6\x5d\xcb\xb2\x70\x52\x2b\xff\x58\x7d\x01\xe5\xaa\x30\x16\xdd\x45\xe7\xea\x1f\x7f\x4d\x73\x78\x08\x12\x3c\x6e\xd5\x09\x0c\xc3\x9d\x69\x62\x07\x46\xa9\x7a\x70\xb7\x58\xa2\x7c\x42\x70\x2b\xf4\x51\xa1\xb0\x1e\x69\xe8\x72\xe5\x7b\x55\x8d\x6f\xb9\x2f\x74\xda\x1c\x77\x78\xf0\x3e\x93\xa9\x61\x32\x7a\xe5\xb4\xcc\x76\xec\xdf\x00\x00\x00\xff\xff\x5e\x01\xc5\x55\xd9\x0b\x00\x00")

func dataClientGoBytes() ([]byte, error) {
	return bindataRead(
		_dataClientGo,
		"data/client.go",
	)
}

func dataClientGo() (*asset, error) {
	bytes, err := dataClientGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/client.go", size: 3033, mode: os.FileMode(420), modTime: time.Unix(1505521439, 0)}
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
	"data/client.go": dataClientGo,
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
	"data": &bintree{nil, map[string]*bintree{
		"client.go": &bintree{dataClientGo, map[string]*bintree{}},
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
