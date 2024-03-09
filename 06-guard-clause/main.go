package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cant divide by zero")
	}

	return a / b, nil
}

func main() {
	// Не ругнется, но в err будет записано: "cant divide by zero"
	res, err := divide(1, 0)

	if err != nil {
		fmt.Printf("Oops, Error!\n%v\n", err)
	} else {
		fmt.Printf("All Fine! res: %v\n", res)
	}
}
