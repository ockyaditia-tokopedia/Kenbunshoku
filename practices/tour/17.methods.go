package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method
func (v Vertex) Abs() float64 {
	fmt.Println("method")
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Note: method is function
func Abs(v Vertex) float64 {
	fmt.Println("method is function")
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// void
func m() {
	fmt.Println("void")
}

// function
func f() int {
	fmt.Println("function")
	return -1
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	m()
	f := f()
	fmt.Println(f)
}
