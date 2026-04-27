package main

import "fmt"

func cetakBintang(n int) {
	if n > 0 {
		cetakBintang(n - 1)
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakBintang(n)
}
