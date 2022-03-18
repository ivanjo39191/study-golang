package main

import (
	"fmt"
	"strconv"
)

func main() {
	local := 10
	show(local)
	incrWrong(local)
	fmt.Printf("show -> local = %d\n", local)
	local = incr(local)
	show(local)
	show(incrBy(local, 5))
	local = incrBy(local, 5)
	show(local)

	// local, _ = incrByStr(local, "2")
	// show(local)
	// show(incrBy(local, 5))
	// show(local)

}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, factor int) int {
	return n * factor
}

func incrByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)
	return n, err
}
