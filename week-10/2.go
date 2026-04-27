package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64
	var totalWadah, sum float64
	var jumlahWadah int

	fmt.Println("Input jumlah ikan dan kapasitas wadah:")
	fmt.Scan(&x, &y)

	fmt.Println("Input semua berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	fmt.Println("Berat wadah:")
	sum = 0
	for i := 0; i < x; i++ {
		sum += berat[i]

		if (i+1)%y == 0 || i == x-1 {
			fmt.Printf("Wadah ke-%d: %.2f\n", jumlahWadah+1, sum)

			totalWadah += sum
			sum = 0
			jumlahWadah++
		}
	}
	fmt.Print("Rata rata per-wadah:")
	fmt.Printf("\n%.2f\n", totalWadah/float64(jumlahWadah))
}
