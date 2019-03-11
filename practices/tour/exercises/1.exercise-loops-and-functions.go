package main

import (
	"fmt"
	"math"
)

func Sqrt(x, y float64) float64 {
	z := x
	for z >= y {
		z -= (z*z - x) / (2 * z)
		fmt.Println("z =", z)
	}
	return z
}

func main() {
	x := 10.0
	fmt.Println(Sqrt(x, math.Sqrt(x)))
	fmt.Println(math.Sqrt(x))
}
