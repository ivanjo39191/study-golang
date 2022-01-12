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
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go aliens
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 9.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}
	var e string
	switch d := args[1]; {
	case d == "micro":
		e = "less than 2.0"
	case d == "very minor":
		e = "2 - 2.9"
	case d == "minor":
		e = "3 - 3.9"
	case d == "light":
		e = "4 - 4.9"
	case d == "moderate":
		e = "5 - 5.9"
	case d == "strong":
		e = "6 - 6.9"
	case d == "major":
		e = "7 - 7.9"
	case d == "great":
		e = "8 - 9.9"
	case d == "massive":
		e = "10+"
	default:
		e = "unknown"
	}
	fmt.Printf("%s's richter scale is %s\n", args[1], e)
}
