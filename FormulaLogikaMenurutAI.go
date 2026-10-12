package main

import "fmt"

func main() {
	jumlahTrue := make([]int, 10)
	jumlahFalse := make([]int, 10)

	for _, p := range []bool{false, true} {
		for _, q := range []bool{false, true} {
			for _, r := range []bool{false, true} {
				for _, s := range []bool{false, true} {
					for _, t := range []bool{false, true} {

						// 1. ((p -> q) || (p -> r)) <-> (p -> (q && r))
						hasil1 := ((!p || q) || (!p || r)) == (!p || (q && r))

						// 2. (p && q || !r) -> (!(p <-> r) && !q)
						hasil2 := !((p && q) || !r) || (!(p == r) && !q)

						// 3. ((!p && q) && (r -> p) && (r -> s) && (s -> t)) -> t
						hasil3 := !((!p && q) && (!r || p) && (!r || s) && (!s || t)) || t

						// 4. ((p -> q) && (!p -> r) && (r -> s)) -> (!q -> s)
						hasil4 := !((!p || q) && (p || r) && (!r || s)) || (q || s)

						// 5. ((p -> q) && (r -> s) && (!p -> !s)) -> (r -> q)
						hasil5 := !((!p || q) && (!r || s) && (p || !s)) || (!r || q)

						// 6. (p XOR q) -> ((p && !q) || (!p && q))
						hasil6 := !(p != q) || ((p && !q) || (!p && q))

						// 7. (p XOR q) -> ((p || q) && (!p && !q))
						hasil7 := !(p != q) || ((p || q) && (!p && !q))

						// 8. (p XOR q) -> !(p <-> q)
						hasil8 := !(p != q) || !(p == q)

						// 9. !(p <-> q) <-> ((p && !q) || (!p && q))
						hasil9 := !(p == q) == ((p && !q) || (!p && q))

						// 10. ((p || q) && (!p && !q)) <-> ((p && !q) || (!p && q))
						hasil10 := ((p || q) && (!p && !q)) == ((p && !q) || (!p && q))

						hasil := []bool{
							hasil1, hasil2, hasil3, hasil4, hasil5,
							hasil6, hasil7, hasil8, hasil9, hasil10,
						}

						for i, nilai := range hasil {
							if nilai {
								jumlahTrue[i]++
							} else {
								jumlahFalse[i]++
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("=== KLASIFIKASI FORMULA ===")

	for i := 0; i < len(jumlahTrue); i++ {
		fmt.Printf(
			"Formula %d: True=%d, False=%d -> ",
			i+1, jumlahTrue[i], jumlahFalse[i],
		)

		switch {
		case jumlahFalse[i] == 0:
			fmt.Println("Valid (tautologi), juga satisfiable")
		case jumlahTrue[i] == 0:
			fmt.Println("Kontradiksi")
		default:
			fmt.Println("Contingency, juga satisfiable")
		}
	}
}
