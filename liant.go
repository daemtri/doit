package lant

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/pflag"
)

var (
	list bool
)

func Execute(roots ...any) {
	for i := range roots {
		addCommand("", reflect.ValueOf(roots[i]))
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
	for key := range commands {
		fmt.Print("  ")
		fmt.Println(key)
	}
}
