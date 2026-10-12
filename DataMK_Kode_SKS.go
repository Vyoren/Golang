package main

import "fmt"

func main() {
	var MK string = "algoritma dan pemrograman"
	var Kode, SKS int
	fmt.Println("tuliskan kode MK dan SKS")
	fmt.Scan(&Kode, &SKS)
	fmt.Println("kredit mk", Kode, "-", MK, "1 adalah", SKS, "SKS")

}
