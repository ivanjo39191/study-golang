// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
	} else if n%4 == 0 && (n%400 == 0 || n%100 != 0) {
		fmt.Printf("%d a leap year\n", n)
	} else {
		fmt.Printf("%d is not a leap year\n", n)
	}
}
