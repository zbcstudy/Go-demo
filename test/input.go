package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal int
	var isPass bool
	fmt.Print("请输入名称：")
	val, _ := fmt.Scanln(&name)
	fmt.Println("val:", val)
	fmt.Println("用户输入的名称是：", name)

	fmt.Println("*********************************")
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

}
