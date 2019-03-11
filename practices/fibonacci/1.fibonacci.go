package main

import (
	"fmt"
	"time"
)

var maps map[int]int

func main() {
	n := 50
	start1 := time.Now()
	for i := 0; i < n; i++ {
		fmt.Print(fibDivideAndConquer(i), " ")
	}
	elapsed1 := time.Since(start1)

	fmt.Println("\n", elapsed1)

	maps = make(map[int]int)
	start2 := time.Now()
	for i := 0; i < n; i++ {
		fmt.Print(fibDynamicProgramming(i), " ")
	}
	elapsed2 := time.Since(start2)

	fmt.Println("\n", elapsed2)
}

func fibDivideAndConquer(n int) int {
	if n < 2 {
		return n
	} else {
		return fibDivideAndConquer(n-2) + fibDivideAndConquer(n-1)
	}
}

func fibDynamicProgramming(n int) int {
	result := 0

	if _, v := maps[n]; v {
		return maps[n]
	} else if n < 2 {
		result = n
	} else {
		result = fibDynamicProgramming(n-2) + fibDynamicProgramming(n-1)
	}

	maps[n] = result

	return result
}
