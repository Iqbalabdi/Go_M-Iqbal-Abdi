package main

import (
	"fmt"
)

func primeNumber(number int) bool {
	if number == 1 {
		return false
	}

	var i int = 2

	for i*i <= number {
		if (number%i == 0) {
			return false
		}

		i += 1
	}

	return true;
}

func main(){
	fmt.Println(primeNumber(100000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}