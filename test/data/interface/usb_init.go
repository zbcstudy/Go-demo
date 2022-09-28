package _interface

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Iphone struct {
	Name string
}

func (p Iphone) Start() {
	fmt.Println("手机开始工作")
}

func (p Iphone) Stop() {
	fmt.Println("手机结束工作")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机结束工作")
}

type Computer struct {
}

// 接受usb 接口类型的参数
func (cp Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
