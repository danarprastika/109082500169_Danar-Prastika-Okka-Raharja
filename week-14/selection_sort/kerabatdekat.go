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
