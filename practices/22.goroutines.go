// A goroutine is a lightweight thread of execution.
package main

import "fmt"

func main() {
	f("string")

	go f("data")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("Done")
}

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i+1)
	}
}
