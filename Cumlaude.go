package main

import "fmt"

func main() {
	var syarat_semester, syarat_nilai int
	var lulus, tidak_lulus bool
	lulus = true
	tidak_lulus = false
	fmt.Scan(&syarat_semester, &syarat_nilai)
	hasil1 := ((syarat_semester > 8 == tidak_lulus) || (syarat_semester < 8 == lulus)) && ((syarat_nilai < 500 == tidak_lulus) || (syarat_nilai > 500 == lulus))
	fmt.Println(hasil1)

}
