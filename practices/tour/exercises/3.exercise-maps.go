package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	//split := strings.Split(s, " ")
	field := strings.Fields(s)
	m := make(map[string]int)

	for i := 0; i < len(field); i++ {
		if _, ok := m[field[i]]; ok {
			m[field[i]] = m[field[i]] + 1
		} else {
			m[field[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
