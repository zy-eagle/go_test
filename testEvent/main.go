package main

import (
	"fmt"
)

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("call actor func", param)
}

func OnEvent(param interface{}) {
	fmt.Println("call global Onevent func", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("call global func", param)
}

func main() {
	a := new(Actor)
	eventName := "OnSkill"
	RegistryEvent(eventName, a.OnEvent)
	RegistryEvent(eventName, GlobalEvent)
	RegistryEvent(eventName, OnEvent)

	CallEvent(eventName, 100)
}
