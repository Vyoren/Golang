package main

import "fmt"

func main() {
	var celcius, fahrenheit, kelvin, reamur float64
	fmt.Scan(&celcius)
	fahrenheit = celcius*9/5 + 32
	reamur = celcius * 4 / 5
	kelvin = celcius + 273.15
	fmt.Printf("%.2f\n", fahrenheit)
	fmt.Printf("%.2f\n", reamur)
	fmt.Printf("%.2f\n", kelvin)

}
