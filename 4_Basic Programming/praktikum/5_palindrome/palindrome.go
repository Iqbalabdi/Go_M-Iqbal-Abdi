package main

import "fmt"

func palindrome(input string) bool {

	for i := 0; i < len(input); i++ {
		j := len(input)-1-i
		if input[i] != input[j] {
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("kaido"))
	fmt.Println(palindrome("luffy"))
}
