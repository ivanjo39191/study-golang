package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}

	if city == "Paris" || city == "Lyon" {
		fmt.Println("France")
	} else if city == "Tokyo" {
		fmt.Println("Japan")
	} else {
		fmt.Println("Where?")
	}
}
