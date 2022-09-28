package _interface

type T interface {
}

type InterfaceA interface {
	Test01()
}

type InterfaceB interface {
	Test02()
}

type InterfaceC interface {
	InterfaceA
	InterfaceB
	TestC()
}
