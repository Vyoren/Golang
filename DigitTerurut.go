package main

import "fmt"

func main() {
	var angka1, angka2, angka3 int
	var benar, salah bool
	fmt.Scan(&angka1, &angka2, &angka3)
	hasil1 := (angka3 < angka2 || angka2 < angka1 || benar)
	hasil2 := (angka1 == angka2 || angka2 == angka3 || salah)
	fmt.Println(hasil1 == salah, "karena berurutan dari besar ke kecil", hasil2 == benar, "karena sama rata")
}
