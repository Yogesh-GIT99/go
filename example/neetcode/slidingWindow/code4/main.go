// You are given two strings s1 and s2.
// Return true if s2 contains a permutation of s1, or false otherwise. That means if a permutation of s1 exists as a substring of s2, then return true.
// Both strings only contain lowercase letters.

package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {

	n1, n2 := len(s1), len(s2)

	if n1 > n2 {
		return false
	}

	var need, window [26]int

	for i := 0; i < n1; i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}

	fmt.Println(need)
	fmt.Println(window)

	if need == window {
		return true
	}

	for i := n1; i < n2; i++ {

		window[s2[i]-'a']++
		window[s2[i-n1]-'a']--

		if need == window {
			return true
		}
	}

	return false
}
