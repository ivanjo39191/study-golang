package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields(
		"lazy car jumps again and again and again",
	)
	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
}
