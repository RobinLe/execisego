package types

import (
	"fmt"
)

// Append slice append test
func Append() {
	p := []int{1, 2, 3}
	fmt.Println(p)
	fmt.Println(p[:1])

	q := append(p[:1], 4)
	fmt.Println(q)
}
