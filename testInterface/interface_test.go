package interface_test

import (
	"log"
	"testing"
)

func TestInterface(t *testing.T) {
	var invoker Invoker

	s := new(Struct)

	invoker = s
	invoker.Call("hello")

	invoker = FuncCaller(func(v interface{}) {
		t.Log("from interface ", v)
	})

	invoker.Call("hello")
}

type Invoker interface {
	Call(p interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	log.Print("from struct ", p)
}

type FuncCaller func(p interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}
