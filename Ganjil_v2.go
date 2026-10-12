package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)
	sisa := angka % 2
	if angka%2 == 0 {
		fmt.Println(false)
	}
	if angka%2 == 1 {
		fmt.Println(true)
	}
	fmt.Println(sisa)

}
