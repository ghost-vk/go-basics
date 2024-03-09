package main

import (
	"fmt"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Printf(concat("hello", "world"))
}
