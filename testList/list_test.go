package list_test

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	//循环双向链表
	numList := list.New()
	//插入尾部
	numList.PushBack(60)
	//插入头部
	numList.PushFront(59)

	element := numList.PushBack("62")

	numList.InsertAfter(63, element)

	numList.InsertBefore(61, element)

	//移除
	numList.Remove(element)

	t.Logf("%#v", numList.Back())
	t.Logf("%#v", numList.Front())

	//遍历链表
	for i := numList.Front(); i != nil; i = i.Next() {
		t.Log(i.Value)
	}

	var tempList list.List
	t.Logf("%#v", tempList)

}
