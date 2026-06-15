package main

import "fmt"

// Nama  : Danar Prastika Okka Raharja
// NIM   : 109082500169
// Kelas : S1IF-13-03

func main() {
	var a [1000]int
	var n, data, i, j, temp, jarak int
	var tetap bool

	n = 0
	for {
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
