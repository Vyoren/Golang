package main

import "fmt"

func main() {
	var n, j, x, total_jarak int
	fmt.Scan(&n)
	for j = 1; j <= n; j++ {
		fmt.Scan(&x)
		total_jarak = x * 5
		fmt.Println(total_jarak)
	}
}
