package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type employee struct {
	ID   string
	Name string
	Age  int
}

func (e *employee) String() string {
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %q, Name: %q, Age: %d", e.ID, e.Name, e.Age)
}

// func (e employee) String() string {
// 	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID: %q, Name: %q, Age: %d", e.ID, e.Name, e.Age)
// }

func TestCreateEmployeeObj(t *testing.T) {
	e := employee{"0", "Bob", 20}
	e1 := employee{Name: "Mike", Age: 30}
	e2 := new(employee)
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("type of e is %T\n", e)
	t.Logf("type of e2 is %T\n", e2)
}

func TestStructOperations(t *testing.T) {
	e := &employee{"0", "Bob", 20}
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
