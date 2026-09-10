// You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.
// You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.
// Return the minimum integer k such that you can eat all the bananas within h hours.

package main

import "fmt"

func main() {
	piles := []int{1, 4, 3, 2}
	fmt.Println(minEatingSpeed(piles, 9))
}

func minEatingSpeed(piles []int, h int) int {

	left := 1 // min eating speed is 1 pile per hr
	right := 0

	for _, p := range piles { // max eating speed is max pile size per hr
		if p > right {
			right = p
		}
	}

	answer := right // asuming max speed is the answer, successful anyway.

	for left <= right {

		k := left + (right-left)/2

		if hoursNeeded(piles, k) <= h {
			answer = k // speed successful but let try eating more slower
			right = k - 1
		} else {
			left = k + 1
		}
	}
	return answer
}

func hoursNeeded(piles []int, k int) int {

	total := 0
	for _, p := range piles {
		total += (p + k - 1) / k
	}
	return total
}
