package main

import (
	"fmt"
)

func main() {
	x := 1
	x = 2

	fmt.Printf("y := x\n")

	y := x

	fmt.Printf("x: %v, y: %v\n", x, y)

	x++

	fmt.Printf("x: %v, y: %v // x incremented, but y the same\n", x, y)
}
