package main

import "fmt"

func hitungSkor(jmlSoal *int, totalSkor *int) {
	var menit int
	*jmlSoal = 0
	*totalSkor = 0

	for i := 1; i <= 8; i++ {
		fmt.Scan(&menit)
		if menit <= 300 {
			*jmlSoal = *jmlSoal + 1
			*totalSkor = *totalSkor + menit
		}
	}
}

func main() {
	var nama, namaPemenang string
	var soal, skor int
	var soalTerbanyak, skorTerkecil int

	soalTerbanyak = -1
	skorTerkecil = 999999

	for {
		fmt.Println("Input Nama & Menit: ")
		fmt.Scan(&nama)
		if nama == "Selesai" || nama == "selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > soalTerbanyak {
			soalTerbanyak = soal
			skorTerkecil = skor
			namaPemenang = nama
		} else if soal == soalTerbanyak {
			if skor < skorTerkecil {
				skorTerkecil = skor
				namaPemenang = nama
			}
		}
	}
	fmt.Println(namaPemenang, soalTerbanyak, skorTerkecil)
}
