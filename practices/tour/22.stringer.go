package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type City struct {
	Code string
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (c City) String() string {
	return fmt.Sprintf("%v (%v city)", c.Code, c.Name)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	c := City{"JKT", "Jakarta"}
	fmt.Println(c)
}
