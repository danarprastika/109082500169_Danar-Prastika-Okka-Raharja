package main

import "fmt"

func main() {
	var n int
	var berat [1000]float64
	var min, max float64

	fmt.Print("Input jumlah Kelinci: ")
	fmt.Scan(&n)

	fmt.Println("Input Berat:")
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])

		if i == 0 {
			min = berat[i]
			max = berat[i]
		}

		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	// Output berat terkecil dan terbesar
	fmt.Printf("Terkecil: %.2f\n", min)
	fmt.Printf("Terbesar: %.2f\n", max)
}
