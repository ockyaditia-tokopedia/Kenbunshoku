package main

/*
Secara bawaan,
pengiriman dan penerimaan ditahan sampai sisi yang lain siap.
Hal ini membolehkan goroutine untuk melakukan sinkronisasi
tanpa melakukan penguncian secara eksplisit
atau menggunakan variabel kondisi.
*/

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)     // deklarasi channel
	go sum(s[:len(s)/2], c) // -9 + 4 + 0 = -5
	go sum(s[len(s)/2:], c) // 7 + 2 + 8 = 17
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}
