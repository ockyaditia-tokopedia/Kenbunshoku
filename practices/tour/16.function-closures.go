package main

/*
Fungsi pada Go bisa closure.
Closure adalah sebuah nilai fungsi
yang merujuk variabel dari blok fungsinya.
Fungsi closure bisa mengakses dan mengisi variabel yang dirujuk;
dalam artian fungsi tersebut "terikat" ke variabel.

Sebagai contohnya, fungsi adder mengembalikan sebuah closure.
Setiap closure terikat dengan variabel sum -nya sendiri.
*/

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
