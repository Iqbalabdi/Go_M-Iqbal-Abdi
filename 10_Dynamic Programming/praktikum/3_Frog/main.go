package main

import (
	"fmt"
	"math"
)
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func Frog(jumps []int) int {
  // your code here
  lenArr := len(jumps)
	var minimum_cost = make([]int, lenArr)
	for i := range minimum_cost {
		minimum_cost[i] = 10000
	}
	// fmt.Println(minimum_cost)
	minimum_cost[0] = 0
	for i:=0; i<lenArr; i++ {
		for j:=i+1; j<=i+2; j++ {
			if (j<lenArr) {
				minimum_cost[j] = min(minimum_cost[j], minimum_cost[i]+int(math.Abs(float64(jumps[i])-float64(jumps[j]))))
			}
		}
	}
	return minimum_cost[lenArr-1]
}


func main() {
   fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
   fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}