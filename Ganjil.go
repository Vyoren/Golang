package main

import (
	"fmt"
)

func main() {

	var testing int
	var salah, benar bool
	benar = true
	salah = false
	fmt.Scan(&testing)
	hasil := testing % 2
	fmt.Println((hasil == 1) == benar || (hasil == 0) == salah)

}
