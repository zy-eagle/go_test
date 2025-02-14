package functest

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

//==================================================

type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	log.Printf("InFunc value: %+v\n", inFunc)
	log.Printf("InFunc ptr: %p\n", &inFunc)
	return inFunc
}

func TestValueFunc(t *testing.T) {
	in := Data{
		complax:  []int{1, 2, 3},
		instance: InnerData{10},
		ptr:      &InnerData{1},
	}
	log.Printf("in value: %+v\n", in)
	log.Printf("in ptr: %p\n", &in)
	log.Printf("in complax ptr: %p\n", &(in.complax))
	out := passByValue(in)
	log.Printf("out complax ptr: %p\n", &(out.complax))
	out.complax = append(in.complax, 4)
	log.Printf("out value: %+v\n", out)
	log.Printf("in value: %+v\n", in)
	log.Printf("out ptr: %p\n", &out)
}

//==================================================

func TestFunc(t *testing.T) {
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
	sum(1, 2, 3)
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clean Resource")
	}()
	t.Log("Started")
	//抛出异常
	panic("Fatal error")
}

func returnMutilValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func sum(op ...int) {
	sum := 0
	for _, tempOp := range op {
		sum += tempOp
	}
	fmt.Println("sum", sum)
}
