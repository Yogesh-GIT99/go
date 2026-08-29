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
