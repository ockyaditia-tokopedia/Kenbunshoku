package main

import (
	"errors"
	"fmt"
)

func main() {
	for _, i := range [] int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("Error", e)
		} else {
			fmt.Println("Work", r)
		}
	}

	for _, i := range [] int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("Error", e)
		} else {
			fmt.Println("Work", r)
		}
	}

	_, e := f2(42)
	// TODO: ok?
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// In this case we use `&argError` syntax to build
		// a new struct, supplying values for the two
		// fields `arg` and `prob`.
		return -1, &argError{arg, "Can't work with 42"}
	}
	return arg + 3, nil
}
