// You are given an array of integers heights where heights[i] represents the height of a bar. The width of each bar is 1.
// Return the area of the largest rectangle that can be formed among the bars.

package main

import "fmt"

func main() {
	height := []int{7, 1, 7, 2, 2, 4}
	fmt.Println(largestRectangleArea(height))
}

func largestRectangleArea(heights []int) int {

	n := len(heights) // no walls, maxarea = 0

	if n == 0 {
		return 0
	}

	left := make([]int, n)
	right := make([]int, n)

	stack := []int{}

	// shortest wall on the left
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	stack = []int{}

	// shortest wall on the right
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// get the max area
	maxarea := 0
	for i := 0; i < n; i++ {
		if area := heights[i] * (right[i] - left[i] - 1); area > maxarea {
			maxarea = area
		}
	}

	return maxarea
}
