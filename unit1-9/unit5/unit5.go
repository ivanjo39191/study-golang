package main

import "fmt"

func main() {
	speed := 100 // speed is int
	force := 2.5 // force is float64
	// int(force) 2.5 => 2

	speed = int(float64(speed) * force)
	// float64(100) = 100.0
	// 100.0 * 2.5 = 250.0
	// int(250.0) = 250

	fmt.Println(speed)
}
