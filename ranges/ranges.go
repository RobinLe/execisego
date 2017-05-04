package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if i == 3 {
			fmt.Println(i, num)
		}
	}
	fmt.Println(sum)

	kvs := map[string]string{"a": "A", "b": "B"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}
}
