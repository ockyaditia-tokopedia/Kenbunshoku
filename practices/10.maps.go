package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println(m)
	fmt.Println(m["k2"])
	fmt.Println(len(m))

	delete(m, "k2")

	fmt.Println(m)

	/*
	The optional second return value when getting a value from a map indicates if the key was present in the map.
	This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	 */
	_, p := m["k2"]
	fmt.Println(p)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
