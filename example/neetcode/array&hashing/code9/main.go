// Given an array of integers nums, return the length of the longest consecutive sequence of elements that can be formed.

// A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element. The elements do not have to be consecutive in the original array.

// You must write an algorithm that runs in O(n) time.

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

// Hashset firstO(1) lookup — makes the inner walk O(n) total, not O(n²)!
// set[num-1] gateGuarantees you only walk each sequence once from its head
// Inner walk is O(n) total across all iterationsEach number is visited at most twice — once in outer loop, once in inner
