package main

import "fmt"

func main() {
	n := 1
	fmt.Println(n)
	fmt.Println(&n)

	func1(n)
	// n still 1
	fmt.Println(1.3, n)
	fmt.Println(1.4, &n)

	// The `&n` syntax gives the memory address of `n`,
	func2(&n)
	// n replace with 0
	fmt.Println(2.3, n)
	fmt.Println(2.4, &n)
}

func func1(n int) {
	n = 0
	fmt.Println(1.1, n)
	fmt.Println(1.2, &n)
}

// `func2` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*n` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
func func2(n *int) {
	*n = 0
	fmt.Println(2.1, *n)
	fmt.Println(2.2, &n)
}
