package main

import "fmt"

func main() {
	for _, tami := range []bool{true, false} {
		for _, reza := range []bool{true, false} {
			for _, anet := range []bool{true, false} {
				for _, manda := range []bool{true, false} {
					for _, jeni := range []bool{true, false} {
						for _, dini := range []bool{true, false} {
							for _, Sam := range []bool{true, false} {
								for _, yuro := range []bool{true, false} {
									for _, mey := range []bool{true, false} {
										for _, aidan := range []bool{true, false} {
											for _, nadip := range []bool{true, false} {
												for _, alfath := range []bool{true, false} { //semua skenario} {
													percobaan_skenario := ((!reza || tami) || (!tami || reza) || (!anet || reza) || (!reza || anet)) == (!reza || (!tami || !anet))
													percobaan_skenario2 := ((!reza || tami) || (!alfath || manda) || (!manda || alfath) || (!tami || reza) || (alfath || tami)) == ((!reza && !tami) && (!alfath && !manda) && (alfath && tami))
													fmt.Println(tami, reza, anet, jeni, dini, manda, Sam, yuro, mey, aidan, nadip, percobaan_skenario, percobaan_skenario2) //nya belum selesai semua, kemungkinan nanti total wanita nya ada 6, lelaki nya adda 6 juga dan membuat kemungkinan siapa aja yang suka dengan siapa juga yang suka bla-bla-bla panjang nantinya.
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}

		}

	}
}
