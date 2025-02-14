package main

import "fmt"

type Operation interface {
	Fly()
	Walk(operation Operation)
	JumpWalk()
}

// 可飞行的
type Common struct{}

func (c *Common) Fly() {
	fmt.Println("can fly")
}

func (c *Common) Walk() {
	fmt.Println("can calk")
	c.JumpWalk()
}

func (c *Common) JumpWalk() {
	fmt.Println("can JumpWalk")
}

// 人类
type Human struct {
	Common // 人类能行走
}

func (h *Human) JumpWalk() {
	fmt.Println("Human can JumpWalk")
}

func main() {
	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()
}
