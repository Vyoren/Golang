package main

import "fmt"

func main() {
	for _, p := range []bool{false, true} {
		for _, q := range []bool{false, true} {
			for _, r := range []bool{false, true} {

				hasil1 := ((!p || q) || (p || r)) == (!p || (q && r))
				hasil2 := !((p && q) || !r) || (!(p == r) && !q)

				fmt.Println(p, q, r, hasil1, "hasil 1", hasil2, "hasil 2")
			}
		}
	}
}
