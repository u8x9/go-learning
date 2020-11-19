package reflect_test

import (
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func checkType(v interface{}, tt *testing.T) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		tt.Log("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		tt.Log("Integer")
	default:
		tt.Log("Other", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	checkType(f, t)
}

type Employee struct {
	ID   string
	Name string `format:"normal"`
	Age  int
}

func (e *Employee) UpdateAge(age int) {
	e.Age = age
}

type Customer struct {
	CookID string
	Name   string
	Age    int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: Value(%[1]v), Type(%[1]T)\n", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Fatal("Failed to get 'Name' failed")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(100)})
	t.Log("Updated age:", e)

}
