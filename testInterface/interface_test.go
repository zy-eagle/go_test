package interface_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type NilS struct {
	name string
}

func TestUnmarsh(t *testing.T) {
	data := struct {
		Name   string `json:"name"`
		UserId int64  `json:"user_id"`
		Age    int    `json:"age"`
	}{
		Name:   "test",
		UserId: 123456789,
		Age:    18,
	}

	b, _ := json.Marshal(data)

	m := make(map[string]interface{})

	_ = json.Unmarshal(b, &m)

	fmt.Println(m)

	m1 := make(map[string]interface{})
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	if err := dec.Decode(&m1); err != nil {
		fmt.Println("decode err ", err)
		return
	}
	fmt.Println(m1["user_id"])
}

func TestNil(t *testing.T) {
	var ns *NilS = nil

	println(IfElse(ns != nil, ns.name))
}

func IfElse(ifv bool, rl string) string {
	if ifv {
		return rl
	}

	return ""
}

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
