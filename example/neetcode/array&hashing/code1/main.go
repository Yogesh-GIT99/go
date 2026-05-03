// Contains Duplicate O(n), O()
// Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.
package main

import "fmt"

func main() {
	var nums = []int{2, 3, 4, 3, 5}
	fmt.Println(hasDuplicate(nums))

}

func hasDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true

	}

	return false

}
