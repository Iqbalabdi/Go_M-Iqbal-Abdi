package main

import (
	"fmt"
)

func Compare (a,b string) string {
	res := ""
	if (len(a) < len(b)){
		res = a
	} else {
		res = b
	}
	return res
}

func main(){
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
}