package lant

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	list bool
)

func Main() int {
	pflag.BoolVarP(&list, "list", "l", false, "list commands")
	pflag.Parse()

	if list {
		printCommands()
		return 0
	}
	if len(os.Args) <= 1 {
		fmt.Println("未指定执行命令")
		return 2
	}
	cmdString := os.Args[1]
	cmd, ok := commands[cmdString]
	if !ok {
		fmt.Println("指定的命令不存在", cmdString)
		return 2
	}
	cmd()
	return 0
}

func printCommands() {
	fmt.Println("Targets:")
	for key := range commands {
		fmt.Print("  ")
		fmt.Println(key)
	}
}
