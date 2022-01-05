package main

import "fmt"

func main() {
	i := 142
	switch { // 等同 switch True {
	case i > 100:
		fmt.Print("big ")
		fallthrough // 不中斷，繼續往下個 case 執行
	case i > 0:
		fmt.Print("positive ")
		fallthrough // 不中斷，繼續往下個 case 執行
	default:
		fmt.Println("number")
	}
}
