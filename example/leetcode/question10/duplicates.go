package main

import "fmt"

func main() {

	nums := []int{1, 2, 2, 3, 3, 3, 6, 6, 6}
	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {

	uniqueNum := make(map[int]int)
	j := 0

	for i, number := range nums {

		_, ok := uniqueNum[number]

		if !ok {

			uniqueNum[number] = nums[i]
			nums[j] = number
			j++
		} else {
			nums[i] = 0
		}
	}

	fmt.Println(nums)
	return j
}
