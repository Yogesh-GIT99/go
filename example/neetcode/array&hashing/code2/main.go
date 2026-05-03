// Valid Anagram
// Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.
// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different

package main

import "fmt"

func main() {

	fmt.Println(isAnagram("racecar", "carrace"))

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range t {
		count[ch]++
	}

	for _, ch := range s {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true

}
