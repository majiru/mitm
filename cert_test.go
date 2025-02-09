// Code generated by go-bindata.
// sources:
// cert.pem
// key.pem
// DO NOT EDIT!

package mitm

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

var _certPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x92\x4f\x93\xa2\x4c\x0c\x87\xef\x7c\x8a\xf7\x3e\xf5\x96\x82\x80\x7a\xd8\x43\xd2\x1d\xa0\x07\x1b\x6d\xfe\x29\xde\x14\x18\x5c\x74\x16\x65\xd0\x5e\xfd\xf4\x5b\xe3\x4e\xcd\x65\x73\x4a\x3d\xa9\x4a\x25\xbf\x7a\xfe\xff\x2c\x24\x5f\x44\xff\x31\x8a\x53\xe1\x09\x06\x29\x3d\xa9\x21\x85\xc0\x9a\x33\x06\xb1\xdb\x80\x16\x08\x8d\x50\x99\x07\x96\x7a\x0f\xf3\xfe\xd7\xce\x0c\x71\xe9\x16\xc7\x3a\xd8\x73\x08\xb1\x69\x2e\x87\x63\xbb\x5c\x29\xc5\xa1\x85\x44\xc6\xa0\x0d\xde\x14\x3c\x57\x2a\x24\x5d\x61\x61\x99\x27\x41\xd1\x4d\xa2\xbd\xe1\x29\x39\x92\x17\xf7\x25\x87\x71\x94\x96\x8f\xc8\xec\x36\x3c\x15\xfa\x93\x45\x8f\x6f\xa6\x0d\x6a\x49\x49\xb0\x7d\x30\x33\x62\x07\x19\xa8\x75\x34\x6c\x13\xe4\xfb\x07\x6e\x25\x4a\x1f\xef\x17\x3f\x91\xf6\x1c\x1a\xf2\x19\xfb\xea\x35\x05\x30\x16\x80\x9e\xc5\x0c\x08\x7d\xef\x5d\x76\x60\x9d\x46\x74\xf0\xaa\x2b\xbe\x0e\x1b\x64\xb3\xe0\xb8\x9b\x97\x3c\xec\x1f\x0b\x90\xce\xfd\x55\x0f\x2b\x77\xfb\x71\x57\xb1\x79\x78\x33\x9d\xd3\xac\x4c\xbc\x7b\x7d\x25\x38\x0f\x46\x30\x89\x74\xb7\x38\xcd\xd2\x69\x6f\xda\x36\x54\x2f\xbf\xeb\x36\x6f\x31\xfb\xba\xaa\xe2\x9a\x70\xa4\x15\x81\x16\xec\xcc\x21\xc5\x26\xca\x83\x24\x23\xfe\x37\x93\xde\xc0\x86\x3c\x54\x25\x87\x14\x56\xcf\x61\x2c\x11\xde\x66\x84\x29\x70\x50\xc1\x48\xa2\x7e\x2e\x22\xa5\x72\x89\x4b\xc6\xac\xdc\x2e\xd6\xa6\xde\xfb\xd9\xb5\xb0\xe6\x83\x71\xd0\xf1\x08\x00\x50\x42\xf7\xfd\x23\x82\x64\x30\x6e\x40\x52\xc6\x84\x62\xe5\xcd\xb9\xd4\x59\xf0\x73\x71\x73\xe3\x5d\x72\xe9\x69\x33\x39\xdb\xe8\xd9\x66\x17\xc5\x95\xd1\xf2\x6c\x08\x60\x2a\x76\xee\x8b\x9a\x28\xd1\x64\x2f\x5b\x57\x95\x53\xb6\x10\x7a\x13\xf2\xae\x55\xee\x47\x31\x2d\xa4\x5d\x05\xe1\x7c\x57\x39\x66\xb8\x4e\xeb\x1e\x9d\x32\x4a\xe5\x0f\xe3\x69\x02\x45\xfc\x5f\x3b\xfe\x04\x00\x00\xff\xff\x04\x64\x83\x53\x3a\x02\x00\x00")

func certPemBytes() ([]byte, error) {
	return bindataRead(
		_certPem,
		"cert.pem",
	)
}

func certPem() (*asset, error) {
	bytes, err := certPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cert.pem", size: 570, mode: os.FileMode(420), modTime: time.Unix(1561697857, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _keyPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x3d\x53\x83\x40\x14\x00\x7b\x7e\x05\xbd\xe3\x84\x08\x2a\x16\x29\x1e\xc7\x0b\x1c\x0a\x72\xf9\x7c\x57\x46\x82\x37\x43\xbc\x1c\xc4\xc0\x91\xfc\x7a\xc7\xd4\xd9\x76\x9b\xdd\xc7\x7f\x22\x4c\x78\xe1\x22\x73\xcb\x05\xdf\xc0\x0a\xdd\x77\x94\x37\xe1\xe4\x69\xc5\x40\x20\x72\x90\x5d\x51\x3f\x34\xc5\x57\x76\x1e\xc2\x9a\xc2\x9f\xcf\xbe\x31\xa7\x83\xb7\xdb\xb7\xc6\xc7\xd4\xf0\xfd\xea\x59\xac\xdb\xcb\xb9\x3c\x18\x30\x09\x63\x5d\xb2\xcc\x83\x37\x07\x2c\xa6\x66\x2d\x62\xa1\x00\x49\x72\x30\x72\x2b\x2f\x2a\xde\x51\xb8\xc4\xcd\x53\x80\x7a\x5a\xe1\x87\xad\x17\xa7\xa9\xcd\xbb\x21\xff\x85\xeb\x91\x1f\x99\x37\x69\xf5\x35\x8b\xd2\xed\xdc\x99\x90\xa6\x71\xe4\x94\xbd\x04\x82\x69\xef\xbb\xf2\x99\xea\x69\x2c\xfb\x81\x1a\x15\xf9\xaf\x73\x3b\x9b\x39\xb7\x58\x2c\xe2\xbb\x0f\x7f\x01\x00\x00\xff\xff\x64\x94\xcf\xd7\xe3\x00\x00\x00")

func keyPemBytes() ([]byte, error) {
	return bindataRead(
		_keyPem,
		"key.pem",
	)
}

func keyPem() (*asset, error) {
	bytes, err := keyPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "key.pem", size: 227, mode: os.FileMode(384), modTime: time.Unix(1561697857, 0)}
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
	"cert.pem": certPem,
	"key.pem": keyPem,
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
	"cert.pem": &bintree{certPem, map[string]*bintree{}},
	"key.pem": &bintree{keyPem, map[string]*bintree{}},
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

