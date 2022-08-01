package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println(nums)

	// My solution
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	fmt.Printf("The maximal value is %d\n", max)

	// Miki solution
	max = nums[0]
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
