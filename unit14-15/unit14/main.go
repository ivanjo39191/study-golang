package main

import (
	"fmt"
)

func main() {
	// const (
	// 	monday    = 1
	// 	tuesday   = 2
	// 	wednesday = 3
	// 	thursday  = 4
	// 	friday    = 5
	// 	saturday  = 6
	// 	sunday    = 7
	// )

	// fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

}
