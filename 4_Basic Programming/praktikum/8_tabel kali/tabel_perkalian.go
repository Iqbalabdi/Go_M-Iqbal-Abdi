package main

import "fmt"

func cetakTablePerkalian(number int){

	for i:=1; i<=number; i++ {

		for j:=1; j<=number; j++ {
			if i*j<10 {
				fmt.Printf("%d     ", i*j)
			} else if i < 100 {
				fmt.Printf("%d    ", i*j)
			} else if i < 1000 {
				fmt.Printf("%d   ", i*j)
			}
		}
		fmt.Println()
	}
}

func main(){
	cetakTablePerkalian(15)
}