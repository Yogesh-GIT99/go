// There are n cars traveling to the same destination on a one-lane highway.
// You are given two arrays of integers position and speed, both of length n.
// position[i] is the position of the ith car (in miles)
// speed[i] is the speed of the ith car (in miles per hour)
// The destination is at position target miles.

// A car can not pass another car ahead of it. It can only catch up to another car and then drive at the same speed as the car ahead of it.
// A car fleet is a non-empty set of cars driving at the same position and same speed. A single car is also considered a car fleet.
// If a car catches up to a car fleet the moment the fleet reaches the destination, then the car is considered to be part of the fleet.
// Return the number of different car fleets that will arrive at the destination.

package main

import "fmt"

func main() {
	position := []int{1, 4}
	speed := []int{3, 2}
	fmt.Println(carFleet(10, position, speed))
}

func carFleet(target int, position []int, speed []int) int {

	n := len(position)

	if n == 0 {
		return 0
	}

	type car struct {
		pos  int
		time float64
	}

	cars := make([]car, n) // intilizing array of struct

	for i := range cars {
		cars[i] = car{
			pos:  position[i],
			time: float64(target-position[i]) / float64(speed[i]),
		}
	}

	stack := make([]float64, 0, n)

	for _, car := range cars {
		stack = append(stack, car.time)

		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
