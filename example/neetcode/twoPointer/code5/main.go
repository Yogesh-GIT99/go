package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}))
}

func trap(height []int) int {

	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	area := 0

	for left < right {

		if height[left] < height[right] {

			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				area += leftMax - height[left]
			}

			left++

		} else {

			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				area += rightMax - height[right]
			}

			right--
		}
	}

	return area
}
