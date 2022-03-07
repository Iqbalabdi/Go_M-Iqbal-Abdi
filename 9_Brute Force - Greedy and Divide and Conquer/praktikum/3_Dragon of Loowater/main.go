package main

import (
   "fmt"
   "sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
 	// your code here
  	sort.Ints(dragonHead)
 	sort.Ints(knightHeight)

	for i := 0; i<len(dragonHead); i++ {
		for j := 0; j<len(knightHeight); j++ {
			if knightHeight[j] >= dragonHead[i] {
				
			}
		}
	}

}

func main() {
   DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
   DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
   DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
   DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
