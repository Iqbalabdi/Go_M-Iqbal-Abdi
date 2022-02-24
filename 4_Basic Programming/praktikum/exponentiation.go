package main

import (
	"fmt"
)
func pangkat(base int, pangkat int) int {
	if pangkat == 0 {
		return 1
	}

	var result int = base
	for i:=2; i<=pangkat; i++ {
		result = result * base
	}

	return result
}

func main(){
	fmt.Println(pangkat(2,3))
}