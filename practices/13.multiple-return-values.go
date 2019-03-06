package main

import "fmt"

func main() {
	fmt.Println(values())

	a, b := values()
	fmt.Println(a, b)

	_, c := values()
	fmt.Println(c)

	d, _ := values()
	fmt.Println(d)
}

func values() (int, int) {
	return 29, 10
}
