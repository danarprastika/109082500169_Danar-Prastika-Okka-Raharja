package main

import "fmt"

func cariFaktor(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		cariFaktor(n, i+1)
	}
}

func main() {
	var n int
	fmt.Println("Input Angka: ")
	fmt.Scan(&n)
	cariFaktor(n, 1)
	fmt.Println()
}
