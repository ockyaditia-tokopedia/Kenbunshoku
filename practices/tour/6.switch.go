package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "mojave":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("OS :", os)
	}
}
