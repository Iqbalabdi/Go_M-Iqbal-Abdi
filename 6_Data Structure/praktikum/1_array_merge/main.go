package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {

	check := make(map[string]int)
	appendArray := append(arrayA, arrayB...)
	result := make([]string,0)
	for _, val := range appendArray {
		check[val] = 1
	}

	for letter, _ := range check {
		result = append(result,letter)
	}
	return result
}

func main(){
	fmt.Println(ArrayMerge([]string{"king", "jin", "lee"}, []string{"kayuza", "jin", "feng"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}