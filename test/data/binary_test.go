package data

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
)

func TestBinary(t *testing.T) {
	fmt.Println(uint32(10))
	fmt.Println(strconv.FormatInt(10000, 2)) //1001110 0010000
	fmt.Println(strconv.ParseInt("1001110", 2, 32))

	var res [4]byte
	binary.BigEndian.PutUint32(res[:], uint32(10000))
	fmt.Println(res)
	fmt.Println("res:", binary.BigEndian.Uint32(res[:]))

	binary.BigEndian.PutUint32(res[:], uint32(100))
	fmt.Println(res)
}
