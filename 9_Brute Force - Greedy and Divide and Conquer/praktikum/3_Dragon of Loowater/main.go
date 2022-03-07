package main

import (
   "fmt"
   "sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
 	// your code here
  	sort.Ints(dragonHead)
 	sort.Ints(knightHeight)
	var counts int = 0
	lenKnight := len(knightHeight)
	lenDragon := len(dragonHead)

	if (knightHeight[lenKnight-1] < dragonHead[lenDragon-1]){
		fmt.Println("knight fall")
		return
	}

	for i := 0; i<lenDragon; i++ {
		for j := 0; j<lenKnight; j++ {
			if (dragonHead == nil) {
				fmt.Println(counts)
				return
			}
			if knightHeight[j] >= dragonHead[i] {
				if dragonHead[i] == 0 {
					return
				}
				counts += knightHeight[j]
				dragonHead[0] = 0
				break
			}
		}
	}
	fmt.Println(counts)
}

func main() {
   DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
   DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
   DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
   DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
