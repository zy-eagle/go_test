package goroutine_test

import (
	"sync"
	"testing"
	"time"
)

func TestSequenceExe(t *testing.T) {
	var wg sync.WaitGroup

	var ch1 = make(chan struct{})
	var ch2 = make(chan struct{})
	var ch3 = make(chan struct{})

	wg.Add(4)

	go func() {
		t.Log("1")
		wg.Done()
		ch1 <- struct{}{}
	}()

	go func() {
		<-ch1
		t.Log("2")
		wg.Done()
		ch2 <- struct{}{}
	}()

	go func() {
		<-ch2
		t.Log("3")
		wg.Done()
		ch3 <- struct{}{}
	}()

	go func() {
		<-ch3
		t.Log("4")
		wg.Done()
	}()

	wg.Wait()
}

func TestGoroutine(t *testing.T) {
	var mut sync.Mutex
	sum := 0
	for i := 0; i < 1000000; i++ {
		go func(j int) {
			defer func() {
				mut.Unlock()
			}()
			//t.Log(j)
			mut.Lock()
			sum++
		}(i)
	}
	time.Sleep(1 * time.Second)
	t.Logf("sum = %d", sum)
}

func TestGoroutineWithWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	sum := 0
	//wg.Add(1000)
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(j int) {
			defer func() {
				mut.Unlock()
			}()
			//t.Log(j)
			mut.Lock()
			sum++
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("sum = %d", sum)
}
