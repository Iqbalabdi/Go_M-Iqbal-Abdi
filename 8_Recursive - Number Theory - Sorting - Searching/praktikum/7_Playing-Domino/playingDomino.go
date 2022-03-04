package main

import(
	"fmt"
	"sort"
)

func playingDomino(cards [][]int,deck[]int) interface{} {

	sort.Slice(cards,func (i,j int) bool {
		return cards[i][0]+cards[i][1] > cards[j][0]+cards[j][1]
	})

	for i := 0; i < len(cards); i++ {
		if (deck[0]==cards[i][0]||deck[0]==cards[i][1]||deck[1]==cards[i][0]||deck[1]==cards[i][1]){
			return cards[i]
		}
	}

	return "tutup kartu"
}

func main()  {
	fmt.Println(playingDomino([][]int{[]int{6,5},[]int{3,4},[]int{2,1},[]int{3,3}},[]int{4,3}))
	fmt.Println(playingDomino([][]int{[]int{6,5},[]int{3,3},[]int{3,4},[]int{2,1}},[]int{3,6}))
	fmt.Println(playingDomino([][]int{[]int{6,6},[]int{2,4},[]int{3,6}},[]int{5,1}))
}