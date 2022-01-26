package main

import "fmt"

func main() {
	i := 10
	switch { // 等同 switch True {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
