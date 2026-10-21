package main

import "fmt"

func main() {
	var i, n int
	var jam, hasil, total float64
	hasil = 0
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&jam)
		hasil = hasil + jam
	}
	total = hasil / float64(n)
	fmt.Println(total)
}
