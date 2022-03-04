package main

import (
	"fmt"
)

// credit https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/

func MaxSequence(arr []int) int {

	max := -int(^uint(0)>>1)
	sum := 0

	for _,el := range arr {
		sum += el
		if max < sum {
			max = sum
		}
		if (sum < 0){
			sum = 0
		}
	}
	return max
}

func main()  {
	fmt.Println(MaxSequence([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(MaxSequence([]int{-2,-5,6,-2,-3,1,5,-6}))
	fmt.Println(MaxSequence([]int{-2,-3,4,-1,-2,1,5,-3}))
	fmt.Println(MaxSequence([]int{-2,-5,6,-2,-3,1,6,-6}))
	fmt.Println(MaxSequence([]int{-2,-5,6,2,-3,1,6,-6}))
}