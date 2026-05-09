package main

import (
	"fmt"
)

func main() {
	var nums = []int{0, 3, 2, 5, 4, 6, 1, 1}
	fmt.Println(longestConsecutive(nums))

}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}
	//fmt.Println(set)

	longest := 0

	for _, num := range nums {

		if !set[num-1] {

			current := num
			streak := 1

			for set[current+1] {
				current++
				streak++
			}

			if streak > longest {
				longest = streak
			}
		}

	}
	return longest
}
