package main

import "fmt"

func runningSum(nums []int) []int {
	output := make([]int, len(nums))
	currentSum := 0
	for index, num := range nums {
		currentSum += num
		output[index] = currentSum
	}
	return output
}
func init() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
