// You are given an array of integers temperatures where temperatures[i] represents the daily temperatures on the ith day.
// Return an array result where result[i] is the number of days after the ith day before a warmer temperature appears on a
// future day. If there is no day in the future where a warmer temperature will appear for the ith day, set result[i] to 0 instead.

package main

func main() {

}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	result := make([]int, n)   // intialized all with zero
	stack := make([]int, 0, n) // stack of len zero

	for i, temp := range temperatures {

		if len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}

		stack = append(stack, i)
	}

	return result
}
