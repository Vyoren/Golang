package main

import "fmt"

func main() {
	var x1, x2, x3, x4, x5 int
	var y1, y2, y3, y4 int
	fmt.Scan(&x1, &x2, &x3, &x4, &x5)
	y1 = ((x1+x2)%4096)>>6 + 64
	y2 = ((x2+x3)%4096)>>6 + 64
	y3 = ((x3+x4)%4096)>>6 + 64
	y4 = ((x4+x5)%4096)>>6 + 64
	fmt.Printf("%c,%c,%c,%c", y1, y2, y3, y4)

}
