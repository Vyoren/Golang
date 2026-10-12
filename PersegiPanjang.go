package main

import "fmt"

func main() {
	var panjang, lebar, luas, keliling int
	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	fmt.Println(keliling, luas, "Luas dan Keliling Persegi Panjang")

}
