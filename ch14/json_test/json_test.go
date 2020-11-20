package jsontest

import (
	"encoding/json"
	"testing"
)

func TestStructToStr(t *testing.T) {
	e := Employee{
		BasicInfo: BasicInfo{
			Name: "张三",
			Age:  123,
		},
		JobInfo: JobInfo{
			Skills: []string{"C", "Rust", "Go"},
		},
	}
	buf, err := json.Marshal(e)
	if err != nil {
		t.Fatal("Marshal failed:", err)
	}
	t.Log(string(buf))
}

func TestStrToStruct(t *testing.T) {
	s := `{
		"basic_info": {
			"name": "张三",
			"age": 123
		},
		"job_info": {
			"skills": ["C", "Rust", "Go"]
		}
	}`
	var e Employee
	if err := json.Unmarshal([]byte(s), &e); err != nil {
		t.Fatal("Unmarshal failed:", err)
	}
	t.Log(e)
}
