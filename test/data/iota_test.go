package data

import (
	"encoding/binary"
	"fmt"
	"testing"
)

// CompressType type of compressions supported by rpc
type CompressType uint16

const (
	Raw CompressType = iota
	Gzip
	Snappy
	Zlib
)

func TestIota(t *testing.T) {
	idx := 0
	header := make([]byte, 32)
	binary.LittleEndian.PutUint16(header[idx:], uint16(Snappy))
	fmt.Println("uint16:", uint16(Snappy))
	fmt.Println(header)

	idx += 2
	str := "Get"
	idx += binary.PutUvarint(header, uint64(len(str)))
	copy(header[idx:], str)
	idx += len(str)
	fmt.Println("header:", header)
}
