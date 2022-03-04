package main

import(
	"fmt"
	"sort"
)

// credit https://www.geeksforgeeks.org/quick-sort/

func MaximumBuyProduct(money int, productPrice []int)  {

	sort.Slice(productPrice,func (i,j int) bool {
		return productPrice[i] < productPrice[j]
	})

	var product,sum int

	for _,el := range productPrice{
		sum += el
		if(sum > money){
			break
		}else{
			product++
		}
	}

	fmt.Println(product)
}

func main()  {
	MaximumBuyProduct(50000,[]int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000,[]int{15000,10000,12000,5000,3000})
	MaximumBuyProduct(10000,[]int{2000,3000,1000,2000,10000})
	MaximumBuyProduct(4000,[]int{7500,3000,2500,2000})
	MaximumBuyProduct(0,[]int{10000,30000})
}
