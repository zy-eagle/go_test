package testCommon

import (
	"fmt"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	fmt.Println("===================1")

loop:
	for {
		fmt.Println("===================2")
		for i := 0; i < 10; i++ {
			if i == 5 {
				break loop
			}
			time.Sleep(time.Second)
		}
	}

	fmt.Println("===================3")

}

func TestGoto(t *testing.T) {
	fmt.Println("===================1")

loop:
	for j := 0; j < 3; j++ {
		fmt.Println("j=", j)
		for i := 0; i < 10; i++ {
			if i == 3 {
				goto loop // goto会跳转到指定标签处，直接进入下一次循环
			}
			time.Sleep(time.Second)
		}
	}

	fmt.Println("===================2")

}

func TestContinue(t *testing.T) {
	fmt.Println("===================1")

loop:
	for j := 0; j < 3; j++ {
		fmt.Println("j=", j)
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue loop // continue会跳过本次循环，直接进入下一次循环
			}
			time.Sleep(time.Second)
		}

	}

	fmt.Println("===================2")

}

func TestAnd(t *testing.T) {
	fmt.Println(0100 & 0001)
}
