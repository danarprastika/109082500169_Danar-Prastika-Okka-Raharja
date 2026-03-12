package main

import "fmt"

func main() {
	var warna1, warna2, warna3, warna4 string
	var benar int
	var status bool

	for k := 1; k <= 5; k++ {
		fmt.Print("Percobaan ", k, ": ")
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu" {
			benar++
		}
	}

	status = (benar == 5)
	fmt.Println("BERHASIL:", status)
}