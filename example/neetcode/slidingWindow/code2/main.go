package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("zxyzxyz"))
}

func lengthOfLongestSubstring(s string) int {

	Occur := make(map[byte]int)
	left := 0
	longest := 0

	if s == " " || len(s) == 1 {
		return 1
	}

	for right := 0; right < len(s); right++ {

		char := s[right]

		if idx, check := Occur[char]; check && idx >= left {
			left = idx + 1
		}

		Occur[char] = right
		windowSize := right - left + 1

		if windowSize > longest {
			longest = windowSize
		}
	}

	return longest
}
