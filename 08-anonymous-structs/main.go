package main

import (
	"fmt"
)

func main() {
	myStruct := struct {
		foo string
		bar string
	}{
		foo: "bla",
		bar: "bla-bla",
	}
	fmt.Printf("Go! %v\n", myStruct.foo)
}
