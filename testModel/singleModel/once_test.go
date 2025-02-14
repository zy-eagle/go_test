package once_test

import (
	"log"
	"sync"
	"testing"
)

type Singleton struct {
}

var singletonInstance *Singleton

var once sync.Once

// 单例模式
func GetSingleObj() *Singleton {
	once.Do(func() {
		log.Print("create single instance")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			obj := GetSingleObj()
			t.Logf("%p\n", obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
