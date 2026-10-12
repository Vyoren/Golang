package main

import "fmt"

func main() {
	m := false
	n := false
	p := false
	hasilrouter := (!m || n) && (!n || p) && (!p || m) && (!n)
	fmt.Println(hasilrouter, "hasil router true semua")

}
