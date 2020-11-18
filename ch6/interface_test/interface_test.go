package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return `fmt.Println("Hello World!")`
}

type CProgrammer struct {
}

func (c *CProgrammer) WriteHelloWorld() string {
	return `printf("Hello World!\n");`
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	p = new(CProgrammer)
	t.Log(p.WriteHelloWorld())

}

func TestEmptyInterfacee(t *testing.T) {
	var p interface{} = 123
	if v, ok := p.(int); ok {
		t.Log("int:", v)
	}
}

func anyData(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow", v)
	}
}
func TestAnyData(t *testing.T) {
	anyData(123)
	anyData("abc")
	anyData(45.6)
}
