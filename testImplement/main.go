package main

func main() {
	ms := &MyStruct{}
	ms.DoSomething()
}

type MyStruct struct {
	MyInterface
}

type MyInterface interface {
	DoSomething() error
}
