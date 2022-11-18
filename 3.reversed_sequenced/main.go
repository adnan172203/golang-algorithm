package main

func main() {
	ReverseSeq(5)
}

func ReverseSeq(n int) []int {
	reverseArray := []int{}

	for i := n; i > 0; i-- {
		reverseArray = append(reverseArray, i)
	}

	return reverseArray
}
