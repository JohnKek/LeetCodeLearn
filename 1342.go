package main

import (
	"fmt"
)

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		switch {
		case num%2 == 0:
			num /= 2
		default:
			num -= 1

		}
		count++
	}
	return count
}

func init() {
	nums := 3
	fmt.Println(fizzBuzz(nums))
}
