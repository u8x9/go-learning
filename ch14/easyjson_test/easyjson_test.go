package easyjsontest

import "testing"

func TestEasyJSON(t *testing.T) {
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
	if err := e.UnmarshalJSON([]byte(s)); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	buf, err := e.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(buf))
}
