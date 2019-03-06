package main

import (
	"fmt"
	"math"
)

const name = "Ocky"

func main() {
	fmt.Println(name)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int(d))

	fmt.Println(math.Sin(30))
	fmt.Println(math.Cos(30))
	fmt.Println(math.Tan(30))
	fmt.Println(math.Sincos(30))
	fmt.Println(math.Abs(math.Sin(30)))
	fmt.Println(math.Round(math.Cos(30)))
	fmt.Println(math.Floor(math.Cos(30)))
	fmt.Println(math.Ceil(math.Cos(30)))

	fmt.Println(math.Cbrt(27)) // akar dari x

	fmt.Println(math.Exp(1))
	fmt.Println(math.Exp2(1))

	fmt.Println(math.Log2(64))
	fmt.Println(math.Log10(100))

	fmt.Println(math.Max(1, 3))
	fmt.Println(math.Min(1, 3))

	fmt.Println(math.Mod(3, 3))
	fmt.Println(math.Remainder(3, 3))
	fmt.Println(math.Mod(19, 3))
	fmt.Println(math.Remainder(19, 3))

	fmt.Println(math.Pow(3, 3))
	fmt.Println(math.Pow10(3))
}
