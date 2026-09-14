// Given an integer array nums and an integer k, return the k most frequent elements within the array.
// The test cases are generated such that the answer is always unique.
// You may return the output in any order.

package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 2, 2, 3, 3, 3, 4, 4}, 3))
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)

	for _, num := range nums {
		m[num] += 1
	}
	fmt.Println(m)

	result := make([]int, k)

	for i := 0; i < k; i++ {
		maxValue := 0
		number := 0
		for num, count := range m {
			if count > maxValue {
				maxValue = count
				number = num
			}
		}
		result[i] = number
		delete(m, number)
		fmt.Println(m, result)
	}

	return result

}

// Pattern name: Count, then select — with bucket sort by frequency.
// When you need the top/bottom k by some count,
// first build a frequency map, then use the count itself as an array index instead of sorting.
// The obvious move is to sort the frequency map — that's O(n log n). The lesson is that you don't need to sort,
//  because a count can never exceed n (an element can't appear more times than the array's length).
