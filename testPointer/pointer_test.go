package pointertest

import (
	"testing"
)

func TestPointer(t *testing.T) {
	var country = "中华人民共和国"
	//获取指针
	ptr := &country
	t.Logf("ptr type is %T", ptr)
	t.Logf("ptr address is %p", ptr)
	//指针取数
	value := *ptr
	t.Logf("value type is %T", value)
	t.Logf("value = %s", value)
	var m = map[int]int{}
	m[1] = 10
	t.Log("m = ", m[1])
}
