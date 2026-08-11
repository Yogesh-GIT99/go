// You are given an array of integers nums and an integer k. There is a sliding window of size k that starts at the left edge of the array.
// The window slides one position to the right until it reaches the right edge of the array.
// Return a list that contains the maximum element in the window at each step.
package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 0, 4, 2, 6}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 || k <= 0 {
		return []int{}
	}

	dq := make([]int, 0, k)
	result := make([]int, 0, len(nums)-k+1)

	for i, num := range nums {

		// remove the front element if it is out of window
		if len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}

		// remove element from back if it is smaller than new element
		if len(dq) > 0 && nums[dq[len(dq)-1]] < num {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i) // add new element index to dq back

		// add result once window is full
		if i >= k-1 {
			result = append(result, nums[dq[0]])
		}

	}
	return result
}
