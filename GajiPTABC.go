package main

import "fmt"

func main() {
	var Gb, Gp, Bl, Pt, lama_waktu_lembur int
	fmt.Scan(&Gp, &lama_waktu_lembur)
	Bl = 45000 * lama_waktu_lembur
	Pt = Gp * 55 / 1000
	Gb = Gp + Bl - Pt
	fmt.Println(Gb, "Gaji Bersih")

}
