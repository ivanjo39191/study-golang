package main

import "fmt"

func main() {

	switch i := 10; { // switch i := 10; true { ，預設 true 可以不寫
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
