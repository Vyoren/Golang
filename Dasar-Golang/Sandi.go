package main

import "fmt"

func main() {
	var i, n, angka, angka1, angka4, yang_diambil, total int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&angka)
		angka1 = angka / 1000
		angka4 = angka % 10
		yang_diambil = angka1 + angka4 + total
		total = yang_diambil
	}
	fmt.Println(angka, yang_diambil)
}
