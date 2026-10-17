package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string
	fmt.Scan(&huruf)

	huruf = strings.ToLower(huruf)

	if len(huruf) != 1 {
		fmt.Println("bukan huruf")
		return
	}

	vokal := "aeiou"
	alfabet := "abcdefghijklmnopqrstuvwxyz"

	if strings.Contains(vokal, huruf) {
		fmt.Println("vokal")
	} else if strings.Contains(alfabet, huruf) {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}
