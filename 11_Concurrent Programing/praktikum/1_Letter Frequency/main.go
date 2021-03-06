package main

import "fmt"

type FreqMap map[string]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[string(r)]++
	}
	return m
}

func ConcurrentFrequency(strings []string) FreqMap {
	sum := FreqMap{}
	count := len(strings)
	results := make(chan FreqMap, count)

	for _, s := range strings {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}

	for i := 0; i < count; i++ {
		for r, freq := range <-results {
			sum[r] += freq
		}
	}
	return sum
}

func main(){
	result := ConcurrentFrequency([]string{"Lorem ipsum dolor sit amet, consectetur adispicing elit"})
	for key, val := range result {
		fmt.Println(key, ":", val)
	}
}