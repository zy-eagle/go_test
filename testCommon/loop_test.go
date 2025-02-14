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
