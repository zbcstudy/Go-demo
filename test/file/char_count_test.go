package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

type CharCount struct {
	ChCount    int // 英文
	NumCount   int // 数字
	SpaceCount int // 空格
	OtherCount int // 其他
}

func FileCount(file string) CharCount {
	fl, err := os.Open(file)
	if err != nil {
		fmt.Println("打开文件失败")
		return CharCount{}
	}

	defer fl.Close()

	var charCount CharCount

	// 读取文件
	reader := bufio.NewReader(fl)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			// 读取结束
			break
		}

		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				charCount.ChCount++
			case v == ' ' || v == '\t':
				charCount.SpaceCount++
			case v >= '0' && v <= '9':
				charCount.NumCount++
			default:
				charCount.OtherCount++
			}
		}
	}
	return charCount
}

func TestCharCount(t *testing.T) {
	count := FileCount("test.txt")
	fmt.Println(count)
}
