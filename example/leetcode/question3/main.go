package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{3, 43, 5, 3, 2}, 7))

}

func twoSum(nums []int, target int) []int {

	for i, num1 := range nums {
		for j, num2 := range nums {

			if i != j && target-num1 == num2 {
				return []int{i, j}
			}
		}

	}

	return nil
}
