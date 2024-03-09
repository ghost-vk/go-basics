package main

import (
	"fmt"
)

// Структуры (struct) это что-то типа объекта в js
type myStruct struct {
	prop  string
	inner innerStruct
}

type innerStruct struct {
	count int
}

func main() {
	// Можно определить сразу или потом
	s := myStruct{}
	s.prop = "bla"
	s.inner = innerStruct{
		count: 25,
	}
	fmt.Printf("s.prop: %v, s.inner.prop: %v\n", s.prop, s.inner.count)
}
