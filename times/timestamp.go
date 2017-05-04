package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(timestamp)
	p, _ := strconv.ParseInt("1490785460", 10, 64)

	tm := time.Unix(p, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))
}
