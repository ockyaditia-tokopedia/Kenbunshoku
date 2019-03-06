package main

import "fmt"

func main() {
	if num := 2; num < 0 {
		fmt.Println(num)
	} else if num == 10 {
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}

	//number := num == 2 ? 1 : 2 // Go doesn't have ternary operation
}
