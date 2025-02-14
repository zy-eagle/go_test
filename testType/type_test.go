package type_test

import (
	"reflect"
	"testing"
)

type Brand struct {
}

func (b Brand) show(t *testing.T) {
	t.Log("调用了show方法")
}

type FakeBrand Brand

type FordBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
	FordBrand
}

func TestType(t *testing.T) {
	var v Vehicle
	//为什么不同于Java要初始化，才能调用方法
	v.Brand.show(t)

	//通过反射获取类型对象
	ta := reflect.TypeOf(v)

	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		t.Logf("FieldName：%v, FieldType：%v", f.Name, f.Type.Name())
	}

}
