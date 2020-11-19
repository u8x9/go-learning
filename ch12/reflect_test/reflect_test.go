package reflect_test

import (
	"errors"
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

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "三"}
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	t.Log("c1 == c2", reflect.DeepEqual(c1, c2))

}

func TestFillNameAndAge(t *testing.T) {
	data := map[string]interface{}{"Name": "张三", "Age": 99}
	e := Employee{}
	if err := fillByMap(&e, data); err != nil {
		t.Fatal(err)
	}
	t.Log(e)

	c := new(Customer)
	if err := fillByMap(c, data); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
func fillByMap(o interface{}, data map[string]interface{}) error {
	if reflect.TypeOf(o).Kind() != reflect.Ptr {
		return errors.New("必须是指针")
	}
	// Elem() 获取指针指向的值
	if reflect.TypeOf(o).Elem().Kind() != reflect.Struct {
		return errors.New("必须是结构体指针")
	}
	if data == nil {
		return errors.New("请指定有效的数据")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range data {
		if field, ok = reflect.ValueOf(o).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(o).Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
