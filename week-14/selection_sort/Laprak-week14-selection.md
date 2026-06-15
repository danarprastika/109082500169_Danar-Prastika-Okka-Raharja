# <h1 align="center">Laporan Praktikum Modul 14 - Review Algoritma dan Pemrograman 2</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 14A]
#### 1.go

```go
package main

import "fmt"

func main() {
	var n, m, x, y, data, t int
	var ganjil, genap [1000]int
	var nGanjil, nGenap int

	fmt.Println("Input: ")
	fmt.Scan(&n)

	x = 1
	for x <= n {
		fmt.Scan(&m)
		nGanjil = 0
		nGenap = 0

		y = 0
		for y < m {
			fmt.Scan(&data)
			if data%2 != 0 {
				ganjil[nGanjil] = data
				nGanjil = nGanjil + 1
			} else {
				genap[nGenap] = data
				nGenap = nGenap + 1
			}
			y = y + 1
		}

		var i, j int

		// sort ganjil membesar
		i = 1
		for i <= nGanjil-1 {
			var idx_min int = i - 1
			j = i
			for j < nGanjil {
				if ganjil[idx_min] > ganjil[j] {
					idx_min = j
				}
				j = j + 1
			}
			t = ganjil[idx_min]
			ganjil[idx_min] = ganjil[i-1]
			ganjil[i-1] = t
			i = i + 1
		}

		// sort genap mengecil
		i = 1
		for i <= nGenap-1 {
			var idx_max int = i - 1
			j = i
			for j < nGenap {
				if genap[idx_max] < genap[j] {
					idx_max = j
				}
				j = j + 1
			}
			t = genap[idx_max]
			genap[idx_max] = genap[i-1]
			genap[i-1] = t
			i = i + 1
		}

		// cetak
		j = 0
		for j < nGanjil {
			fmt.Print(ganjil[j], " ")
			j = j + 1
		}
		j = 0
		for j < nGenap {
			fmt.Print(genap[j], " ")
			j = j + 1
		}
		fmt.Println()
		x = x + 1
	}
}
```

##### Output 
![Screenshot Output 14A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-14/output/dekat.png)

##### Penjelasan
Program ini mengurutkan nomor rumah secara membesar (ascending) dengan teknik Selection Sort, di mana setiap elemen diperbandingkan untuk mencari nilai terkecil (idx_min) yang kemudian ditukar (swap) ke posisi terdepan hingga seluruh daftar rumah urut.

### 2. [Soal 14B]
#### 2.go

```go
package main

import "fmt"

func main() {
	var n, m, x, y int
	var a [1000]int
	var idx_min, t int

	fmt.Println("Input: ")
	fmt.Scan(&n)

	x = 1
	for x <= n {
		fmt.Scan(&m)

		y = 0
		for y < m {
			fmt.Scan(&a[y])
			y = y + 1
		}

		// selection sort persis dari modul 14
		var i, j int
		i = 1
		for i <= m-1 {
			idx_min = i - 1
			j = i
			for j < m {
				if a[idx_min] > a[j] {
					idx_min = j
				}
				j = j + 1
			}
			t = a[idx_min]
			a[idx_min] = a[i-1]
			a[i-1] = t
			i = i + 1
		}

		j = 0
		for j < m {
			fmt.Print(a[j], " ")
			j = j + 1
		}
		fmt.Println()

		x = x + 1
	}
}
```

##### Output 
![Screenshot Output 14B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-14/output/kerabat.png)

##### Penjelasan
Program ini membagi data ke dalam dua array (ganjil dan genap), lalu mengurutkan data ganjil secara membesar (ascending) dan data genap secara mengecil (descending) menggunakan Selection Sort sebelum akhirnya mencetak kedua kelompok data tersebut secara berurutan dalam satu baris.