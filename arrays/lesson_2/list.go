package main

import "fmt"

func main() {
	platforms := [4]string{"patna", "jhodpur", "srinagar"}
	fmt.Println(platforms[1:2]) // slicing 1st and last
	// other ways
	fmt.Println(platforms[:2])
	fmt.Println(platforms[1:]) // we cannot use -1 here unlike other programs and cannot use index value +1 more than the last index

	// build another slice from slice
	station := platforms[1:]
	stations := station[:1]

	fmt.Println(stations)
	fmt.Println(len(stations), cap(stations)) // 1 3 output represent len count the content and cap counts the remaining space to right. ( it doesnot account what was sliced from left, but what is left at right side)

	stations = stations[:3]
	fmt.Println(stations)
	fmt.Println(len(stations), cap(stations)) // 3 3 justifies what we have disscused in the previous line. there is capacity left on the right side, which can be extended since array was [4]
	// slice is part of array/ref to a array ( similar what happens in pointer ), golang does not create another mem location for slices, it ref to the same array.

}
