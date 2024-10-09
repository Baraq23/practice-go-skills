package main

import "fmt"

func main() {
	sum := sumNonNeg(5)
	fmt.Println(sum)
}


/* Write a recursive function that given an input
`n` sums all non-negative integers up to `n`. */
func sumNonNeg(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumNonNeg(n-1)
}