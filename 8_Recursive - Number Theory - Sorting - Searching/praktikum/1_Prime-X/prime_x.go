package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {

	check := 2
	N := 5

	if number == 1 {
		return 2
	}else if number == 2 {
		return 3
	}

	cap := int(math.Sqrt(float64(N)))

	for check != number {
		for i := 2; i <= cap; i++ {
			// manfaatkan branching untuk cek setiap iterasi
			if N%i==0 && N != i {
				break
			}
			if N%i!=0 && i == cap {
				check++
				if check == number { return N }
			}
		}
		N+=2
		cap = int(math.Sqrt(float64(N)))
	}

	return 0
}

func main()  {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}