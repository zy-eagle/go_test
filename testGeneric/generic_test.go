package testGeneric

import (
	"fmt"
	"testing"
)

func TestGeneric(t *testing.T) {
	fmt.Println(PtrWithNil(""))
	fmt.Println(PtrWithNil("22"))
	fmt.Println(Ptr(""))
	fmt.Println(Ptr("22"))
}

func PtrWithNil[T comparable](v T) *T {
	var tmp T
	if v == tmp {
		return nil
	}
	return &v
}

func Ptr[T any](v T) *T {
	return &v
}
