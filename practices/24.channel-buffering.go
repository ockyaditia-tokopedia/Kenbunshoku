package main

import "fmt"

func main() {
	var len = 10
	messages := make(chan string, len)

	var msg = "message" + 12

	for i := 0; i < len; i++ {
		messages <- "message" + i
	}

	for i := 0; i < len; i++ {
		fmt.Println(<-messages)
	}
}
