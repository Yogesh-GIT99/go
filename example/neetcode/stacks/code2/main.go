// Design a stack class that supports the push, pop, top, and getMin operations.

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// Each function should run in
// O(1) time.

package main

import "fmt"

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) push(val int) {
	this.stack = append(this.stack, val)

	if len(this.min) == 0 || val < this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) getMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	minStack := Constructor()
	minStack.push(1)
	minStack.push(2)
	minStack.push(0)
	fmt.Println(minStack.getMin()) // 0
	minStack.pop()
	fmt.Println(minStack.top())    // 2
	fmt.Println(minStack.getMin()) // 1
}
