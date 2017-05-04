package main

import (
	"fmt"
)

type point struct {
	x, y int
}

type line struct {
	point
	z int
}

type cube struct {
	p1 point
	p2 point
}

func main() {
	// p := line{point{1, 2}, 3}
	// q := p.point
	// fmt.Println(q)
	p := cube{
		point{1, 2},
		point{3, 4},
	}
	fmt.Printf("%+v", p)
}
