package main

import "fmt"

func main() {
	var berat, kg, gr, biayaKg, biayaGr int

	fmt.Print("Berat Parsel per gram: ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gr = berat % 1000

	biayaKg = kg * 10000

	if berat <= 10000 {
		if gr >= 500 {
			biayaGr = gr * 5
		} else {
			biayaGr = gr * 15
		}
	} else {
		biayaGr = 0
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGr)
	fmt.Printf("Total biaya pengiriman: Rp. %d\n", biayaKg + biayaGr)
}