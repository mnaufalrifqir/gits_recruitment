package main

import (
	"fmt"
	"sort"
)

func denseRanking(scores []int, score int) int {
	var ranks []int
	var result int

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	rank := 1
	for i := 0; i<len(scores); i++{
		if i != len(scores)-1 {
			if scores[i] != scores[i+1] {
				ranks = append(ranks, rank)
				rank++
			} else {
				ranks = append(ranks, rank)
			}
		} else {
			ranks = append(ranks, rank)
		}
	}

	for i, s := range scores {
		if s == score {
			result = ranks[i]
		}
	}
	return result
}

func main() {
	var nPlayer, score, nGITS int
	var scores, rankResult []int

	fmt.Scanln(&nPlayer)

	for i := 0; i < nPlayer; i++ {
		fmt.Scan(&score)
		scores = append(scores, score)
	}
	fmt.Scanln()

	fmt.Scanln(&nGITS)

	for i := 0; i < nGITS; i++ {
		fmt.Scan(&score)
		scores = append(scores, score)
		rankResult = append(rankResult, denseRanking(scores, score))
	}
	fmt.Scanln()

	for _, rank := range rankResult {
		fmt.Print(rank, " ")
	}
}
