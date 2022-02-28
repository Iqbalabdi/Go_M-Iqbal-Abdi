package main

import (
	"fmt"
)

func PairSum(arr []int, target int) []int {

	arrLen := len(arr)
	low := 0
	high := arrLen - 1
	var result []int
	for low < high {
		if arr[low] + arr[high] == target {
			result = append(result, low, high)
			break
		}

		if arr[low] + arr[high] < target {
			low = low+1

		} else {
			high = high-1
		}
	}

	return result
}

func main(){
	fmt.Println(PairSum([]int{1,2,3,4,6}, 6))
	fmt.Println(PairSum([]int{1,3,5,7}, 12))
	// PairSum([]int{1,2,3,4,6}, 6)
}