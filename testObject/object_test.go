package object_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name string
	Age  int
}

func TestObject(t *testing.T) {
	e := Employee{"Tome", 30}
	e1 := Employee{Name: "Jake", Age: 20}
	e2 := new(Employee) //这里是返回的引用/指针,相当于e:=&Employee{}
	t.Logf("%#v", e2)

	e2.Name = "Marry1"
	e2.Age = 10
	t.Log(e.Name, e1.Name, e2.Name)

	t.Logf("e前 = %s\n", e.Name)
	e.String()
	t.Logf("e后 = %s\n", e.Name)

	fmt.Println("======================================")

	t.Logf("e1前 = %s\n", e1.Name)
	e1.String1()
	t.Logf("e1后 = %s\n", e1.Name)

}

func (e *Employee) String() {
	e.Name = "Tom1"
	fmt.Println(e.Name, e.Age)
}

func (e Employee) String1() {
	e.Name = "Jake1"
	fmt.Println(e.Name, e.Age)
}
