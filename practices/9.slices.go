package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println(s)

	s = append(s, "4")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)

	fmt.Println(c[2:4])
	fmt.Println(c[2:])
	fmt.Println(c[:4])

	l := make([][]int, 3)
	for i := 0; i < len(l); i++ {
		length := i + 1
		l[i] = make([]int, length)
		for j := 0; j < length; j++ {
			l[i][j] = i + j
		}
	}

	fmt.Println(l)
}
