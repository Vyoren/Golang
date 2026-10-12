package main

import "fmt"

func main() {
	for _, p := range []bool{false, true} {
		for _, q := range []bool{false, true} {
			for _, r := range []bool{false, true} {
				for _, s := range []bool{false, true} {
					for _, t := range []bool{false, true} {

						hasil1 := ((!p || q) || (!p || r)) == (!p || (q && r))
						hasil2 := !((p && q) || !r) || (!(p == r) && !q)
						hasil3 := !((!p && q) && (!r || p) && (!r || s) && (!s || t)) || t
						hasil4 := !((!p || q) && (p || r) && (!r || s)) || (q || s)
						hasil5 := !((!p || q) && (!r || s) && (p || !s)) || (!r || q)
						hasil6 := !(p != q) || ((p && !q) || (!p && q))
						hasil7 := !(p != q) || ((p || q) && (!q && !p))
						hasil8 := (!(p != q) || !(p == q))
						hasil9 := !(p == q) == ((p && !q) || (!p && q))
						hasil10 := ((p || q) && (!p && !q)) == ((p && !q) || (!p && q))

						fmt.Println(p, q, r, s, t, hasil1, "hasil1", hasil2, "hasil2", hasil3, "hasil3", hasil4, "hasil4", hasil5, "hasil5", hasil6, "hasil6", hasil7, "hasil7", hasil8, "hasil8", hasil9, "hasil9", hasil10, "hasil10")

					}
				}
			}
		}
	}

}
