package main

import "fmt"

func main() {
	var i, n, jumlah_langkah int
	fmt.Scan(&n)
	for i = jumlah_langkah; i <= 2; jumlah_langkah++ {
		fmt.Scan(&jumlah_langkah)
		total_langkah := jumlah_langkah / 7
		fmt.Println(n, total_langkah)

	}

}
