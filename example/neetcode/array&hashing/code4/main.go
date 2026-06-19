// Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.
// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

package main

import "fmt"

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(strs))

}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int

		for _, ch := range s {
			count[ch-'a']++
		}

		m[count] = append(m[count], s)
	}

	fmt.Println(m)

	var result [][]string
	for _, group := range m {
		result = append(result, group)
	}

	return result
}

// undestand code
// func groupAnagrams(strs []string) [][]string {
// 	m := make(map[[26]int][]string)

// 	fmt.Println("Starting grouping...\n")

// 	for _, s := range strs {
// 		var count [26]int

// 		fmt.Printf("Processing word: %s\n", s)

// 		// Count characters
// 		for _, ch := range s {
// 			index := ch - 'a'
// 			count[index]++
// 			fmt.Printf("  -> char: %c, index: %d, count: %v\n", ch, index, count)
// 		}

// 		fmt.Printf("Final count key for '%s': %v\n", s, count)

// 		// Add to map
// 		m[count] = append(m[count], s)

// 		fmt.Printf("Map state after inserting '%s':\n", s)
// 		for k, v := range m {
// 			fmt.Printf("  Key: %v -> Value: %v\n", k, v)
// 		}

// 		fmt.Println("--------------------------------------------------")
// 	}

// 	// Build result
// 	fmt.Println("\nBuilding final result...\n")

// 	var result [][]string
// 	for _, group := range m {
// 		fmt.Printf("Adding group: %v\n", group)
// 		result = append(result, group)
// 	}

// 	fmt.Println("\nFinal grouped anagrams:", result)

// 	return result
// }
