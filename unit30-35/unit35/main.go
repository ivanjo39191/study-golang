package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: [feet]")
		return
	}
	arg := os.Args[1]
	f, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("Error: %q is not a number.\n", arg)
		return
	}
	m := f * 0.3048
	fmt.Printf("%g feet is %g meters.\n", f, m) // %g 根據情況顯示科學計數法或浮點數
}
