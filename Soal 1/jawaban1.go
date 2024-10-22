package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1 
	for i := n; i >= 1; i-- {
		*hasil *= i
	}
}

func permutasi(n, r int) int {
	var pembilang, penyebut int
	factorial(n, &pembilang)
	factorial(n-r, &penyebut)
	return pembilang / penyebut
}

func kombinasi(n, r int) int {
	var pembilang, penyebut1, penyebut2 int
	factorial(n, &pembilang)               
	factorial(r, &penyebut1)               
	factorial(n-r, &penyebut2)             
	return pembilang / (penyebut1 * penyebut2) 
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
