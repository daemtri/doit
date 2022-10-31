package main

import (
	"fmt"

	"github.com/duanqy/lant"
)

func main() {
	lant.RegisterCommand(Etcdmanager{})
	lant.Execute()
}

type Etcdmanager lant.Namespace

func (e Etcdmanager) Build() {
	fmt.Println("in build", e)
}

func (e Etcdmanager) Run() {
	fmt.Println("in run", e)
	lant.Vars["args"] = "-al"
}
