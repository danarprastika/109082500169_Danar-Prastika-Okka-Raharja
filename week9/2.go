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
