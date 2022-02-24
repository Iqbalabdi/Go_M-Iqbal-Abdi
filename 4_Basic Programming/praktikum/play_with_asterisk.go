package main

import "fmt"

func playWithAsterisk(n int){

	for rows:=1; rows<=n; rows++{

		for j:=n-rows; j>0; j-- {
			fmt.Print(" ")
		}

		for k:=1; k<=rows; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main(){
	playWithAsterisk(5)
}