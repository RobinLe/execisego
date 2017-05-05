package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
)

func main() {
	p := md5.Sum([]byte("byte"))
	q := hex.EncodeToString(p[:])
	fmt.Println(reflect.TypeOf(q))
	fmt.Println(q)
}
