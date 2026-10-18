package main

import "fmt"

func main() {
	var i, n, jumlah_langkah, langkah1, langkah2, langkah3, langkah4, langkah5, langkah6, langkah7 int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&langkah1, &langkah2, &langkah3, &langkah4, &langkah5, &langkah6, &langkah7)
		jumlah_langkah = langkah1 + langkah2 + langkah3 + langkah4 + langkah5 + langkah6 + langkah7
		fmt.Println(jumlah_langkah)

	}

}

//for i = 1; i <= n; i++ {
//     ↑          ↑
// mulai dari     sampai orang
// orang ke-1     ke-n (orang terakhir)
