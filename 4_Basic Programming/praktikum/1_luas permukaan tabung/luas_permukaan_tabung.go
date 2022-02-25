package main

import "fmt"

func main() {

	const Pi = 3.14
	var tinggi float64
	var jari_jari float64

	fmt.Scanf("%f", &tinggi)
	fmt.Scanf("%f", &jari_jari)

	luas_permukaan := (2*Pi*jari_jari) * (jari_jari + tinggi)

	fmt.Println(luas_permukaan)
}