package main

import "fmt"

func main() {
	var studentName string
	var studentScore int
	
	fmt.Scanf("%s", &studentName)
	fmt.Scanf("%d", &studentScore)

	if studentScore >= 80 {
		fmt.Printf("Mahasiswa %s dengan Nilai A", studentName)
	} else if studentScore >= 65 {
		fmt.Printf("Mahasiswa %s dengan Nilai B", studentName)
	} else if studentScore >= 50 {
		fmt.Printf("Mahasiswa %s dengan Nilai C", studentName)
	} else if studentScore >= 35 {
		fmt.Printf("Mahasiswa %s dengan Nilai D", studentName)
	} else if studentScore >= 0 {
		fmt.Printf("Mahasiswa %s dengan Nilai E", studentName)
	} else {
		fmt.Printf("Mahasiswa %s dengan Nilai Invalid", studentName)
	}
}