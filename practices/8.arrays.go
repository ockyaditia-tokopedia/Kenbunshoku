package main

import "fmt"

func main() {
	var a [5]int;
	fmt.Println(a)

	b := [5]int{1, 2}
	fmt.Println(b)

	a[2] = 3
	b[4] = 5

	fmt.Println(a, b)

	var c [2][3][4]int

	fmt.Println(len(c))

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			for k := 0; k < len(c[0][0]); k++ {
				c[i][j][k] = i + j + k
			}
		}
	}

	fmt.Println(c)
}
