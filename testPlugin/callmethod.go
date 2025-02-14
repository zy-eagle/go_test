package main

import (
	"log"

	"github.com/dullgiulio/pingo"
)

func main() {
	p := pingo.NewPlugin("tcp", "plugin/method")

	log.Printf("%p\n", p)

	p.Start()
	defer p.Stop()

	var resp string
	if err := p.Call("MyPlugin.SayHello", "Tom", &resp); err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
}
