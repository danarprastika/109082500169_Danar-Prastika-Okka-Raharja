package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nrFact int
	factorial(n, &nFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / nrFact
}

func combination(n, r int, hasil *int) {
	var nFact, rFact, nrFact int
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / (rFact * nrFact)
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int

	fmt.Println("Input Angka: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	fmt.Println(p1, c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Println(p2, c2)
}
