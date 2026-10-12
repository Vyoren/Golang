package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	volume_kubus := s * s * s
	fmt.Println(volume_kubus)

	var alas, tinggi int
	fmt.Scan(&alas, &tinggi)
	ruas_segitiga := alas * tinggi / 2
	fmt.Println(ruas_segitiga)

	var usd, idr int
	fmt.Scan(&idr)
	usd = 1
	convert := idr / 15000 * usd
	fmt.Println(convert)

	fmt.Scan(&usd)
	nilai := usd * 15000
	fmt.Println(nilai)

}
