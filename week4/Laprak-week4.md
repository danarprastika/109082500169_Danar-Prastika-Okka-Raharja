# <h1 align="center">Laporan Praktikum Modul 3 - Review Algoritma dan Pemrograman 2</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 4A]
#### soal-4A.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nrFact int
	factorial(n, &nFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / nrFact
}

func combination(n, r int, hasil *int) {
	var nFact, rFact, nrFact int
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / (rFact * nrFact)
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int

	fmt.Println("Input Angka: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	fmt.Println(p1, c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Println(p2, c2)
}

```

##### Output 
![Screenshot Output 4A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week4/output/output-4A.png)

##### Penjelasan
Program ini menggunakan sistem tiga "robot" kecil: robot factorial sebagai penyedia hitungan dasar, yang hasilnya kemudian dipakai oleh robot permutation dan combination untuk menentukan jawaban akhir. Karena memakai teknik pass by reference, kita cukup memberikan "kunci" atau alamat memori (dengan simbol * dan &) agar setiap prosedur bisa langsung mengisi hasilnya ke tempat yang sudah disiapkan. Cara ini membuat kode jadi lebih rapi dan efisien karena kita tidak perlu menulis ulang rumus yang sama berkali-kali.

### 2. [Soal 4B]
#### soal-4B.go

```go
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

```

##### Output 
![Screenshot Output 4B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week4/output/output-4B.png)

##### Penjelasan
Program ini bertindak seperti juri yang dibantu robot hitungSkor untuk mendata peserta. Robot ini bertugas menghitung jumlah soal yang selesai tepat waktu dan total menitnya. Dengan teknik pass by reference, juri cukup memberikan "catatan" (alamat memori lewat * dan &) agar robot bisa langsung menulis hasilnya di sana. Di akhir, program membandingkan data semua orang untuk menentukan pemenang berdasarkan soal terbanyak atau waktu tercepat.