package structtest

import (
	"log"
)

type Cat struct {
	Color string
	Name  string
}

type BlackCat struct {
	Cat
	Age int
}

func (b *Cat) eat() {
	log.Print("cat eat")
}

func (b *BlackCat) eat() {
	b.Name = "Tom"
	log.Print("blackcat eat")
}

// 基于非指针类型的接收器方法内部修改字段值，不会影响原来实例字段的值
func (b BlackCat) eat1() {
	b.Name = "BTom"
	log.Print("blackcat eat ", b.Name)
}

func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	cat.Age = 10
	cat.Name = "ss"
	cat.Cat.Name = "ff"
	return cat
}
