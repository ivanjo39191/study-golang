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
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}
	earthquake, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}
	var d string
	switch e := earthquake; {
	case e < 2:
		d = "micro"
	case e >= 2 && e < 3:
		d = "very minor"
	case e >= 3 && e < 4:
		d = "minor"
	case e >= 4 && e < 5:
		d = "light"
	case e >= 5 && e < 6:
		d = "moderate"
	case e >= 6 && e < 7:
		d = "strong"
	case e >= 7 && e < 8:
		d = "major"
	case e >= 8 && e < 10:
		d = "great"
	case e >= 10:
		d = "massive"
	}
	fmt.Printf("%.2f is %s\n", earthquake, d)
}
