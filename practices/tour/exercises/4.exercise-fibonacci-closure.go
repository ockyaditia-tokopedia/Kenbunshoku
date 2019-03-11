package main

import "fmt"

// fibonacci adalah sebuah fungsi yang mengembalikan sebuah fungsi yang
// mengembalikan sebuah integer.
func fibonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z, x, y = x, y, x+y
		// loop 1
		// x, y, z = 0, 1, 0
		// z, x, y = 0, 1, 1
		// return 0

		// loop 2
		// x, y, z = 1, 1, ?
		// z, x, y = 1, 1, 2
		// return 1

		// loop 3
		// x, y, z = 1, 2, ?
		// z, x, y = 1, 2, 3
		// return 1

		// loop 4
		// x, y, z = 2, 3, ?
		// z, x, y = 2, 3, 5
		// return 2

		// loop 5
		// x, y, z = 3, 5, ?
		// z, x, y = 3, 5, 8
		// return 3

		// loop 6
		// x, y, z = 5, 8, ?
		// z, x, y = 5, 8, 13
		// return 5
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
