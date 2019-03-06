package main

import "fmt"

func main() {
	fmt.Println(sumFunc(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sumFunc(nums...))
}

func sumFunc(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
