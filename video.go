// Package video check if given path is video file
package video // import "github.com/shamsher31/goisvdo"

import (
	"github.com/shamsher31/govdoext"
	"path"
	"strings"
)

// Call Get from vdoext package
var extensions = vdoext.Get()

// Extensions is the extensions for different text file
var Extensions map[string]bool

func init() {
	Extensions = map[string]bool{}
	for _, ext := range extensions {
		Extensions[ext] = true
	}
}

// Is checks if a path is a video file
func Is(p string) bool {
	ext := strings.TrimLeft(path.Ext(p), ".")
	return Extensions[strings.ToLower(ext)]
}
