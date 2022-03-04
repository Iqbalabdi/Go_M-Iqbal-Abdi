package main

import (
	"fmt"
)

func FindMinAndMax(numbers []int) string {
	min := int(^uint(0)>>1)
	max := -min
	var inMin int
	var inMax int

	for in, number := range numbers {
		if(min>number){
			min = number
			inMin = in
		}
		if(max<number){
			max = number
			inMax = in
		}
	}
	res := fmt.Sprintf("min: %d index: %d max: %d index: %d",min,inMin,max,inMax)
	return res
}

func main()  {
	fmt.Println(FindMinAndMax([]int{5,7,4,-2,-1,8}))
	fmt.Println(FindMinAndMax([]int{2,-5,-4,22,7,7}))
	fmt.Println(FindMinAndMax([]int{4,3,9,4,-21,7}))
	fmt.Println(FindMinAndMax([]int{-1,5,6,4,2,18}))
	fmt.Println(FindMinAndMax([]int{-2,5,-7,4,7,-20}))
}