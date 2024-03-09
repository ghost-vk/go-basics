package main

import (
	"fmt"
)

// Можно через декларацию типа
// type Cb func(a int) int
// func canHaveCb(fn Cb, val int) int {
func canHaveCb(fn func(a int) int, val int) int {
	return fn(val)
}

func add(a int) int {
	return a
}

func main() {
	res := canHaveCb(add, 1)
	fmt.Printf("Go cb: %v", res)
}
