// Code generated by vfsgen; DO NOT EDIT.

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 9, 10, 21, 40, 58, 715425786, time.UTC),
		},
		"/repositories.html": &vfsgen۰CompressedFileInfo{
			name:             "repositories.html",
			modTime:          time.Date(2018, 9, 10, 21, 40, 58, 715425786, time.UTC),
			uncompressedSize: 1904,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xc9\x6e\xe4\x36\x10\xbd\xf7\x57\x94\x89\x39\x8e\x44\x18\x39\x8c\xd3\xa0\x84\x04\x1e\x1f\x0c\x64\x19\x18\x31\x90\x60\x30\x07\x36\x55\x92\x68\x53\xa4\x86\x2c\xb5\xd1\x51\xf4\xef\x01\xb5\xa0\x57\x1b\xce\x21\xba\x50\x7c\xf5\xea\xb1\x16\xaa\xd4\xf7\x05\x96\xda\x22\x30\x8f\xad\x0b\x9a\x9c\xd7\x18\xd8\x30\xac\xc4\xd5\xe7\xdf\x6f\xff\xf8\xeb\xcb\x1d\xd4\xd4\x98\x7c\x25\xae\x92\xe4\xab\x2e\xc1\x10\xdc\xdf\xc1\xa7\x6f\x39\x8c\x8f\x88\x56\x50\x46\x86\x90\x31\xeb\x92\xa7\x00\x86\x12\x8d\x3f\x4e\xcb\xcd\xb4\x7c\x62\x39\x88\xab\xaf\x68\x0b\x5d\x7e\x4b\x92\xbd\xda\xa1\xd4\x3b\xd4\xde\x90\xb9\x79\x8f\xcc\x6b\xfe\x15\xcd\x12\x11\xc8\x2f\xf8\x8f\x8e\x49\x72\xec\x5c\xa3\x2c\xf2\xd5\x78\x60\x83\x24\x41\xd5\xd2\x07\xa4\x8c\x75\x54\x26\x37\x6c\x36\x6d\x64\x40\xa8\x3d\x96\x19\xe3\x0c\x0e\xf9\x35\x51\x9b\xe0\xf7\x4e\x6f\x33\xf6\x67\xf2\xf8\x73\x72\xeb\x9a\x56\x92\xde\x18\x64\xa0\x9c\x25\xb4\x94\xb1\xfb\xbb\x0c\x8b\x0a\x3f\xaa\xda\xbb\x06\xb3\xeb\x45\x97\x34\x19\xcc\xfb\x1e\xd2\x07\xac\x74\x20\xbf\xfb\xec\x1a\xa9\x2d\x0c\x83\xe0\x93\x71\x22\x1a\x6d\x9f\xc1\xa3\xc9\x98\x56\xce\x32\xa0\x5d\x8b\x19\xd3\x8d\xac\x90\x6b\xe5\xd8\x12\x5c\x20\x49\x5a\xf1\x52\x6e\x23\x2f\x8d\xa6\x33\x85\x40\x3b\x83\xa1\x46\xa4\x53\x37\x15\x02\x9f\xac\xa9\x0a\x81\x01\xcf\x57\x82\x4f\x15\x12\x1b\x57\xec\x66\xa9\xfa\xfa\xb5\x90\xeb\xeb\x99\x52\x3a\xdf\x4c\xaf\xe3\x56\xdb\xb6\x23\xb0\xb2\xc1\x8c\x95\xda\x10\xfa\x25\x85\x80\xd2\xab\x9a\xe5\x42\x2e\xbd\x52\x06\xa5\x67\xf9\xb8\x08\x2e\x67\x41\x3e\x29\x4e\x9b\x42\x6f\x17\xf6\x8b\x97\x6d\x8b\x9e\x1d\x1c\x46\x72\xb3\xd4\x6d\x8f\xf9\x63\x60\x02\xeb\xfc\x61\xf9\x62\x76\xf0\x9b\x6c\x50\x70\xaa\x2f\x13\xbf\x74\xc6\xc0\xad\x6b\x1a\x69\x8b\x73\x96\xe0\xa7\x07\xf4\x3d\x78\x69\x2b\x84\x0f\xcf\xb8\xfb\x08\x1f\xb6\xd2\x74\x08\xeb\x2c\x96\x6d\xff\x91\xc2\x30\xbc\x27\xce\x02\xb6\xd2\xe8\xca\x66\x8c\x5c\xcb\xce\x19\x23\x4b\x2e\xcd\x8c\x53\x80\xf7\xfd\x7c\x66\x1a\xf3\x82\x7f\xa0\xf3\xe6\x7b\x87\x7e\x07\xc3\xc0\x49\x56\xe1\x15\x99\x13\xbf\x93\xf8\xf6\xf9\xca\x0b\x71\x72\x2a\xe6\x0e\x9d\x86\x3f\x47\xef\x75\x55\x13\x03\xeb\x62\xd7\xfe\xd7\x34\x84\x72\x05\xe6\x85\x53\xcf\xe8\xa1\x8d\xbd\xdb\x2b\x3d\x3e\xdc\x8f\x97\x75\xa4\xfc\xd7\xf4\x4e\x90\xf3\xae\xa3\x2d\x0e\xab\x26\xf8\xc1\x75\x14\xbc\xd0\xdb\x0b\x97\xb8\x74\x8e\x8e\xef\x70\x9b\xff\x2a\x0b\x84\x17\x4d\xf5\x9c\x8b\xf8\x61\x8e\x18\x36\xbb\x7d\x95\xe2\xf8\x09\x6b\xce\x2b\x4d\x75\xb7\x49\x95\x6b\xf8\x13\x86\x50\x7a\xf9\x37\xcb\x7f\x5a\x5e\x63\x3e\x82\xb7\x47\xfa\xb7\x35\xaa\x67\xd7\x11\x50\x8d\x10\x5c\xe7\x15\x42\xd4\x07\x49\xeb\x37\xf5\x2b\xb4\x9d\xb6\x48\xce\x99\xc0\x3d\x56\x2c\x7f\xc3\x78\xe9\xe4\x5f\x64\x20\x78\x6c\x0b\x49\x58\xac\x63\xc9\xd2\x88\xcc\xc0\xd8\x99\x63\xfe\x64\x81\x7b\x4b\xe8\xb7\xd2\xac\xc7\x32\xa7\x13\xba\x80\x87\x6e\x53\x95\xe3\xa8\xe7\xe9\x54\xd9\x38\xe9\x47\x4b\x50\x5e\xb7\x04\xc1\xab\xfd\xd0\x7b\x0a\x7c\x82\x43\x1a\xff\x11\x62\xde\xc5\xd1\x37\xcd\x3c\xc1\xa7\xbf\x67\xdf\xa3\x2d\x86\x61\xf5\x6f\x00\x00\x00\xff\xff\x82\xfc\x97\x4b\x70\x07\x00\x00"),
		},
		"/tags.html": &vfsgen۰CompressedFileInfo{
			name:             "tags.html",
			modTime:          time.Date(2018, 9, 10, 21, 40, 58, 715425786, time.UTC),
			uncompressedSize: 2868,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xdf\x6f\xdb\x36\x10\x7e\xf7\x5f\x71\x25\x0a\xb4\x05\x2c\x31\xe9\xd6\xb5\x73\x24\x61\x85\x9b\x61\x1d\xf6\x0b\x43\x32\x6c\x08\xf2\x40\x4b\x27\x89\x0e\x4d\xaa\xe4\xc9\x89\xa7\xe9\x7f\x1f\x68\xc9\x8e\xec\xd8\x41\x86\xee\xa1\x7e\x11\x7d\x3f\x3e\x7d\xf7\x9d\xc8\x63\xd3\x64\x98\x4b\x8d\xc0\x48\x14\x8e\xb5\xed\x28\x7a\xf6\xe1\xd7\xe9\xc5\x5f\xbf\x9d\x43\x49\x0b\x95\x8c\xa2\x67\x41\x70\x25\x73\x50\x04\x1f\xcf\xe1\xed\x75\x02\xeb\x5f\xe4\xbd\x90\x2a\xe1\x5c\xcc\xb4\x09\xe6\x0e\x14\x05\x12\xbf\xed\x1e\xef\xba\xc7\x5b\x96\x40\xf4\xec\x0a\x75\x26\xf3\xeb\x20\xb8\x47\x1b\x42\x3d\x01\xed\x11\x98\x77\x4f\x81\x39\x96\x5f\x50\x0f\xe1\x0d\xc9\x81\xfc\x75\x62\x10\xec\x26\x97\x28\xb2\x64\xb4\x7e\xe1\x02\x49\x40\x5a\x0a\xeb\x90\x62\x56\x53\x1e\xbc\x63\xbd\x6b\x26\x1c\x42\x69\x31\x8f\x19\x67\x30\x8c\x2f\x89\xaa\x00\x3f\xd5\x72\x19\xb3\x3f\x83\xcb\xf7\xc1\xd4\x2c\x2a\x41\x72\xa6\x90\x41\x6a\x34\xa1\xa6\x98\x7d\x3c\x8f\x31\x2b\x70\x9c\x96\xd6\x2c\x30\x3e\xdd\xe0\x92\x24\x85\x49\xd3\x40\xf8\x3b\x16\xd2\x91\x5d\x7d\x30\x0b\x21\x35\xb4\x2d\xf7\xd6\x5f\xc4\x02\xa1\x6d\x23\xde\x05\x76\x49\x4a\xea\x1b\xb0\xa8\x62\x26\x53\xa3\x19\xd0\xaa\xc2\x98\xc9\x85\x28\x90\xcb\xd4\xb0\x0d\x51\x47\x82\x64\xca\x73\xb1\xf4\x71\xa1\x77\x3d\x40\x70\xb4\x52\xe8\x4a\x44\xda\x4f\x4b\x9d\xe3\x9d\x37\x4c\x9d\x63\xc0\x93\x51\xc4\x3b\xb5\xa2\x99\xc9\x56\x3d\x54\x79\xfa\x14\xfa\xe5\x69\x1f\x9e\xc9\xe5\xa6\x27\xb7\x56\x54\x15\xda\x9e\x53\x27\x87\x98\x6d\xaa\xbc\xb7\xd9\x5d\x43\x67\x2c\x13\x8f\x1d\x71\x2a\x0f\x7b\x2f\x44\x71\xdc\x39\xb5\x28\x08\xb3\xc3\x01\x4d\x23\x73\x08\x7f\x10\xee\x8f\x5a\x69\xd7\xb6\x3e\xc1\x2f\xd1\x8a\x99\x54\x92\x24\xba\x75\x62\xd3\xa0\xce\xda\x76\x97\x2c\xdf\x67\xdb\x34\x60\x85\x2e\x10\x9e\xdf\xe0\x6a\x0c\xcf\x97\x42\xd5\x08\x93\xd8\x2b\x56\x19\x27\xc9\x58\x89\x0e\xf6\x71\x0e\x17\x9d\xc1\x52\x28\x59\xe8\x98\x29\xcc\x89\x81\x36\x5e\xc3\x87\x91\xdb\x32\x9e\x0f\xeb\x10\x9b\x06\x5b\xac\x8c\x6f\x4f\x47\xa6\x6b\xd2\x3f\x50\x5b\xf5\xa9\x46\xbb\xf2\xbd\x23\x51\x0c\x02\x2e\x44\xe1\x8d\x4b\x8f\xc3\x0e\xd6\x3d\xa8\x76\x08\x7a\x34\x6a\x9f\x1a\x17\x47\x61\x23\x4e\xd9\x61\x2d\x7a\x29\xac\x2c\xca\x2f\x5b\x8b\x2e\xe7\x4b\x91\x62\xc3\xaa\xdf\x04\xe1\xf7\xc6\x2e\x04\x01\x3b\x79\x0d\x3f\x0a\x3d\x86\xd7\x27\x27\xdf\xc0\xe9\x9b\xc9\xc9\xd7\x93\x93\x37\x70\x79\x31\x65\x87\xc8\x1f\xe6\xf2\xa0\x9c\xcf\x64\xfb\x3f\xf5\x09\x64\x16\xb3\x07\xdf\xe6\x64\x3f\x9a\x1d\x26\xb1\x7b\x6a\x39\x59\x68\xa1\x58\x12\xf1\x4c\x2e\x8f\xb0\xe6\xe2\x40\x9b\x8e\x08\xf6\xc4\x43\x04\x75\x36\x6c\x43\xc4\x07\x47\x65\x4f\xe5\xc1\x01\x9b\x1b\x43\xbb\xe7\x6b\x95\xfc\x2c\x32\x84\x5b\x49\x25\x44\xa9\xc9\x30\x89\xbe\x8a\xf8\x7a\x01\xb3\xd5\xbd\xdc\x7e\xa8\xb9\x09\xe7\x85\xa4\xb2\x9e\x85\xa9\x59\xf0\x39\x3a\x97\x5b\xf1\x37\x4b\xbe\xdb\x2c\x7d\x9d\x11\xaf\x76\xf0\xa7\x25\xa6\x37\xa6\x26\xa0\x12\xc1\x99\xda\xa6\x08\x1e\x1f\x04\x4d\x1e\xc5\x2f\x50\xd7\x52\x23\x19\xa3\x1c\xb7\x58\xb0\xe4\x11\xe7\xa1\x37\xff\x24\x1c\xc1\x65\x95\xf9\x8f\x7a\xe2\x25\x0b\xbd\xa5\x37\xac\x87\x50\x35\x94\xcb\xdf\x04\x78\xd8\x49\xe4\x2f\x02\x6b\x8f\x4b\xad\xac\x08\x9c\x4d\xef\xe7\xe0\xdc\xf1\xce\xec\x42\x7f\x85\x88\xfa\x7f\x5d\xc6\xfe\xb0\xd8\x81\xe9\xe6\x32\xe1\x1d\xf1\xb9\x58\x8a\xce\xba\x6d\xc7\x52\x58\x10\x73\x71\x37\x15\x4a\x39\x88\xe1\x6a\x34\x68\xf7\x7f\x9e\x19\x2f\x3e\x6b\x87\x84\x73\x67\xf4\x8b\xf1\xe8\xd8\x07\x77\x7d\xd6\x2f\x6e\xa5\xce\xcc\x6d\x68\xb4\x32\x22\x83\x18\xf2\x5a\xa7\x24\x8d\x7e\xf9\x0a\x9a\x6d\xf6\x7b\x6b\xc5\x2a\xac\xac\x21\xe3\x25\x08\x73\x63\xcf\x45\x5a\x86\xa9\x50\xea\xe5\xb6\xe4\xf1\x7d\x72\x6d\xd5\x18\xa4\xce\xf0\xee\x55\x33\xa8\xc9\xbf\x63\x38\x79\x57\x53\x53\x6b\xf2\xd1\xaf\xce\xb6\x61\xed\x76\xdd\x9e\xf5\xed\xdd\x6d\x50\xb7\xc1\x22\xde\xdd\x59\x22\xde\xdd\x84\x37\xf6\x7f\x03\x00\x00\xff\xff\x88\x1c\x69\x6a\x34\x0b\x00\x00"),
		},
		"/vulns.html": &vfsgen۰CompressedFileInfo{
			name:             "vulns.html",
			modTime:          time.Date(2018, 9, 10, 21, 40, 58, 715425786, time.UTC),
			uncompressedSize: 2715,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xdf\x6f\xdb\x36\x10\x7e\xcf\x5f\x71\x21\xfa\x18\x99\x48\x82\xa1\x59\x40\x0b\x5b\x93\x60\x2b\xd0\xfd\x40\x9b\x0e\x1b\x8a\x62\xa0\xa4\xb3\xc4\x86\x22\x35\xf2\xe4\xcd\x13\xf4\xbf\x0f\x12\x2d\x5b\x96\xec\x2c\xc3\xfc\x20\x4b\xbe\xbb\xef\xee\xfb\xee\x74\x74\xd3\x64\xb8\x52\x06\x81\xad\x6b\x6d\x3c\x6b\xdb\x33\x71\x7e\xff\xd3\xdd\xe3\x6f\x3f\x3f\x40\x41\xa5\x8e\xcf\xc4\x79\x14\x7d\x52\x2b\xd0\x04\x6f\x1f\xe0\xf5\xe7\x18\xfa\x8f\xe8\xac\x90\x6a\xe9\xfd\x92\x19\x1b\x7d\xf1\xa0\x29\x52\xf8\x75\xf8\xba\x09\x5f\xaf\x59\x0c\xe2\xfc\x13\x9a\x4c\xad\x3e\x47\xd1\x1e\x6d\x0c\xf5\x02\xb4\x67\x60\x6e\x5e\x02\x73\x2a\x3e\xa7\x2d\x44\xf7\x43\x7c\x24\xbe\x0f\x8c\xa2\xc3\xe0\x02\x65\x16\x9f\xf5\x09\x4b\x24\x09\x69\x21\x9d\x47\x5a\xb2\x9a\x56\xd1\x0d\xdb\x9a\x12\xe9\x11\x0a\x87\xab\x25\xe3\x0c\xc6\xfe\x05\x51\x15\xe1\x1f\xb5\x5a\x2f\xd9\xaf\xd1\xc7\x6f\xa3\x3b\x5b\x56\x92\x54\xa2\x91\x41\x6a\x0d\xa1\xa1\x25\x7b\xfb\xb0\xc4\x2c\xc7\x8b\xb4\x70\xb6\xc4\xe5\xe5\x80\x4b\x8a\x34\xc6\x4d\x03\x8b\xf7\x98\x2b\x4f\x6e\xf3\xf1\xfd\x3b\x68\x5b\x1e\x7e\xaa\x2c\xb4\xed\x6d\x77\xff\x28\x73\x68\x5b\xf8\xa5\xd6\x06\x9d\x4c\x94\x56\xb4\x81\xce\xc1\x91\xe0\x01\x25\x20\x6a\x65\x9e\xc0\xa1\x5e\x32\x95\x5a\xc3\x80\x36\x15\x2e\x99\x2a\x65\x8e\x5c\xa5\x96\x0d\x2c\x3c\x49\x52\x29\x5f\xc9\x75\xe7\xb7\xe8\x4c\x33\x04\x4f\x1b\x8d\xbe\x40\xa4\x69\x58\xea\x3d\x4f\xac\x25\x4f\x4e\x56\x8b\x52\x99\x45\xea\x3d\x03\x1e\x9f\x09\x1e\x14\x15\x89\xcd\x36\x5b\xc4\x4c\xad\x87\x46\x74\x8a\x48\x65\xd0\x6d\xb3\xf5\x76\xbb\xeb\x53\xe2\x50\x66\xa9\xab\xcb\x64\x64\xdf\x56\x15\x0b\xb9\x6b\xc1\x11\xc9\x04\x97\xb1\xe0\x5a\xcd\xe2\x06\x6c\x99\x92\x5a\xe3\x10\x3a\x95\xf6\x30\x54\x70\xab\xe3\xb3\xfd\xe3\x88\x41\x25\x73\x8c\x3a\x8e\x07\x1c\xc2\xc4\x5e\xfe\x87\x56\x0a\x5f\x4a\xad\xe3\xe3\x1d\x0d\x36\xc1\x8b\xcb\x71\x4d\x99\x5a\x8f\x1e\xab\xa1\x22\xc2\xbf\x28\x72\x2a\x2f\x88\xc5\xdf\x61\x87\x46\x98\x81\x35\xb7\xd0\x34\x8b\x7b\x49\xd8\xb1\xab\x46\x74\x9a\x26\xbc\x2e\x8b\x37\x32\xeb\xf2\x7b\xf8\xaa\x6d\x8f\x92\x95\x1a\x1d\x41\x7f\x8d\x32\x69\x72\x74\x0c\x9c\xd5\xb8\xb5\x4c\x04\x68\x9a\x1d\x62\xdb\xc2\xf7\x2a\x2f\x2e\xe0\xce\x29\x52\xa9\xd4\x17\x20\x4d\xc6\xad\x83\x7b\x5c\xa5\xd6\x5c\xc2\x7a\x44\x5c\xa1\x87\x95\xad\x4d\x76\x8a\x6c\xd3\xa0\xc9\xda\x76\xd4\x91\xe2\x2a\xfe\x50\x97\xa5\x74\x1b\xc1\x8b\xab\x91\x2c\xf5\x6e\x98\xb4\xf2\x14\xe5\xce\xd6\xd5\x7c\x98\xe6\x3e\x91\x22\x2c\x27\x8e\xbd\xb3\xaf\xa4\xd9\xcd\xa7\xcc\xf2\x30\x42\x1a\x0d\x2c\x82\x78\x9d\xbc\x9d\xd3\x91\xd8\x24\x7e\xb4\x24\xb5\xe0\xc9\xa4\x82\xd9\xa0\x36\x8d\xeb\xf4\x85\x57\x4f\xb8\xb9\x80\x57\x6b\xa9\x6b\x84\xdb\xe5\x36\xc7\x9b\xcd\x07\x5c\xa3\x53\xb4\x19\x35\xea\xff\x51\xd1\x32\x41\x0d\xfd\x35\x6a\x9a\xd4\x6a\xeb\xfa\xe4\x6d\x0b\x55\xad\xf5\x30\x50\x5b\xaa\xdb\x82\x4e\x53\x6d\x9a\x3e\x1a\xa6\xf5\x1d\x21\x1a\x5a\xb9\xf7\xa8\x0f\x5e\xb5\xe2\x2a\xbe\x47\x92\x4a\xfb\xc3\xc6\xee\x04\xea\x4f\xb7\x7f\x91\x66\xef\x3c\xe8\x18\xc2\x4e\x8c\x79\x25\x0d\x6a\xe8\xaf\x51\x86\x2b\x59\xeb\xe9\x68\xcf\xbc\xfb\x15\xa0\x4c\x7e\x4c\xe6\xe2\xfa\xd0\xb5\xdf\xcf\x9d\x94\xa1\x9c\xc5\x8f\xb2\xc4\x89\x4e\x2f\x6f\x51\x80\xd8\x73\x9e\x74\x6b\x66\x3f\x39\x9b\xbc\xb8\x9e\x0e\xe5\xc1\x3b\x77\x9c\x76\xb7\xd6\x8f\x70\xde\x25\xbe\x47\x9f\x3a\x55\x91\xb2\x66\x36\x0a\x2f\x80\x5f\x59\x4b\xb3\xc5\xda\xbb\x0e\xcb\x7f\x97\xe9\x9d\x32\x4f\x6d\xcb\x80\xa4\xcb\xbb\x03\xfb\xf7\x44\x4b\xf3\x34\xd2\x20\x38\x74\xe7\xc2\x73\x75\x9c\xd8\x34\xd3\xe7\xbd\x7f\x28\xf1\x60\xf9\xa6\x68\xe6\x55\x8b\x2a\xfe\x41\x66\x08\x7f\x2a\x2a\x40\xa4\x36\xc3\x58\x5c\x0b\xde\xdf\x40\xb2\xd9\x33\xea\xfe\x42\xf8\x5b\xce\x73\x45\x45\x9d\x2c\x52\x5b\xf2\x2f\xe8\xfd\xca\xc9\xbf\x59\xfc\xcd\x70\x1b\xce\xb7\x6a\x96\xe3\xae\xc0\xf4\xc9\xd6\x04\x54\x20\x78\x5b\xbb\x14\xa1\xcb\x01\x92\x6e\x9f\xcd\x91\xa3\xa9\x95\x41\xb2\x56\x7b\xee\x30\x67\xf1\x33\xc6\x59\x76\xc1\x83\x0e\xdb\x23\x3e\x88\x28\x78\x38\xf5\x05\x0f\xff\x37\x07\xf1\xfe\x09\x00\x00\xff\xff\x30\xe0\x25\xc3\x9b\x0a\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/repositories.html"].(os.FileInfo),
		fs["/tags.html"].(os.FileInfo),
		fs["/vulns.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
