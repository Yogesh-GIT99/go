// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] where nums[i] + nums[j] + nums[k] == 0, and the indices i, j and k are all distinct.
// The output should not contain any duplicate triplets. You may return the output and the triplets in any order.
// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]

package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] { // removing duplicates for i position
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				for j < k && nums[j] == nums[j+1] { //  removing duplicates for j position
					j++
				}

				for j < k && nums[k] == nums[k-1] { //  removing duplicates for k position
					k--
				}

				j++
				k--
			}
		}
	}

	return result
}
