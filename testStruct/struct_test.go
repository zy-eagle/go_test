package structtest

import (
	"log"
	"testing"
)

func TestStruct(t *testing.T) {
	cat := NewBlackCat("black")
	log.Printf("%#v\n", cat)
	cat.Name = "Jake"
	log.Printf("%#v\n", cat)
	cat.eat()
	log.Printf("%#v\n", cat)
	cat.eat1()
	cat.Cat.eat()
	log.Printf("%#v\n", cat)
}
