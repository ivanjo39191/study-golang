package main

import (
	"fmt"
	"time"
)

func main() {

	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("good evening")
	case h >= 12:
		fmt.Println("good afternoon")
	case h >= 6:
		fmt.Println("good morning")
	default:
		fmt.Println("good night")
	}
}
