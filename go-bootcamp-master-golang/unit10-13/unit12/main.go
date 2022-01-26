package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	msg := os.Args[1]
	l := len(msg)
	r := strings.Repeat("!", l)
	s := r + msg + r
	fmt.Println(s)

}
