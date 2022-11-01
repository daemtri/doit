package doit

import (
	"io"
	"os"

	"github.com/valyala/fasttemplate"
)

var (
	Vars = map[string]interface{}{
		"ext": fasttemplate.TagFunc(func(w io.Writer, tag string) (int, error) {
			if os.Getenv("GOOS") == "windows" {
				return w.Write([]byte(".exe"))
			}
			return 0, nil
		}),
	}
)

// func BindVar(key string, val any) {
// 	vars[key] = val
// }

// func GetVar(key string) any {
// 	return vars[key]
// }

func Setenv(key, val string) {
	if err := os.Setenv(key, val); err != nil {
		panic(err)
	}
}
