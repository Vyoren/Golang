package main

import "fmt"

func main() {
	var celcius, reamur, fahrenheit, kelvin float64
	fmt.Scan(&celcius)
	reamur = celcius * 4.0 / 5.0
	fahrenheit = celcius*9.0/5.0 + 32.0
	kelvin = celcius + 273.15
	fmt.Println(reamur, fahrenheit, kelvin)
}
