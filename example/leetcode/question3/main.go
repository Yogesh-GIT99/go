package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{3, 43, 5, 3, 2}, 7))

}

// using maps:

func twoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {

		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{i, j}
		}

		seen[num] = i

	}

	return nil
}

// brute force approach

// func twoSum(nums []int, target int) []int {

// 	for i, num1 := range nums {
// 		for j, num2 := range nums {

// 			if i != j && target-num1 == num2 {
// 				return []int{i, j}
// 			}
// 		}

// 	}

// 	return nil
// }
