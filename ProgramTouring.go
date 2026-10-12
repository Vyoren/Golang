package main

import "fmt"

func main() {
	var kecepatan, menit, jam, total_jarak int
	fmt.Scan(&kecepatan)
	total_jarak = 100 + 60 + 170
	menit = (total_jarak + 60) / kecepatan
	jam = menit / 60
	menit = menit % 60
	println(menit, jam)

}
