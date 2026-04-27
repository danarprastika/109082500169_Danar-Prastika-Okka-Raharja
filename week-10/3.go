package main

import "fmt"

func main() {
	var n int
	var berat [100]float64
	var total, min, max float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
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
		total += berat[i]
	}

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", total/float64(n))
}
