package types

import (
	"fmt"
	"strings"
)

func splitStr() {
	url1 := "http://google.com"
	url2 := "https://google.com:8080"

	p1 := strings.Split(url1, ":")
	fmt.Println(p1[0])
	p2 := strings.Split(url2, ":")
	fmt.Println(p2[0])
	if p1[0] == "http" {
		fmt.Println("p1 http")
	}
}
