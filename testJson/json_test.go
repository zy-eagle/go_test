package jsontest

import (
	"encoding/json"
	"testing"
	"time"
)

var jsonStr = `{
	"basic_info":{
		"name":"Tom",
		"age":30
	},
	"job_info":{
		"skills":["Java", "C#", "Go"]
	}
}`

func TestReflectJson(t *testing.T) {
	start := time.Now()
	e := new(Employee)
	//json->object
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	t.Log(*e)
	//object->json
	if v, err := json.Marshal(e); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}

	t.Log(time.Since(start))

}

// func TestEasyJson(t *testing.T) {
// 	start := time.Now()
// 	e := new(Employee)
// 	//json->object
// 	err := e.UnmarshalJSON([]byte(jsonStr))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(*e)
// 	//object->json
// 	if v, err := e.MarshalJSON(); err == nil {
// 		t.Log(string(v))
// 	} else {
// 		t.Error(err)
// 	}

// 	t.Log(time.Since(start))
// }
