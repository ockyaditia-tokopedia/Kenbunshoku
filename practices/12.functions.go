package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(mul(1, 2, 3))
	div(2, 4)
}

func sum(a int, b int) int {
	return a + b
}

func mul(a, b, c int) int {
	return a * b * c
}

func div(a, b int) {
	c := b / a
	fmt.Println(c)
}
