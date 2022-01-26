package main

import (
	"fmt"
	"study-golang/unit1"

	"rsc.io/quote"
)

func main() {
	fmt.Println(unit1.UnitOne())
	q := quote.Go()
	fmt.Println(q)
}
