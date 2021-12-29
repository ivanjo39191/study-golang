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
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Give me args")
	} else if len(args) == 2 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if len(args) == 3 {
		fmt.Printf("There is two: \"%s %s\"\n", args[1], args[2])
	} else if len(args) == 6 {
		fmt.Println("There are 5 arguments")
	}
}
