// Create a method each_cons that accepts a list and a number n, and returns cascading subsets of the list of size n, like so:

// each_cons([1,2,3,4], 2)
//   #=> [[1,2], [2,3], [3,4]]

// each_cons([1,2,3,4], 3)
//   #=> [[1,2,3],[2,3,4]]
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}

	fmt.Println(EachCons(arr, 2))
}

func EachCons(arr []int, num int) [][]int {
	var newArray [][]int

	for i := 0; i < len(arr); i++ {
		var smallArray []int

		for j := 0; j < num; j++ {
			index := i + j
			if index >= len(arr) {
				return newArray
			}
			smallArray = append(smallArray, arr[index])
		}
		newArray = append(newArray, smallArray)
	}

	return newArray
}
