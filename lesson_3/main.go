// Type converter

package main

import "fmt"

func main() {
	var distance int
	time := 2
	speed := 10.5

	distance = int(speed * float64(time))

	fmt.Println(distance)
}
