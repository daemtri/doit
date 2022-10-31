package doit

import (
	"os"
)

var (
	Vars = map[string]interface{}{
		"ext": ext(),
	}
)

func ext() string {
	if os.Getenv("GOOS") == "windows" {
		return ".exe"
	}
	return ""
}

// func BindVar(key string, val any) {
// 	vars[key] = val
// }

// func GetVar(key string) any {
// 	return vars[key]
// }
