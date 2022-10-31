package lant

import (
	"reflect"
	"strings"
)

type command func()

var (
	commands = map[string]command{}
)

func RegisterCommand(x any) {
	refVal := reflect.ValueOf(x)
	namespace := firstLower(reflect.Indirect(refVal).Type().Name())
	addCommand(namespace, refVal)
}

func addCommand(namespace string, refVal reflect.Value) {
	refTyp := refVal.Type()
	for i := 0; i < refTyp.NumMethod(); i++ {
		if !refTyp.Method(i).IsExported() {
			continue
		}
		methodName := firstLower(refTyp.Method(i).Name)
		{
			methodVal := refVal.Method(i)
			commands[joinNamespaceMethod(namespace, methodName)] = func() {
				methodVal.Call(nil)
			}
		}
	}
}

func joinNamespaceMethod(namespace, method string) string {
	if namespace == "" {
		return method
	}
	return namespace + ":" + method
}

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
