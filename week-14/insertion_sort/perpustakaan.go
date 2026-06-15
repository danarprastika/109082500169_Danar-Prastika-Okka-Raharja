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
