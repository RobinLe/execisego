package main

import "fmt"

var a int

func main() {
	a = 1
	fmt.Println(a)
	add()
	fmt.Println(a)

}

func add() {
	a++
}
