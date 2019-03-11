package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (ve *Vertex) Scale(f float64) { // baca dan assign v lewat ve
	ve.X = ve.X * f
	ve.Y = ve.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	v = Vertex{3, 4}
	(&v).Scale(10) // tidak perlu &v.Scale(10)
	// karena memiliki pointer-receiver
	fmt.Println(v.Abs())
}
