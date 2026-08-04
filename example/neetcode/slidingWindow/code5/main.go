package main

import "fmt"

func main() {

	fmt.Println(minWindow("OUZODYXAZV", "XYZ"))

}

func minWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	} // count of what char are needed

	left := 0
	required := len(t)
	formed := 0
	minlen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ { // starting with an empty window and expanding right pointer

		ch := s[right]
		window[ch]++

		if count, ok := need[ch]; ok && window[ch] == count {
			formed++ // keep track of window which char are satisfied.
		}

		for formed == required {

			if right-left+1 < minlen { // record answer for the smallest window so far
				minlen = right - left + 1
				start = left
			}

			ch = s[left]
			window[ch]-- // shrink window until window becomes invalid

			if count, ok := need[ch]; ok && window[ch] < count { // invalid window if not satified after shrink
				formed--
			}

			left++

		}
	}

	if minlen == len(s)+1 {
		return ""
	}

	return s[start : start+minlen]
}
