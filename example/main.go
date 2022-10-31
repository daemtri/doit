package main

import (
	"fmt"
	"github.com/duanqy/doit"
)

func main() {
	doit.Command(Etcdmanager{})
	doit.Execute(Etcdmanager{})
}

type Etcdmanager doit.Namespace

func (e Etcdmanager) Build() {
	fmt.Println("in build", e)
}

func (e Etcdmanager) Run() {
	fmt.Println("in run", e)
	doit.Vars["args"] = "-al"
}
