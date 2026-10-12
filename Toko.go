package main

import "fmt"

func main() {
	var sawit, bensin, beras float64
	fmt.Scan(&sawit, &bensin, &beras)
	hargasawit := sawit*50/1000 + sawit
	hargabensin := bensin*50/1000 + bensin
	hargaberas := beras*50/1000 + beras
	fmt.Println("hasil sawit", hargasawit)
	fmt.Println("hasil bensin", hargabensin)
	fmt.Println("hasil beras", hargaberas)

}
