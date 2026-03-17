# <h1 align="center">Laporan Praktikum Modul 2 - Review Algoritma dan Pemrograman 1</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 3A]
#### soal-3A.go

```go
package main

import "fmt"

func faktorial(n int) int64 {
	var h int64 = 1
	for i := 2; i <= n; i++ {
		h *= int64(i)
	}
	return h
}

func permutasi(n, r int) int64 {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int64 {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c1, d int
	fmt.Println("Input Angka: ")
	fmt.Scan(&a, &b, &c1, &d)

	fmt.Println("Hasilnya: ")
	fmt.Println(permutasi(a, c1), kombinasi(a, c1))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}

****
```

##### Output 
![Screenshot Output Unguided 2A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week3/output/output-3A.png)

##### Penjelasan
Program ini mengimplementasikan konsep matematika diskrit menggunakan subprogram faktorial berbasis int64 untuk menghitung nilai permutasi melalui pembagian faktorial $n$ dengan $(n-r)$, serta menghitung kombinasi secara efisien dengan membagi hasil permutasi tersebut dengan faktorial $r$. Seluruh proses perhitungan dijalankan secara terstruktur dalam fungsi utama untuk menghasilkan keluaran yang akurat sesuai dengan spesifikasi format input-output yang diminta.

### 2. [Soal 3B]
#### soal-3B.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Println("Input Angka: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println("Hasilnya: ")
	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}

```

##### Output 
![Screenshot Output Unguided 2B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week3/output/output-3B.png)

##### Penjelasan
Program ini mengimplementasikan konsep fungsi komposisi matematika melalui pendefinisian tiga subprogram modular yaitu f, g, dan h, yang kemudian diproses menggunakan teknik pemanggilan fungsi bersarang (nested function calls). Dengan struktur ini, nilai input dari pengguna akan diproses secara berurutan mulai dari fungsi yang berada di posisi paling dalam hingga fungsi terluar guna menghasilkan nilai keluaran yang presisi sesuai dengan aturan notasi komposisi yang diminta.

### 3. [Soal 3C]
#### soal-3C.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var x1, y1, r1, x2, y2, r2, tx, ty float64

	fmt.Println("Input Angka: ")
	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&tx, &ty)

	d1 := didalam(x1, y1, r1, tx, ty)
	d2 := didalam(x2, y2, r2, tx, ty)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

##### Output 
![Screenshot Output Unguided 2C](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week3/output/output-3C.png)

##### Penjelasan
Program ini mengimplementasikan logika geometri untuk menentukan posisi koordinat titik terhadap dua buah lingkaran menggunakan subprogram jarak yang menghitung jarak Euclidean serta fungsi didalam yang memvalidasi radius secara boolean. Dengan memanfaatkan struktur kontrol percabangan if-else, sistem dapat mengklasifikasikan secara akurat apakah titik tersebut berada di dalam salah satu lingkaran, di dalam cakupan kedua lingkaran sekaligus, atau justru berada di luar area keduanya sesuai dengan parameter input yang diberikan.