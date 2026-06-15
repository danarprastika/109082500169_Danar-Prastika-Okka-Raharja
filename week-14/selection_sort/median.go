package main

import "fmt"

func main() {
	var a [1000000]int
	var n, data, i, j, temp int

	n = 0
	for {
		fmt.Print("Input: ")
		fmt.Scan(&data)
		if data == -5313 {
			break
		} else if data == 0 {

			// insertion sort
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

			if n%2 != 0 {
				fmt.Println("Hasil:", a[n/2])
			} else {
				fmt.Println("Hasil:", (a[(n/2)-1]+a[n/2])/2)
			}
		} else {
			a[n] = data
			n = n + 1
		}
	}
}
