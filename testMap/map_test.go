package map_test

import (
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var sence sync.Map
	sence.Store("tom", 20)
	sence.Store("jake", 10)
	sence.Store("mary", 20)
	sence.Store("kate", 20)
	t.Log(sence.Load("tom"))
	sence.Delete("mary")
	//return true 继续遍历；false终止遍历
	//sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，
	//sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
	sence.Range(func(k, v interface{}) bool {
		t.Log(k, v)
		return true
	})

	tempMap := new(map[int]string)
	t.Logf("%#v", tempMap)

	tempSlice := new([]string)
	t.Logf("%#v", len(*tempSlice))

	tempMap1 := make(map[int]string)
	tempSlice1 := make([]string, 10)
	t.Logf("%#v", tempMap1)
	t.Logf("%#v", tempSlice1)
	t.Logf("%#v", len(tempSlice1))

}
