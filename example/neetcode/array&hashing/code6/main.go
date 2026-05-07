package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := Solution{}
	encode := s.Encode([]string{"Hello", "Hi"})
	fmt.Println(encode)

	fmt.Println(s.Decode(encode))

}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""

	for _, str := range strs {
		result += strconv.Itoa(len(str))
		result += "#"
		result += str
	}

	return result

}

func (s *Solution) Decode(encoded string) []string {

	var result []string
	i := 0

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		j++

		word := encoded[j : j+length]
		result = append(result, word)

		i = j + length
	}
	return result

}
