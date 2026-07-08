// You are given a string s consisting of only uppercase english characters and an integer k. You can choose up to k characters of the string and replace them with any other uppercase English character.
// After performing at most k replacements, return the length of the longest substring which contains only one distinct character.
package main

import (
	"fmt"
)

func main() {

	fmt.Println(characterReplacement("AAABABB", 1))

}

func characterReplacement(s string, k int) int {
	// total length - max repeating characters <= k

	count := make(map[byte]int)
	left := 0
	longest := 0
	maxCount := 0

	for right := 0; right < len(s); right++ {

		count[s[right]]++
		if count[s[right]] > maxCount {
			maxCount = count[s[right]]
		}

		if (right-left+1)-maxCount > k { // shrink when condition doesn't match
			count[s[left]]--
			left++
		}

		window := right - left + 1
		if window > longest {
			longest = window
		}

	}

	return longest

}
