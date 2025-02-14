package reflect_test

import (
	"errors"
	"log"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(val int) {
	e.Age = val
}

type Consumer struct {
	ConsumerId string
	Name       string
	Age        int
}

func fillBySetting(obj interface{}, setting map[string]interface{}) error {
	log.Print(reflect.TypeOf(obj).Kind())

	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		//指针获取元素类型
		log.Print("==============", reflect.TypeOf(obj).Elem().Kind())
		if reflect.TypeOf(obj).Elem().Kind() != reflect.Struct {
			return errors.New("this first param should be a pointer to the struct type")
		}
	}
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("this first param should be a pointer type")
	}

	if setting == nil {
		return errors.New("setting is nil.")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range setting {
		if field, ok = reflect.ValueOf(obj).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(obj).Elem()
			//vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestReflectFillSetting(t *testing.T) {
	setting := map[string]interface{}{"Name": "Tom", "Age": 10}
	e := Employee{}
	if err := fillBySetting(&e, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Consumer)
	if err := fillBySetting(c, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

func TestReflect(t *testing.T) {
	a := map[string]int{"one": 1, "two": 2, "three": 3}
	b := map[string]int{"one": 1, "two": 2}
	t.Logf("%p-%p", &a, &b)
	//值的比较
	t.Log(reflect.DeepEqual(a, b))

}
