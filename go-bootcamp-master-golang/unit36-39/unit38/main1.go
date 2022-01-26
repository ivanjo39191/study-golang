package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var ( // 宣告
		n   int
		err error
	)
	if a := os.Args; len(a) != 2 {
		// only: a variable
		fmt.Println("Gice me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil { // 重新賦值
		// only : a, n and err variables
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		// all the variables in the if statenment
		n *= 2
		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}
	fmt.Printf("n is %d. QQ - you`ve been shadowed ;)\n", n) // 正常
}
