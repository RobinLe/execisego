package main

import (
	"fmt"
)

func vSlices() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, s[2], len(s))
	s = append(s, "1", "2")
	fmt.Println(s)

	p := []int{1, 2, 3}
	fmt.Println(p)

	k := make([][]int, 3)
	fmt.Println(k)

	v := [][]int{{1, 2, 3}}
	fmt.Println(v)
}

func vMaps() {
	m := make(map[int]string)
	m[1] = "abc"
	m[3] = "cde"
	fmt.Println(m, len(m))

	v := m[1]
	_, p := m[3]
	fmt.Println(v, p)
}

func main() {
	vMaps()
	vSlices()
}
