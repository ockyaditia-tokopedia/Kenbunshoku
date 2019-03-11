package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(ve *Vertex, f float64) { // baca dan assign v lewat ve
	ve.X = ve.X * f
	ve.Y = ve.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10) // perlu &v, agar ve menunjuk ke v
	fmt.Println(Abs(v))
}
