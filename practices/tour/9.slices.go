package main

/*
Slice adalah referensi ke array

Sebuah slice tidak menyimpan data,
ia hanya mendeskripsikan sebuah bagian dari array.
*/

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // akan merubah names[1]
	fmt.Println(a, b)
	fmt.Println(names)
}
