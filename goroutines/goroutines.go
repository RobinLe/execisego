package goroutines

import (
	"fmt"
)

// F print string msg
func F(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// P call F
func P() {
	F("direct First")
	F("direct Second")
	go F("goroutine Third")
	go F("goroutine Fourth")
}
