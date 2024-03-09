package main

import (
	"fmt"
)

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("Int: %T\n", v)
	case string:
		fmt.Printf("String: %T\n", v)
	default:
		fmt.Printf("Undetected: %T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
