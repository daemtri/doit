package doit

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"strings"
	_ "unsafe"

	"github.com/valyala/fasttemplate"
)

func Run(cmd string) {
	cmd = fasttemplate.New(cmd, "{", "}").ExecuteFuncString(tagFunc)
	log.Println("run: ", cmd)
	err := RunCommand(context.Background(), &RunCommandOptions{
		Command: cmd,
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Env:     os.Environ(),
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func RunString(cmd string) string {
	var data bytes.Buffer
	cmd = fasttemplate.New(cmd, "{", "}").ExecuteFuncString(tagFunc)
	log.Println("run: ", cmd)
	err := RunCommand(context.Background(), &RunCommandOptions{
		Command: cmd,
		Stdin:   os.Stdin,
		Stdout:  &data,
		Stderr:  os.Stderr,
		Env:     os.Environ(),
	})
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes.TrimSpace(data.Bytes()))
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
