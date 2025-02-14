package panic_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	tempSlice := GetTempSlice()
	for _, tempItem := range tempSlice {
		t.Log("进入slice")
		t.Log(tempItem)
	}
	t.Log("执行完了...")
}

func GetTempSlice() []int {
	return nil
}

func TestLen(t *testing.T) {
	var serverSqlDBConfigMap = make(map[string]int)
	t.Logf("number of database read: %d", len(serverSqlDBConfigMap))
}

func TestDefer2(t *testing.T) {
	defer func() {
		t.Log("1")
	}()
}

func TestPanic(t *testing.T) {
	defer func() {
		//捕捉了异常
		if err := recover(); err != nil {
			t.Log("recover panic", err)
		}
	}()
	panic(errors.New("报错了"))
	os.Exit(0)
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("execute defer")
	}()
	t.Log("execute func")
	fmt.Println("execute func")

	//time.Sleep(1 * time.Second)
	os.Exit(0)
}

func TestDefer1(t *testing.T) {
	startedAt := time.Now()
	//time.Sleep(time.Second * 2)
	//fmt.Printf("1.%v\n", startedAt)
	// 值拷贝
	defer fmt.Printf("3.%v\n", time.Since(startedAt))
	time.Sleep(time.Second * 2)
	fmt.Printf("2.%v\n", startedAt)
}
