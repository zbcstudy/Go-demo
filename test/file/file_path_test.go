package file

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T) {
	fmt.Println(filepath.Join("/tmp", "rosedb")) // \tmp\rosedb
}
