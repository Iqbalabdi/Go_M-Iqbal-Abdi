package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	newmsg:=""
	for i:=0;i<len(input);i++ {
		newmsg += string((int(input[i]) - 'a'+ offset)%26+'a');
	}
	return newmsg
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(3, "def"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}

