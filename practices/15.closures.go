package main

import "fmt"

func main() {
	val := intSeq()
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())

	newVal := intSeq()
	fmt.Println(newVal())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
