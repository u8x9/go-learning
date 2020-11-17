package interface_test

import (
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
