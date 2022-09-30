package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Address  []string `json:"address"`
}

func main() {
	mo := Monster{
		Name:     "zbc123",
		Age:      18,
		Birthday: "2022-10-23",
		Sal:      189.03,
		Address:  []string{"address1", "address2"},
	}
	str, err := json.Marshal(mo)
	if err != nil {
		fmt.Println("encode error")
	}
	// {"name":"zbc123","age":18,"Birthday":"2022-10-23","Sal":189.03}
	fmt.Println(string(str))

	jsonStr := "{\"name\":\"zbc1234\",\"age\":19,\"Birthday\":\"2032-10-23\",\"Sal\":199.03,\"address\":[\"address1\",\"address2\"]}"

	var monster Monster
	err = json.Unmarshal([]byte(jsonStr), &monster)
	if err != nil {
		fmt.Println("decode error")
	}
	fmt.Println(monster)

	// json反序列化为map
	var m map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println("decode to map error")
	}
	fmt.Println("map:", m)

	jsonSlice := "[{\"name\":\"zbc1234\",\"age\":19,\"Birthday\":\"2032-10-23\",\"Sal\":199.03," +
		"\"address\":[\"address1\",\"address2\"]}]"
	var ms []map[string]interface{}
	err = json.Unmarshal([]byte(jsonSlice), &ms)
	if err != nil {
		fmt.Println("decode slice error")
	}
	fmt.Println("decode slice:", ms)

}
