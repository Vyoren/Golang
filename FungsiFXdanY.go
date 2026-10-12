package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	A := 3*(x*x) + 10
	B := 1.0 / A
	C := y*10 + 7
	Function := B + C
	fmt.Println(Function)
}
