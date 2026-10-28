package main

import "fmt"

func main() {
	var i, n, hasil int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		hasil = hasil + n
		fmt.Print(hasil, " ")
	}
}
//kemungkinan kode ini masih ada yang salah
