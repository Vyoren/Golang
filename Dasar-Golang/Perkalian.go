package main

import "fmt"

func main() {
	var i, n, hasil, angka int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&angka)
		hasil = i * angka
		fmt.Print(hasil, " ")
	}

}
