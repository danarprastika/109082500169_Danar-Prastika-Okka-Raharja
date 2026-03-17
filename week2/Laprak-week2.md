# <h1 align="center">Laporan Praktikum Modul 2 - Review Algoritma dan Pemrograman 1</h1>
<p align="center">Danar Prastika Okka Raharja - 109082500169</p>

## Unguided 

### 1. [Soal 2A]
#### soal-2A.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

****
```

##### Output 
![Screenshot Output 2A](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week2/output/output-2A.png)

##### Penjelasan
Program ini melakukan pergeseran nilai secara melingkar. Awalnya, tiga input string disimpan dalam variabel satu, dua, dan tiga. Untuk menukar isinya tanpa ada data yang hilang, nilai satu dipindahkan sementara ke variabel temp. Kemudian, posisi diisi secara berantai: satu mengambil nilai dua, dua mengambil nilai tiga, dan tiga mengambil nilai dari temp. Hasilnya, urutan input berubah dengan posisi pertama berpindah ke posisi paling akhir.

### 2. [Soal 2B]
#### soal-2B.go

```go
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

```

##### Output 
![Screenshot Output 2B](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week2/output/output-2B.png)

##### Penjelasan
Program ini dirancang untuk memvalidasi urutan empat warna spesifik—merah, kuning, hijau, dan ungu—melalui lima kali tahapan input. Keberhasilan diukur dari konsistensi pengguna dalam memasukkan urutan yang tepat pada setiap percobaan. Jika seluruh rangkaian percobaan terpenuhi tanpa ada kesalahan, program akan memberikan nilai balik berupa boolean true, namun jika terdapat satu saja kesalahan, hasilnya adalah false.

### 3. [Soal 2C]
#### soal-2C.go

```go
package main

import "fmt"

func main() {
	var berat, kg, gr, biayaKg, biayaGr int

	fmt.Print("Berat Parsel per gram: ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gr = berat % 1000

	biayaKg = kg * 10000

	if berat <= 10000 {
		if gr >= 500 {
			biayaGr = gr * 5
		} else {
			biayaGr = gr * 15
		}
	} else {
		biayaGr = 0
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGr)
	fmt.Printf("Total biaya pengiriman: Rp. %d\n", biayaKg + biayaGr)
}
```

##### Output 
![Screenshot Output 2C](https://github.com/danarprastika/109082500169_Danar-Prastika-Okka-Raharja/blob/main/week2/output/output-2C.png)

##### Penjelasan
Program ini menghitung biaya kirim parsel dengan membagi berat ke satuan kg dan sisa gram. Jika total berat di atas 10 kg, sisa gram digratiskan. Namun, jika 10 kg atau kurang, sisa gram dikenakan biaya Rp5/gram (untuk sisa >= 500 gr) atau Rp15/gram (untuk sisa < 500 gr). Di akhir, program menampilkan rincian berat dan total biayanya.