package main

import "fmt"

func main() {

	retunadd := add(1, 2)

	fmt.Println(retunadd)

}

// this is a limitation of interfaces:  have to write this big code , need more generic code.

// func add(a, b any) any { // any and interfaces{} are the same thing

// 	aInt, aIsint := a.(int)
// 	bInt, bIsint := b.(int)

// 	if aIsint && bIsint {
// 		return aInt + bInt
// 	}

// 	aFloat, aIsFloat := a.(float64)
// 	bFloat, bIsFloat := b.(float64)

// 	if aIsFloat && bIsFloat {
// 		return aFloat + bFloat
// 	}

// 	aString, aIsString := a.(string)
// 	bString, bIsString := b.(string)

// 	if aIsString && bIsString {
// 		return aString + bString
// 	}

// 	return nil
// }

// generic code

func add[T int | float64 | string](a, b T) T {

	return a + b
}

// using generics we can simplyfy the code to this much level and overcome limitations of interfaces.
