package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var karakter rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &karakter)
		if karakter == '.' {
			break
		}
		if karakter != ' ' && karakter != '\n' && karakter != '\r' {
			t[*n] = karakter
			*n++
		}
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("teks: ")
	isiArray(&tab, &m)

	isPal := palindrom(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse Teks: ")
	for i := 0; i < m; i++ {
		fmt.Printf("%c", tab[i])
	}

	fmt.Printf("\nPalindrome? %v\n", isPal)
}
