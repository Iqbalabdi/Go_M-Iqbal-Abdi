package main

import (
	"fmt"
	"sort"
)

type pair struct{
	name string
	count int
}

func MostAppearItem(items[]string) []pair {
	var itemInfo []pair
	sort.Strings(items)

	if len(items) == 0{
		return itemInfo
	}else {
		itemInfo = append(itemInfo,pair{items[0],0})
	}
	
	i := 0

	for _ ,el := range items {
		if el == itemInfo[i].name{
			itemInfo[i].count++
		}else{
			itemInfo = append(itemInfo,pair{el,1})
			i++
		}
	}

	sort.Slice(itemInfo,func (i,j int) bool {
		return itemInfo[i].count < itemInfo[j].count
	})

	return itemInfo
}

func main()  {
	fmt.Println(MostAppearItem([]string{"js","js","golang","ruby","ruby","js","js"}))
	fmt.Println(MostAppearItem([]string{"A","B","B","C","A","A","B","A","D","D"}))
	fmt.Println(MostAppearItem([]string{"football","basketball","tenis"}))
	fmt.Println(MostAppearItem([]string{}))
}