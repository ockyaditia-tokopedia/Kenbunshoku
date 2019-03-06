package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	rr := &r
	fmt.Println(rr.area())
	fmt.Println(rr.perimeter())
}
