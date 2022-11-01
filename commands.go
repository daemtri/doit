package doit

import (
	"bytes"
	"reflect"
)

type command func()

var (
	commands = map[string]command{}
)

func Command(x any) {
	refVal := reflect.ValueOf(x)
	namespace := commandString(reflect.Indirect(refVal).Type().Name())
	addCommand(namespace, refVal)
}

func addCommand(namespace string, refVal reflect.Value) {
	refTyp := refVal.Type()
	for i := 0; i < refTyp.NumMethod(); i++ {
		if !refTyp.Method(i).IsExported() {
			continue
		}
		methodName := commandString(refTyp.Method(i).Name)
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

func commandString(s string) string {
	num := len(s)
	data := make([]byte, 0, num*2)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' {
			data = append(data, '-')
		}
		data = append(data, d)
	}
	return string(bytes.ToLower(data))
}
