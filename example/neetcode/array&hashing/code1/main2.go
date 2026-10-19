// You’re given an integer array nums and an integer k. Return true if there are two different indices i and j such that:
// nums[i] == nums[j], and
// abs(i - j) <= k.

package main

import "fmt"

func main2() {
	nums := []int{1, 2, 3, 1, 2, 3}

	fmt.Println(containsNearByDuplicate(nums, 2))
}

func containsNearByDuplicate(nums []int, k int) bool {

	check := make(map[int]int)

	for i, num := range nums {

		if idx, ok := check[num]; ok {

			if i-idx <= k {
				return true
			}
		}

		check[num] = i
	}
	return false
}

// though sliding window is a better solution for this question on space wise O(k) instead of O(n) here.
// When the difference matters:

// If k is small and n is huge (e.g. n = 10⁵, k = 10), the window stores 10 numbers while your map may store 100,000.
// If k ≥ n, the two use the same space.
