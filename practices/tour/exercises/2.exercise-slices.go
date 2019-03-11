package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx)
	fmt.Println(dy)

	s := make([][]uint8, dx)

	for y := 0; y < dy; y++ {
		s[y] = make([]uint8, dy)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			//s[x][y] = uint8((x + y) / 2)
			//s[x][y] = uint8(x * y)
			s[x][y] = uint8(x ^ y)
		}
	}

	return s
}

func main() {
	pic.Show(Pic)
}
