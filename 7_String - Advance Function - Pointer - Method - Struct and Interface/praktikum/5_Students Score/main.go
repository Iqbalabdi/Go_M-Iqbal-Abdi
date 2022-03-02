package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var temp float64 = 0
	for _, val := range s.score {
		temp += float64(val)
	}
	avg := temp/float64(len(s.score))
	return avg
}


func (s Student) Min() (min int, name string) {
	mapper := map[string]int{}
	for idx, val := range s.name {
		mapper[val] = s.score[idx]
	}
	min = mapper[s.name[0]]
	for key, val := range mapper{
		if val < min {
			min = val
			name = key
		}
	}
	return min, name
}


func (s Student) Max() (max int, name string) {
	mapper := map[string]int{}
	for i, val := range s.name {
		mapper[val] = s.score[i]
	}
	max = mapper[s.name[0]]
	for key, val := range mapper{
		if val > max {
			max = val
			name = key
		}
	}
	return max, name
}


func main() {
	var a = Student{}
	for i := 0; i<3; i++ {
		var name string
		fmt.Print("Input " + strconv.Itoa(i+1) + " Student’s Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		fmt.Println(a.name)
		var score int

		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}
	fmt.Println(a.name[1])
	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

}