// Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].

// Each product is guaranteed to fit in a 32-bit integer.

// Follow-up: Could you solve it in O(n)
// O(n) time without using the division operation?

package main

import "fmt"

func main() {
	var nums = []int{1, 2, 4, 6}
	fmt.Println(productExceptSelf(nums))

}

func productExceptSelf(nums []int) []int {

	n := len(nums)
	output := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		output[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= suffix
		suffix *= nums[i]
	}

	return output
}
