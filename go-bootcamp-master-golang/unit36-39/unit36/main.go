package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("42")
	if err == nil {
		fmt.Println("There is no error, n is", n)
	}
}
