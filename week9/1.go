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
