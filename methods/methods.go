package main

import (
	"fmt"
)

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	r.width = 2
	return 2 * (r.width + r.height)
}

func main() {
	rect1 := rect{1, 2}
	fmt.Println(rect1.area(), rect1.perim())
}
