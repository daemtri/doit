package main

import (
	"fmt"
	"os"

	"github.com/duanqy/lant"
)

func main() {
	os.Exit(lant.Main())
}

func init() {
	lant.AddCommand(&Etcdmanager{Name: "some"})
}

type Etcdmanager struct {
	Name string
}

func (e *Etcdmanager) Build() {
	fmt.Println("in build", e)
}

func (e *Etcdmanager) Run() {
	fmt.Println("in run", e)
	lant.Vars["args"] = "-al"
}
