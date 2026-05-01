package main

import "fmt"

func main() {

	var nums = []int{2, 5, 5, 11}
	var target = 10

	fmt.Println(twoSum(nums, target))

}

//brute force

// func twoSum(nums []int, target int) []int {

// 	var result = []int{0, 0}

// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				result[0] = i
// 				result[1] = j
// 				return result
// 			}

// 		}
// 	}

// 	return result
// }

//better approach

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}

		m[num] = i
	}

	return []int{}
}
