package main

import "fmt"

func main() {
	fmt.Println("Hello" + "World")
	fmt.Println("1+1=", 1+1)
	fmt.Println("100.0/29.0=", 100.0/29.0)

	name := "Ocky"
	fmt.Printf("Hello, I'm %v\n", name)

	var fullName string = "Ocky Aditia Saputra"
	fmt.Printf("Full name is %s\n", fullName)

	old := 24
	fmt.Printf("I'm %#T years old", old)
}
