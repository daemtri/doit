package doit

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/spf13/pflag"
)

var (
	list bool
)

func Execute(root any) {
	if root != nil {
		addCommand("", reflect.ValueOf(root))
	}

	pflag.BoolVarP(&list, "list", "l", false, "list commands")
	pflag.Parse()

	if list {
		printCommands()
		return
	}
	if len(os.Args) <= 1 {
		fmt.Println("未指定执行命令")
		os.Exit(2)
	}
	cmdString := os.Args[1]
	cmd, ok := commands[cmdString]
	if !ok {
		fmt.Println("指定的命令不存在", cmdString)
		os.Exit(2)
	}
	cmd()
}

func printCommands() {
	fmt.Println("Targets:")
	globalKeys := make([]string, 0, len(commands))
	keys := make([]string, 0, len(commands))
	for key := range commands {
		if strings.Contains(key, ":") {
			keys = append(keys, key)
		} else {
			globalKeys = append(globalKeys, key)
		}
	}
	sort.Strings(globalKeys)
	for _, key := range globalKeys {
		fmt.Print("  ")
		fmt.Println(key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Print("  ")
		fmt.Println(key)
	}
}
