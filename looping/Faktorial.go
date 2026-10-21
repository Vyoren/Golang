package main

import "fmt"

func main() {
	var i, n, hasil int
	hasil = 1
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		hasil = i * hasil
	}
	fmt.Println(hasil)
}

//ISI KODE GUA YANG SALAH
//var i, n, hasil int
//	fmt.Scan(&n)
//	for i = 1; i <= n; i++ {
//		hasil = i * n
//		fmt.Println(hasil)
