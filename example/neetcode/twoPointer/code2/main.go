// Given an array of integers numbers that is sorted in non-decreasing order.
// Return the indices (1-indexed) of two numbers, [index1, index2], such that they add up to a given target number target and index1 < index2. Note that index1 and index2 cannot be equal, therefore you may not use the same element twice.
// There will always be exactly one valid solution. Your solution must use O(1) additional space.
// key here is ( 1-indexed), In go array are ( 0-indexed)

package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{1, 2, 3, 4}, 3))

}

func twoSum(numbers []int, target int) []int {

	l := 0
	r := len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]

		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}

}
