// +build dev

package assets

import (
	"net/http"
	"path"
	"runtime"
)

var _, file, _, _ = runtime.Caller(0)

var Assets = http.Dir(path.Join(path.Dir(file), "res"))
