package lant

import (
	"fmt"
	"reflect"
	"strings"
)

type command func()

var (
	commands = map[string]command{}
)

func AddCommand(x any) {
	refVal := reflect.ValueOf(x)
	refTyp := refVal.Type()
	switch refTyp.Kind() {
	case reflect.Pointer:
		switch refTyp.Elem().Kind() {
		case reflect.Struct:
			registerStructMethodsCommand(refVal)
		default:
			panic(fmt.Errorf("暂不支持注册该类型: %s", refTyp))
		}
	case reflect.Struct:
		registerStructMethodsCommand(refVal)
	case reflect.Func:
		panic("稍后支持")
	default:
		panic(fmt.Errorf("暂不支持注册该类型: %s", refTyp))
	}
}

func registerStructMethodsCommand(refVal reflect.Value) {
	refTyp := refVal.Type()
	className := firstLower(reflect.Indirect(refVal).Type().Name())
	for i := 0; i < refTyp.NumMethod(); i++ {
		if !refTyp.Method(i).IsExported() {
			continue
		}
		methodName := firstLower(refTyp.Method(i).Name)
		{
			methodVal := refVal.Method(i)
			commands[className+":"+methodName] = func() {
				methodVal.Call(nil)
			}
		}
	}
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
