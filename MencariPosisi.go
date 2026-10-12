package main

import "fmt"

func main() {
	var posisi, posisi0 int
	var kecepatan, waktu int
	fmt.Scan(&posisi0, &kecepatan, &waktu)
	posisi = posisi0 + (kecepatan * waktu)

	fmt.Println(posisi)
}
