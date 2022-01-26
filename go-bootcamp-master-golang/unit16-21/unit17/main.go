package main

import "fmt"

func main() {
	var (
		ops  = 2350
		ok   = 543
		fail = 433
	)
	fmt.Println("total:", ops, "succuss", ok, "/", fail)
	fmt.Printf(
		"total: %d success %d / %d\n", ops, ok, fail,
	)
}
