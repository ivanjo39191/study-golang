package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("SUCCESS: Converted %q to %d.\n:", age, n)

}
