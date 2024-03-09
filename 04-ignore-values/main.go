package main

import (
	"fmt"
)

func getNames() (string, string) {
	return "Sue", "Bob"
}

func main() {
	// Bob are ignored
	sue, _ := getNames()

	// Нельзя использовать игнорированные переменные
	// Компилятор просто их выбрасывает, в итоговом бинаре их попросту нет
	// compiler: cannot use _ as value or type [InvalidBlank]
	// bob := _

	fmt.Printf("Go, %v!\n", sue)
}
