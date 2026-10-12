package main

import "fmt"

func main() {
	var jariJari, luas, keliling float64
	fmt.Scan(&jariJari)
	luas = 3.14 * jariJari * jariJari
	keliling = 2 * 3.14 * jariJari
	fmt.Println(keliling, luas)
}
