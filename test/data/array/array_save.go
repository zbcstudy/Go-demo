package main

import "fmt"

func main() {
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%v的成绩\n", i)
		fmt.Scanln(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Printf("第%v的成绩为%.2f\n", i, score[i])
	}
}
