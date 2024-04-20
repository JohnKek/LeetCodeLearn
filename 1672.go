package main

import (
	"fmt"
	"math"
)

func maximumWealth(accounts [][]int) int {
	max := 0
	for i := 0; i < len(accounts); i++ {
		curr := 0
		for j := 0; j < len(accounts[i]); j++ {
			curr += accounts[i][j]
		}
		max = int(math.Max(float64(curr), float64(max)))
	}
	return max
}

func init() {
	nums := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(nums))
}
