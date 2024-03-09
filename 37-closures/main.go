// ASSIGNMENT Keeping track of how many emails we send is mission-critical at
// Mailio. Complete the adder() enclosing function.
//
// It should return a function that adds its input (an int) to an enclosed sum
// value, then return the new sum. In other words, it keeps a running total of
// the sum variable within a closure.

package main

func adder() func(int) int {
	t := 0
	return func(x int) int {
		t += x
		return t
	}
}
