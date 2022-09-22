package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 测试结构体 标签
func TestStructTag(t *testing.T) {
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	// 序列化为json字符串
	marshal, _ := json.Marshal(monster)
	// 不加tag标签 返回的是 {"Name":"牛魔王","Age":500,"R":"芭蕉扇"}
	fmt.Println("序列化后：", string(marshal))
}
