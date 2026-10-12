package main

import "fmt"

func main() {
	sandy_memasukkan_kode_dengan_benar := true
	Sandy_dapat_masuk_ke_apartemen := true
	hasil_soal_1 := sandy_memasukkan_kode_dengan_benar && !Sandy_dapat_masuk_ke_apartemen
	fmt.Println(hasil_soal_1, "sandy_memasukkan_kode_dengan_benar tetapi sandy_tidak_bisa_masuk_ke_apartment")

	Sandy_sudah_melakukan_pengecekan_sistem_di_pintu_apartemennya := true
	hasil_soal_2 := Sandy_sudah_melakukan_pengecekan_sistem_di_pintu_apartemennya && !Sandy_dapat_masuk_ke_apartemen
	fmt.Println(hasil_soal_2)

	jika_sudah_melakukan_pengecekan_sistem_di_pintu_apartemennya_dan_memasukan_kode_dengan_benar := false
	hasil_soal_3 := Sandy_dapat_masuk_ke_apartemen != jika_sudah_melakukan_pengecekan_sistem_di_pintu_apartemennya_dan_memasukan_kode_dengan_benar
	fmt.Println(hasil_soal_3)

	p = false
	q = true
	r = true
	hasilp := !q and r 
	fmt.Println(p)
	
}
