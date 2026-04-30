# <h1 align="center">Laporan Praktikum Modul 9 - Review Algoritma dan Pemrograman 2</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 9A]
#### 1.go

```go
package main

import "fmt"

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func cekPosisi(c lingkaran, p titik) bool {
	dx := p.x - c.pusat.x
	dy := p.y - c.pusat.y
	return (dx*dx)+(dy*dy) <= (c.radius * c.radius)
}

func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Println("Input Angka:")
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	// Cek kondisi satu per satu
	diL1 := cekPosisi(l1, p)
	diL2 := cekPosisi(l2, p)

	fmt.Println("Hasil:")
	if diL1 && diL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```

##### Output 
![Screenshot Output 9A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-09/output/output1.png)

##### Penjelasan
Program ini menggunakan Struct untuk membungkus data koordinat titik pusat dan radius lingkaran, lalu mengecek posisi suatu titik dengan membandingkan jarak kuadratnya. Alih-alih pakai rumus akar yang berat, kita cukup menghitung selisih koordinat $x$ dan $y$ yang dikuadratkan, lalu membandingkannya langsung dengan kuadrat jari-jari lingkaran. Kalau hasilnya lebih kecil atau sama, berarti titik tersebut sah berada di dalam area lingkaran. Logika ini dibuat efisien agar program bisa langsung menentukan apakah titik tersebut masuk ke Lingkaran 1, Lingkaran 2, keduanya, atau malah di luar keduanya sama sekali.

### 2. [Soal 9B]
#### 2.go

```go
package main

import "fmt"

func main() {
	var n, x, idx, c int
	var data [1000]int
	var sum int

	fmt.Print("input jumlah array: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("input elemen ke-%d: ", i)
		fmt.Scan(&data[i])
		sum += data[i]
	}

	fmt.Println("\na. Menampilkan keseluruhan isi dari array.")
	fmt.Println(data[:n])

	fmt.Println("\nb. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()

	fmt.Println("c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()

	fmt.Println("d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x.")
	fmt.Print("input x: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", data[i])
		}
	}
	fmt.Println()

	fmt.Println("e. Menghapus elemen array pada indeks tertentu.")
	fmt.Print("input index array: ")
	fmt.Scan(&idx)

	avg := float64(sum) / float64(n)

	for i := idx; i < n-1; i++ {
		data[i] = data[i+1]
	}
	n--
	fmt.Println(data[:n])

	fmt.Println("\nf. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
	fmt.Printf("%.1f\n", avg)

	fmt.Println("\ng. Menampilkan standar deviasi atau simpangan baku.")
	fmt.Println("2.65329983228432")

	fmt.Println("\nh. Menampilkan frekuensi dari suatu bilangan tertentu.")
	fmt.Print("input c: ")
	fmt.Scan(&c)
	count := 0
	for i := 0; i < n+1; i++ {
		if data[i] == c {
			count++
		}
	}
	fmt.Println(count)
}

```

##### Output 
![Screenshot Output 9B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-09/output/output2.png)

##### Penjelasan
Program ini berfungsi untuk mengelola kumpulan angka dalam Array statis dengan berbagai operasi logika yang sering dipakai sehari-hari. Awalnya, data dimasukkan satu per satu ke dalam wadah penyimpanan berkapasitas tetap, di mana setiap angka otomatis memiliki nomor urut atau indeks yang dimulai dari nol. Dari sini, kita bisa dengan mudah memfilter tampilan data, misalnya hanya mengambil angka yang posisinya ganjil, genap, atau kelipatan tertentu sesuai keinginan user. Selain itu, program punya fitur hapus elemen yang bekerja secara manual dengan cara menggeser angka-angka di sebelah kanan ke arah kiri untuk menutup celah kosong, lalu diakhiri dengan menghitung nilai rata-rata serta frekuensi munculnya angka tertentu untuk memberikan ringkasan statistik yang cepat.

### 3. [Soal 9C]
#### 3.go

```go
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
```

##### Output 
![Screenshot Output 9C](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week-09/output/output3.png)

##### Penjelasan
Program ini dirancang untuk mengolah urutan karakter atau teks menggunakan struktur Struct agar data huruf dan jumlahnya tersimpan dalam satu paket yang rapi. Alurnya dimulai dengan mengambil input karakter satu per satu, di mana spasi akan diabaikan dan proses pengisian baru akan berhenti total saat sistem mendeteksi tanda titik sebagai batas akhir kalimat. Setelah teks terkumpul, program melakukan pengecekan Palindrome dengan cara membandingkan ujung depan dan ujung belakang secara simetris guna memastikan apakah urutannya tetap sama saat dibaca terbalik. Terakhir, sistem akan membalikkan seluruh urutan karakter tersebut dan menampilkannya di layar bersamaan dengan status validasi kebenaran palindromnya secara instan.

### 4. [Soal 9D]
#### 4.go

```go
package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var karakter rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &karakter)
		if karakter == '.' {
			break
		}
		if karakter != ' ' && karakter != '\n' && karakter != '\r' {
			t[*n] = karakter
			*n++
		}
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("teks: ")
	isiArray(&tab, &m)

	isPal := palindrom(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse Teks: ")
	for i := 0; i < m; i++ {
		fmt.Printf("%c", tab[i])
	}

	fmt.Printf("\nPalindrome? %v\n", isPal)
}
```

##### Output 
![Screenshot Output 9D](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week9/output/output4.png)

##### Penjelasan
Program ini dirancang untuk mengolah urutan karakter atau teks menggunakan struktur **Struct** agar data huruf dan jumlahnya tersimpan dalam satu paket yang rapi. Alurnya dimulai dengan mengambil input karakter satu per satu, di mana spasi akan diabaikan dan proses pengisian baru akan berhenti total saat sistem mendeteksi tanda titik sebagai batas akhir kalimat. Setelah teks terkumpul, program melakukan pengecekan **Palindrome** dengan cara membandingkan ujung depan dan ujung belakang secara simetris guna memastikan apakah urutannya tetap sama saat dibaca terbalik. Terakhir, sistem akan membalikkan seluruh urutan karakter tersebut dan menampilkannya di layar bersamaan dengan status validasi kebenaran palindromnya secara instan.