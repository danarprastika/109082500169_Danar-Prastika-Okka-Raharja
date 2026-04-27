package main

import "fmt"

func faktorial(n int) int64 {
	var h int64 = 1
	for i := 2; i <= n; i++ {
		h *= int64(i)
	}
	return h
}

func permutasi(n, r int) int64 {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int64 {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c1, d int
	fmt.Println("Input Angka: ")
	fmt.Scan(&a, &b, &c1, &d)

	fmt.Println("Hasilnya: ")
	fmt.Println(permutasi(a, c1), kombinasi(a, c1))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
