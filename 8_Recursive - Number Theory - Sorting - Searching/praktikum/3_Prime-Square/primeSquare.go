package main

import(
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int){
	size := high * wide
	primes := []int{}
	sum := 0
	in := 0

	if start < 2 {
		primes = append(primes,2)
		primes = append(primes,3)
		sum += 5
		size-=2
	}else if start < 3 {
		primes = append(primes,3)
		sum += 3
		size-=1
	}

	for i := start;; i++ {
		cap := int(math.Sqrt(float64(i)))
		for j := 2; j <= cap; j++ {
			if i % j == 0{
				break
			}else if j == cap && i != start{
				sum += i
				primes = append(primes,i)
				size--
			}
		}
		if size == 0{
			break
		}
	}

	for j := wide; j > 0; j-- {
		for k := high; k > 0; k-- {
			fmt.Printf("%d ",primes[in])
			in++
		}
		fmt.Println()
	}
	fmt.Println(sum)
}

func main()  {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}