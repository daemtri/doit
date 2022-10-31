package doit

import (
	"bytes"
	"github.com/valyala/fasttemplate"
	"io"
	"log"
	"os"
	"strings"
	_ "unsafe"
)

func Run(cmd string) {
	cmd = fasttemplate.New(cmd, "{", "}").ExecuteFuncString(tagFunc)
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
	cmd = fasttemplate.New(cmd, "{", "}").ExecuteFuncString(tagFunc)
	log.Println("run: ", cmd)
	exe := ParseCommand(cmd)
	data, err := exe.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes.TrimSpace(data))
}

//go:linkname keepUnknownTagFunc github.com/valyala/fasttemplate.keepUnknownTagFunc
func keepUnknownTagFunc(w io.Writer, startTag, endTag, tag string, m map[string]interface{}) (int, error)

func tagFunc(w io.Writer, tag string) (int, error) {
	if !strings.HasPrefix(tag, "env.") {
		return keepUnknownTagFunc(w, "{", "}", tag, Vars)
	}
	envKey := strings.TrimPrefix(tag, "env.")
	envVal := os.Getenv(envKey)
	if len(envVal) == 0 {
		return 0, nil
	}
	return w.Write([]byte(envVal))
}
