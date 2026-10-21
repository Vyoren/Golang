package main

import "fmt"

func main() {
	var i, n, sisi, luas, keliling int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&sisi)
		luas = sisi * sisi
		keliling = 4 * sisi
		fmt.Println(luas, keliling)
	}
}
