package main

import "fmt"

func main() {
	var angka4, angka5, angka6 int
	var benar, salah bool
	benar = true
	salah = false
	fmt.Scan(&angka4, &angka5, &angka6)
	hasil3 := (angka4 > angka5 || angka5 > angka6 || benar)
	hasil4 := (angka4 == angka5 || angka5 == angka6 || salah)
	hasil5 := (angka6 < angka5 || angka5 < angka6 || benar)
	hasil6 := (angka6 < angka5 || angka5 < angka6 || salah)
	fmt.Println(hasil3 == benar, "karena dari kecil ke besar", hasil4 == salah, "karena sama rata dari angka 4 sampai",
		hasil5 == benar, "karena dari besar ke terkecil", hasil6 == salah, "karena ngga berurutan")

}
