package main

import (
	"fmt"
)

type rectangle struct {
	width  int
	height int
}

// Не особо очевидно как это работает, т.е. не пришлось делать r.area = area
// Добавление метода произошло автоматом

// Стоит заметить, что использовать это лучше только для одноразовых Computed Props, но не хранения состояния
// потому что при изменении, например, width, area не изменится автоматически.
//
// Вид сигнатуры отличается от обычной функции: func add(a int) int { ..
func (r rectangle) area() int {
	return r.width * r.height
}

func main() {
	r := rectangle{
		width:  2,
		height: 3,
	}
	fmt.Printf("Go! area(): %v\n", r.area())
}
