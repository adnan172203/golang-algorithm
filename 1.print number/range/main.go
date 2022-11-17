package main

import "fmt"

func main() {
	printNumberWithRange()
}
func printNumberWithRange() {

	for i, _ := range [10]int{} {
		fmt.Println(i)
	}
}
