package structs

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

type animals struct {
	cat string
	dog string
	pig string
}

func pet() {
	pets := animals{dog: "apple"}
	if pets.cat != "" {
		fmt.Printf("%v", pets)
	} else {
		fmt.Println("missing cat")
	}
}
