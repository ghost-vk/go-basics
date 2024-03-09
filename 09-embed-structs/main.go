package main

import (
	"fmt"
)

type car struct {
	height float64
	wheels int
}

type truck struct {
	car
	bedHeight float64
}

func main() {
	t := truck{
		car: car{
			height: 2.3,
			wheels: 6,
		},
		bedHeight: 0.4,
	}
	fmt.Printf("Go! %v\n", t.height)
}
