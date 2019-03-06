package main

import "fmt"

// Structs are mutable.
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Ocky", 24})
	fmt.Println(person{name: "Ocky"})

	p := person{"Ocky", 24}
	fmt.Println(p.name)
	fmt.Println(p.age)
}
