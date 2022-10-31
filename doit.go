package doit

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func init()  {
	commands["list"] = printCommands
}

func Execute(root any) {
	if root != nil {
		addCommand("", reflect.ValueOf(root))
	}
	cmdString := "list"
	if len(os.Args) > 1 {
		cmdString = os.Args[1]
	}
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
