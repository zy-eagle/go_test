package main

import (
	"log"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	defer c.L.Unlock()

	for !done {
		c.Wait()
	}

	log.Println(name, "start reading")
}

func write(name string, c *sync.Cond) {
	log.Println(name, "start writing")
	time.Sleep(time.Second)

	c.L.Lock()
	done = true
	c.L.Unlock()

	log.Println(name, "wakes all")

	c.Broadcast()
}
func main() {
	c := sync.NewCond(&sync.Mutex{})

	go read("read1", c)
	go read("read2", c)
	go read("read3", c)
	write("write", c)

	time.Sleep(3 * time.Second)
}
