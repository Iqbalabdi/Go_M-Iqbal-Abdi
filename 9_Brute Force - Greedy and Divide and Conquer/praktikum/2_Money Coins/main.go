package main

import (
	"fmt"
	"sort"
)

func moneyCoins(money int) []int {
  // your code here
	var converts []int
	var coins = []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	var i int = 0

	for money > 0 {
		if money / coins[i] > 0 {
			converts = append(converts, coins[i])
			money = money - coins[i]
		} else {
			i += 1
		}
	}
  	return converts
}


func main() {
   fmt.Println(moneyCoins(123))   // [100 20 1 1 1]
   fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]
   fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]
   fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]
   fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]
}