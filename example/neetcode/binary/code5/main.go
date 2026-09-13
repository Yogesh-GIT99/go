// Design a time-based key-value data structure that can store multiple values for the same key at different time stamps
// and retrieve the key's value at a certain timestamp.

// Implement the TimeMap class:

// TimeMap() Initializes the object of the data structure.
// void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
// String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp.
// If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".

package main

import "fmt"

func main() {
	tm := Constructor()
	tm.Set("alice", "happy", 1)
	fmt.Println(tm.Get("alice", 1)) // happy
	fmt.Println(tm.Get("alice", 2)) // happy
	tm.Set("alice", "sad", 3)
	fmt.Println(tm.Get("alice", 3)) // sad
	fmt.Println(tm.Get("bob", 1))   // "" (empty line)
}

type entry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	store map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]entry)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], entry{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {

	entries := this.store[key]
	answer := ""

	left, right := 0, len(entries)-1

	for left <= right {
		mid := left + (right-left)/2

		if entries[mid].timestamp <= timestamp {
			answer = entries[mid].value // valid but later one might also work
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return answer
}
