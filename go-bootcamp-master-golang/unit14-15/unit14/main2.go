package main

import (
	"fmt"
)

func main() {
	// const (
	// 	EST = -5
	// 	MST = -7
	// 	PST = -8
	// )
	// fmt.Println(EST, MST, PST)

	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
	fmt.Println(EST, MST, PST)

}
