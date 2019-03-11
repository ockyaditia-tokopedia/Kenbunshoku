package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("Fav number", rand.Intn(10))
}
