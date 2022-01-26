package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number     :", n)
	fmt.Println("Returnbed error value:", err)
}
