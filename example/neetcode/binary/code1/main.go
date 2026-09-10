// You are given an array of distinct integers nums, sorted in ascending order, and an integer target.
// Implement a function to search for target within nums. If it exists, then return its index, otherwise, return -1.
// Your solution must run in O(logn) time.

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(search(nums, 5))

}

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := left + (right-left)/2

		switch {
		case nums[mid] == target:
			return mid

		case nums[mid] < target:
			left = mid + 1

		default:
			right = mid - 1
		}
	}
	return -1
}
