package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string
	name = "Ivan"
	fmt.Println("name", name, "byte is", len(name))

	name = "譽瑋"
	fmt.Println("name", name, "byte is", len(name))
	fmt.Println("name", name, "length is", utf8.RuneCountInString(name))
}
