package lant

import (
	"bytes"
	"log"
	"os"

	"github.com/valyala/fasttemplate"
)

func Run(cmd string) {
	cmd = fasttemplate.New(cmd, "{", "}").ExecuteStringStd(Vars)
	log.Println("run: ", cmd)
	exe := ParseCommand(cmd)
	exe.Stdin = os.Stdin
	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	exe.Env = os.Environ()
	if err := exe.Run(); err != nil {
		log.Fatalln(err)
	}
}

func RunString(cmd string) string {
	cmd = fasttemplate.New(cmd, "{", "}").ExecuteStringStd(Vars)
	log.Println("run: ", cmd)
	exe := ParseCommand(cmd)
	data, err := exe.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes.TrimSpace(data))
}
