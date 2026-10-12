package main

import "fmt"

func main() {
	var fx float64
	fmt.Scan(&fx)
	hasil := 2/(fx-5) - 5
	fmt.Printf("%.0f\n", hasil)

}
