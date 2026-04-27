package main

import "fmt"

func main() {
	var klub1, klub2 string
	var skor1, skor2 int
	var daftarHasil [1000]string

	fmt.Print("Klub A : ")
	fmt.Scan(&klub1)
	fmt.Print("Klub B : ")
	fmt.Scan(&klub2)

	i := 0
	for {
		fmt.Printf("Pertandingan %d : ", i+1)
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		if skor1 > skor2 {
			daftarHasil[i] = klub1
		} else if skor2 > skor1 {
			daftarHasil[i] = klub2
		} else {
			daftarHasil[i] = "Draw"
		}
		i++
	}

	fmt.Println()
	for j := 0; j < i; j++ {
		fmt.Printf("Hasil %d : %s\n", j+1, daftarHasil[j])
	}

	fmt.Println("Pertandingan selesai")
}
