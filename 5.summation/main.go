// Write a program that finds the summation of every number from 1 to num. The number will always be a positive integer greater than 0.

// For example:

// summation(2) -> 3
// 1 + 2

// summation(8) -> 36
// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8

package main

import "fmt"

func main() {
	fmt.Println(Summation(8))
}

func Summation(n int) int {
	newSummation := 0

	for i := 0; i <= n; i++ {
		newSummation += i
	}

	return newSummation
}
