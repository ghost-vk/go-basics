package main

import (
	"fmt"
)

// Благодаря именнованному return type получается более выразительный код
func getZeroCoords() (x, y int) {
	// compiler: no new variables on left side of := [NoNewVar]
	// x, y := 0, 0

	// Переменные уже определены в начале
	x, y = 0, 0

	// Можно сделать так:
	// return
	// Но это сложно понимать, теряется выразительность кода, лучше делать так:

	return x, y

	// А можно вообще вот так (но только в маленьких функциях, иначе ничего не разобрать):
	// return 0, 0
}

func main() {
	a, b := getZeroCoords()
	fmt.Printf("x: %v, y: %v\n", a, b)
}
