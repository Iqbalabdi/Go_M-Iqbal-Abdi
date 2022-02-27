package main

import (
	"fmt"
)

func pow(base, expo int) int {
	if expo == 0 {
		return 1
	}

	var temp int = pow(base, expo/2)

	if expo%2 == 0 {
		return temp*temp
	} else {
		if expo > 0 {
			return base*temp*temp
		} else {
			return temp*temp/base
		}
	}
}

func main(){
	fmt.Println(pow(2,3))
	fmt.Println(pow(5,3))
	fmt.Println(pow(10,2))
	fmt.Println(pow(2,5))
	fmt.Println(pow(7,3))
}