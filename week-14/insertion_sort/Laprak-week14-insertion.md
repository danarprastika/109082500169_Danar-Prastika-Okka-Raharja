# <h1 align="center">Laporan Praktikum Modul 14 - Review Algoritma dan Pemrograman 2</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 14A]
#### 1.go

```go
package main

import "fmt"

func main() {
	var a [1000]int
	var n, data, i, j, temp, jarak int
	var tetap bool

	n = 0
	for {
		fmt.Print("Input: ")
		fmt.Scan(&data)
		if data < 0 {
			break
		}
		a[n] = data
		n = n + 1
	}

	i = 1
	for i <= n-1 {
		j = i
		temp = a[j]
		for j > 0 && temp < a[j-1] {
			a[j] = a[j-1]
			j = j - 1
		}
		a[j] = temp
		i = i + 1
	}

	i = 0
	for i < n {
		fmt.Print(a[i], " ")
		i = i + 1
	}
	fmt.Println()

	tetap = true
	if n > 1 {
		jarak = a[1] - a[0]
		i = 1
		for i <= n-1 {
			if a[i]-a[i-1] != jarak {
				tetap = false
			}
			i = i + 1
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```

##### Output 
![Screenshot Output 14A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-14/insertion_sort/output/jarak.png)

##### Penjelasan
Di sini, Insertion Sort berperan sebagai "pematang" data agar selisihnya mudah dicek. Setelah diurutkan dari terkecil ke terbesar, program melakukan pengecekan konstansi. Ia mengambil selisih dua angka pertama sebagai referensi jarak, kemudian membandingkan selisih setiap angka berikutnya dengan referensi tersebut. Jika ada satu saja selisih yang meleset, variabel tetap akan bernilai false, menandakan bahwa data tersebut tidak memiliki pola jarak yang tetap.

### 2. [Soal 14B]
#### 2.go

```go
package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func main() {
	var pustaka DaftarBuku
	var n, r int

	fmt.Print("Jumlah buku: ")
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	fmt.Print("Rating yang dicari: ")
	fmt.Scan(&r)

	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, r)
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int = 0
	for i < n {
		fmt.Print("Input buku ke-", i+1, ": ")
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
		i = i + 1
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var max, idx_max, i int

	if n > 0 {
		max = pustaka[0].rating
		idx_max = 0
		i = 1
		for i <= n-1 {
			if pustaka[i].rating > max {
				max = pustaka[i].rating
				idx_max = i
			}
			i = i + 1
		}
		fmt.Println(pustaka[idx_max].judul, pustaka[idx_max].penulis, pustaka[idx_max].penerbit, pustaka[idx_max].tahun)
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	/* I.S. Array pustaka berisi n data buku
	   F.S. Array pustaka terurut menurun/mengecil */
	var temp Buku
	var i, j int

	i = 1
	for i <= n-1 {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i int = 0
	if n < 5 {
		for i < n {
			fmt.Println(pustaka[i].judul)
			i = i + 1
		}
	} else {
		for i < 5 {
			fmt.Println(pustaka[i].judul)
			i = i + 1
		}
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var kr, kn, med int
	var found bool

	found = false
	kr = 0
	kn = n - 1

	for kr <= kn && !found {
		med = (kr + kn) / 2
		if r > pustaka[med].rating {
			kn = med - 1
		} else if r < pustaka[med].rating {
			kr = med + 1
		} else {
			found = true
		}
	}

	if found {
		fmt.Println(pustaka[med].judul, pustaka[med].penulis, pustaka[med].penerbit, pustaka[med].tahun, pustaka[med].eksemplar, pustaka[med].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
```

##### Output 
![Screenshot Output 14B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-14/insertion_sort/output/perpustakaan.png)

##### Penjelasan
Program ini mengelola data buku menggunakan struct dan array, menerapkan alur "urutkan sebelum cari" melalui Insertion Sort untuk menyusun rating secara menurun (descending) dan Binary Search untuk pencarian efisien, serta dilengkapi fitur pencarian buku terfavorit dan daftar lima buku teratas.