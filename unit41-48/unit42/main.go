package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris": // case 目的1
		fmt.Println("France") // case 條件表達式
		vip := true           //在 case 中宣告，只能在當前的 case 中使用
		fmt.Println("VIP trip?", vip)
	case "Tokyo": // case 目的2
		fmt.Println("Japan") // case 條件表達式
	default:
		fmt.Println("Where?")
	}
}
