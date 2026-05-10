// Container With Most Water:
// You are given an integer array heights where heights[i] represents the height of the ith bar.
// You may choose any two bars to form a container. Return the maximum amount of water a container can store.
package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 7, 2, 5, 4, 7, 3, 6}))
}

func maxArea(heights []int) int {

	maxAmount := 0
	i := 0
	j := len(heights) - 1

	for i < j {
		width := j - i
		area := 1

		if heights[i] > heights[j] {
			area = heights[j] * width
			j--
		} else {
			area = heights[i] * width
			i++
		}

		if area > maxAmount {
			maxAmount = area
		}
	}
	return maxAmount
}
