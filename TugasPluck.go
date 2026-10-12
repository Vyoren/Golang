package main

import "fmt"

func main() {
	pluck := false
	qluck := false
	skenario := (!pluck || !qluck)
	fmt.Println(skenario)
	//--------------------------------------------------------Tugas Tambahan Dari AI------------------------------------
	raka_bertugas := true
	seno_bertugas := true
	supervisor := (!raka_bertugas || !seno_bertugas)
	kemungkinanA := (raka_bertugas && seno_bertugas)
	kemungkinanB := (!raka_bertugas && !seno_bertugas)
	kemungkinanC := (raka_bertugas && !seno_bertugas)
	kemungkinanD := (!raka_bertugas && seno_bertugas)
	fmt.Println(supervisor, kemungkinanA, kemungkinanB, kemungkinanC, kemungkinanD)

	kartuA := true
	kartuB := false
	kartu_aktif := (kartuA != kartuB)
	kemungkinanKartuA := (kartuA && kartuB)
	kemungkinanKartuB := (kartuA && !kartuB)
	kemungkinanKartuC := (!kartuA && kartuB)
	kemungkinanKartuD := (!kartuA && !kartuB)
	fmt.Println(kartu_aktif, kemungkinanKartuA, kemungkinanKartuB, kemungkinanKartuC, kemungkinanKartuD)

}
