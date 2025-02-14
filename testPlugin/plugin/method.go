package main

import (
	"log"

	"github.com/dullgiulio/pingo"
)

type MyPlugin struct {
}

func (m *MyPlugin) SayHello(name string, msg *string) error {
	*msg = "hello " + name
	return nil
}

func main() {
	plugin := &MyPlugin{}
	log.Printf("%p\n", plugin)
	pingo.Register(plugin)
	pingo.Run()
}
