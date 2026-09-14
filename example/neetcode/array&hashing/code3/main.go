// Two sum
// Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.
// You may assume that every input has exactly one pair of indices i and j that satisfy the condition.
// Return the answer with the smaller index first.

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{4, 3, 5, 6}, 9))
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, num := range nums {

		diff := target - num
		if idx, ok := m[diff]; ok {

			return []int{idx, i}
		}

		m[num] = i
	}

	return []int{}

}

// Pattern: complement lookup in a hash map. One pass,
// and the ordering matters: check for target - x before inserting x, otherwise you match an element with itself.
// rule: unsorted array + find a pair → hash map; sorted array + find a pair → two pointers.
