package main

import (
	"flag"
	"fmt"
	"time"
)

var name = flag.String("name", "Tom", "姓名")
var age = flag.Int("age", 18, "年龄")
var married = flag.Bool("married", false, "婚否")
var delay time.Duration

//Init 初始化参数
func init() {
	flag.DurationVar(&delay, "d", 0, "时间间隔")
}

func main() {
	//Init()
	flag.Parse()

	fmt.Printf("args=%s, num=%d, nflag=%d\n", flag.Args(), flag.NArg(), flag.NFlag())

	fmt.Println("name=", *name)
	fmt.Println("age=", *age)
	fmt.Println("married=", *married)
	fmt.Println("delay=", delay)
}
