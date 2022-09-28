package _interface

import (
	"fmt"
	"testing"
)

func TestUsbFunc(t *testing.T) {
	// 创建对象
	computer := Computer{}
	phone := Iphone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}

// 多态数组
func TestUsbArray(t *testing.T) {
	var usbArray [3]Usb
	usbArray[0] = Iphone{Name: "小米"}
	usbArray[1] = Iphone{Name: "苹果"}
	usbArray[2] = Camera{Name: "佳能"}
	fmt.Println(usbArray)
}
