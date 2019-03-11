package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append bekerja pada slice yang nil.
	s = append(s, 0)
	printSlice(s)

	// Slice bertambah seperlunya.
	s = append(s, 1)
	printSlice(s)

	// Kita juga bisa menambahkan lebih dari satu elemen sekaligus.
	s = append(s, 2, 3, 4)
	printSlice(s)

	/*
	len=0 cap=0 []
	len=1 cap=2 [0]
	len=2 cap=2 [0 1]
	len=5 cap=8 [0 1 2 3 4]
	*/
	// NOTE: cap menjadi 8
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
