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
