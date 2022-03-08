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
	var minimum_cost = make([]int, len(jumps))
	// minimum_cost[1] = int(math.Abs(float64(jumps[1])-float64(jumps[0])))
	for i:= 3; i<=len(jumps); i++ {
		for j :=
	//   jumps1 := minimum_cost[i - 1]  + int(math.Abs(float64(jumps[i]) - float64(jumps[i - 1]))) //40
	//   jumps2 := minimum_cost[i - 2]  + int(math.Abs(float64(jumps[i]) - float64(jumps[i - 2]))) //
	//   if jumps1 < jumps2 {
	// 	  minimum_cost = append(minimum_cost, jumps1)
	//   } else {
	// 	  minimum_cost = append(minimum_cost, jumps2)
	//   }
		minimum_cost[i] = min(minimum_cost[i - 1] + int(math.Abs(float64(jumps[i]) - float64(jumps[i - 1]))),
                              minimum_cost[i - 2] + int(math.Abs(float64(jumps[i]) - float64(jumps[i - 2]))))
  }

	return minimum_cost[len(jumps)-1]
}


func main() {
   fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
   fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}