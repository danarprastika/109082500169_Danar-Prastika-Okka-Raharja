# <h1 align="center">Laporan Praktikum Modul 10 - Review Algoritma dan Pemrograman 2</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 10A]
#### 1.go

```go
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

	fmt.Printf("Terkecil: %.2f\n", min)
	fmt.Printf("Terbesar: %.2f\n", max)
}
```

##### Output 
![Screenshot Output 10A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week10/output/1.png)

##### Penjelasan
Program ini mendata berat anak kelinci dengan menyimpan input ke dalam Array berkapasitas 1000. Melalui satu kali perulangan, sistem membandingkan setiap berat yang masuk untuk menentukan nilai terkecil dan terbesar secara otomatis.

### 2. [Soal 10B]
#### 2.go

```go
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
```

##### Output 
![Screenshot Output 10B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week10/output/2.png)

##### Penjelasan
Program ini mendistribusikan ikan ke dalam wadah sesuai kapasitas menggunakan penyimpanan Array. Sistem menjumlahkan berat tiap ikan dan mencetak totalnya setiap kali wadah penuh atau data habis. Terakhir, program menghitung rata-rata beban dari semua wadah yang terpakai untuk memantau distribusi berat secara akurat.

### 3. [Soal 10C]
#### 3.go

```go
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
```

##### Output 
![Screenshot Output 10C](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week10/output/3.png)

##### Penjelasan
Program ini mendata berat balita menggunakan Array 100 elemen agar data tersimpan rapi. Lewat bantuan prosedur, sistem otomatis mencari berat badan terkecil dan terbesar tanpa perlu dicek satu per satu. Selain itu, fungsi rata-rata memudahkan petugas mengetahui kondisi kesehatan kelompok balita secara instan.