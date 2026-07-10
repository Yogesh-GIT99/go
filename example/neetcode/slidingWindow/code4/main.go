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
