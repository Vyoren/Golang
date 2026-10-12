package main

import "fmt"

func main() {
	var bola1 string
	var weight1 int
	var bola2 string
	var weight2 int
	var weight int
	var nama_bola string

	bola1 = "biru"
	bola2 = "ungu"
	fmt.Scan(&weight1, &weight2)

	if weight1 > weight2 {
		println(weight1, "winner")
		weight = weight1
		nama_bola = bola1
	} else if weight1 == weight2 {
		println(weight2, "winner")
		weight = weight2
		nama_bola = bola2
	} else {
		//weight1 < weight2
	}

	fmt.Println(weight, nama_bola, "kedua variable ini menandakan bola paling berat dan nama_bola nya")

}
