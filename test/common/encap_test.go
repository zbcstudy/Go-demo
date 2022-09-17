package common

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (emp Employee) String() string {
	return fmt.Sprintf("ID:%s--name:%s--age:%d", emp.Id, emp.Name, emp.Age)
}

//这种方式在调用时 没有对象复制
func (emp *Employee) ToString() string {
	return fmt.Sprintf("ID:%s--name:%s--age:%d", emp.Id, emp.Name, emp.Age)
}

func TestObjectCreate(t *testing.T) {
	e := Employee{"1", "2", 3}
	t.Log(e.String())
	//指针也可以直接访问
	e2 := &Employee{"2", "3", 4}
	t.Log(e2.String())

	//指针和实例都可以直接访问
	t.Log(e.ToString())
	t.Log(e2.ToString())
}

func TestStructOperations(t *testing.T) {
	e := Employee{"1", "2", 3}
	fmt.Printf("address is %x", unsafe.Pointer(&e))
	t.Log(e.ToString())
}
