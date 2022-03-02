package main

import (
	"fmt"
)

func caesar(offset int, input string) string {

	shift, asci := rune(offset), rune(26)
	// len:=len(input)
	// newmsg:=""
	// input=strings.ToUpper(input)

	runes := []rune(input)
	for index, char := range runes {
		if char >= 'a'+shift && char <= 'z' || char >= 'A'+shift && char <= 'Z' {
			char = char - shift
		} else if char >= 'a' && char < 'a'+shift || char >= 'A' && char < 'A'+shift {
			char = char - shift + asci
		}

		runes[index] = char
	}
	return string(runes)
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}