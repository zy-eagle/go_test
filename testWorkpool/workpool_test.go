package testWorkpool

import (
	"fmt"
	"testing"
	"time"
)

func testPool(i int, p *WorkerPool) {
	fmt.Printf("run %d \n", i)
	defer fmt.Printf("done %d \n", i)
	if p.Stopped() {
		fmt.Printf(" %d stoped\n", i)
		return
	}
	if i == 6 {
		fmt.Println("sleep 10--- ")
		time.Sleep(time.Duration(5) * time.Second)
		fmt.Println("sleep 10--- done")
	}
	if i == 8 {
		var stop string
		fmt.Printf("stop -- %d : \n", i)
		// fmt.Scanln(&stop)
		stop = "yes"
		if stop == "yes" {
			p.Stop()
			return
		}
	}
	time.Sleep(time.Duration(i) * time.Second)
}

func TestNewPool(t *testing.T) {
	var pool *WorkerPool
	pool = New(3)
	for x := 1; x <= 15; x++ {
		a := x
		pool.Submit(func() {
			testPool(a, pool)
		})
		fmt.Printf("commit task %d\n", x)
	}
	fmt.Printf("commit done\n")

	// pool.StopWait()
	for {
		if pool.IsFullDone() {
			break
		}
		time.Sleep(time.Second * 1)
	}
}
