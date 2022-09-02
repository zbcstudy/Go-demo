package pkg

import (
	"fmt"
	"github.com/golang-module/carbon"
	"testing"
)

func TestCarbon(t *testing.T) {
	car := carbon.Now()
	date := carbon.Now().ToString()
	fmt.Println(date)
	fmt.Println(car.ToDateTimeString())
	fmt.Println(car.ToDateString())
	format := car.Format("2006/01/02 15:04:05")
	fmt.Println(format)

}
