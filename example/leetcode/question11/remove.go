package main

import "fmt"

func main() {

	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))

}

func removeElement(nums []int, val int) int {

	k := 0

	for _, number := range nums {
		if int(number) == val {
			nums[k] = number
			k++

		}
	}
	return k
}
