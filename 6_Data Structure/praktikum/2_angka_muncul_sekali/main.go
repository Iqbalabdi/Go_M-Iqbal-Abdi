package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {

	//split strinng and convert string to int
	var arrNum []int
	splitStr := strings.Split(angka, "")
	for i := 0; i < len(splitStr); i++ {
		strConv, _ := strconv.Atoi(splitStr[i])
		arrNum = append(arrNum, strConv)
	}

	//remove duplicate values
	var result []int
	for i := 0; i < len(arrNum); i++ {
		counter := 0;
		for j := 0; j < len(arrNum); j++ {
			if (arrNum[i] == arrNum[j]) {
				counter++;
			}
		}

		if (counter == 1) {
			result = append(result, arrNum[i])
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("1122334455"))
}