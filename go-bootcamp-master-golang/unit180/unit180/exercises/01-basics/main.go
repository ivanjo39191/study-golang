// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var null *computer
	// compare the pointer variable to a nil value, and say it's nil
	if null == nil {
		fmt.Println("null computer is nil")
	}
	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "apple"}
	// put the apple into a new pointer variable
	applePtr := apple
	// fmt.Println(applePtr)
	// compare the apples: if they are equal say so and print their addresses
	if apple == applePtr {
		fmt.Println("compare the apples")
		fmt.Printf("apple:    %p\napplePtr: %p\n", &apple, &applePtr)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "sony"}
	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if apple != sony {
		fmt.Println("compare apple to sony")
		fmt.Printf("apple:    %p\nsony: %p\n", &apple, &sony)
	}
	// put apple's value into a new ordinary variable
	appleVal := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Println("apple's addresses and appleVal's addresses")
	fmt.Printf("apple:    %p\n", &apple)
	fmt.Printf("appleVal:    %p\n", &appleVal)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *apple == appleVal {
		fmt.Println("apple and appleVal are equal")
	}

	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Println("apple's addresses and appleVal's value")
	fmt.Printf("apple:    %p\n", *apple)
	fmt.Printf("appleVal:    %p\n", appleVal)
	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	// write a func that returns the value that is pointed by the given *computer
	fmt.Println("change sony")
	fmt.Printf("sony:    %p\n", sony.brand)
	// print the returned value
	fmt.Printf("appleVal:    %p\n", valueOf(apple))
	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
}

func change(c *computer, brand string) {
	c.brand = brand
}

func valueOf(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
