package testtime

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestNumCPU(t *testing.T) {
	var numCPU = runtime.NumCPU()
	//runtime.go
	fmt.Println("本地核心数据：" + strconv.Itoa(numCPU))

}

func TestTime(t *testing.T) {
	//500ms执行一次 ticker.C 只读通道
	ticker := time.NewTicker(time.Millisecond * 500)
	//两秒后执行 stop.C 只读通道
	stop := time.NewTimer(time.Second * 2)
	var i int
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Println("打点计数器结果为：" + strconv.Itoa(i))
		case <-stop.C:
			fmt.Println("stop")
			goto StopHere
		}
	}

StopHere:
	fmt.Println("done")
}
