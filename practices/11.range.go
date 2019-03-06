package main

import "fmt"

func main() {
	i := []int{1, 2, 3}
	sum := 0
	for _, num := range i {
		sum += num
	}
	fmt.Println(sum)

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k, v := range "ocky" {
		fmt.Println(k, v) // unicode code points
	}
}
