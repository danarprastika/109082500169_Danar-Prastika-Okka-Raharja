# <h1 align="center">Laporan Praktikum Modul 5 - Review Algoritma dan Pemrograman 2</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 5A]
#### soal-5A.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	fmt.Println("Input Angka: ")
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

```

##### Output 
![Screenshot Output 5A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-05/output/output-5A.png)

##### Penjelasan
Program mencari nilai Fibonacci dengan cara fungsi "bertanya" pada dua angka sebelumnya ($n-1$ dan $n-2$). Proses ini terus turun hingga menyentuh base-case (0 atau 1), lalu hasilnya dijumlahkan kembali ke atas. Rekursif mempermudah penulisan rumus matematika yang kompleks menjadi kode yang sangat singkat.

### 2. [Soal 5B]
#### soal-5B.go

```go
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

```

##### Output 
![Screenshot Output 5B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-05/output/output-5B.png)

##### Penjelasan
Program mencetak segitiga bintang dengan teknik "menabung" tugas. Karena fungsi rekursif dipanggil sebelum perintah cetak, program akan antre di memori (stack) hingga angka terkecil, baru kemudian mulai mencetak bintang dari baris 1 naik hingga baris ke-$n$. Ini membuktikan rekursif memiliki sistem antrean tugas yang sangat teratur.

### 3. [Soal 5C]
#### soal-5C.go

```go
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
```

##### Output 
![Screenshot Output 5C](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-05/output/output-5C.png)

##### Penjelasan
Fungsi rekursif di sini bertindak sebagai pemeriksa otomatis. Dimulai dari angka 1, fungsi mengecek apakah angka tersebut habis membagi $N$. Jika iya, angka dicetak, lalu fungsi memanggil dirinya sendiri untuk mengecek angka berikutnya ($i+1$) sampai mencapai batas $N$. Teknik ini menggantikan perulangan for dengan logika yang lebih modular.