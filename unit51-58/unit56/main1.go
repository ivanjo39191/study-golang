package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields(
		"lazy car jumps again and again and again",
	)
	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}
}
