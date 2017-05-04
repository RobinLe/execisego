package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 10 {
		return -1, errors.New("New error with 10")
	}
	return arg, nil
}

type argError struct {
	errCode int
	errMsg  string
}

func (arg *argError) Error() string {
	return fmt.Sprintf("%d - %s", arg.errCode, arg.errMsg)
}

func f2(arg int) (int, error) {
	if arg == 10 {
		return 0, &argError{arg, "can't work with it"}
	}
	return arg, nil
}

func main() {
	for _, i := range []int{3, 4, 10} {
		if r, e := f1(i); e != nil {
			fmt.Println(r, e)
		} else {
			fmt.Println(r, e)
		}
	}
	for _, i := range []int{1, 10, 3} {
		if r, e := f2(i); e != nil {
			fmt.Println(r, e)
		} else {
			fmt.Println(r, e)
		}
	}
}
