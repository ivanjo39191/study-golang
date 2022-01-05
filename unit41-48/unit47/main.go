package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch h := t.Hour(); {
	case 6 <= h && h < 11:
		fmt.Println("good morning")
	case 11 <= h && h < 14:
		fmt.Println("good afternoon")
	case 14 <= h && h < 22:
		fmt.Println("good evening")
	case (22 <= h && h < 24) || (0 <= h && h < 6):
		fmt.Println("good night")
	}
}
