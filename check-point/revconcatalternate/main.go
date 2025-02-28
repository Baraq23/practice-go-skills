package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	s1 := slice1
	s2 := slice2
	fs := []int{}

	if len(s1) == len(s2) {
		for i := len(s1) - 1; i >= 0; i-- {
			fs = append(fs, s1[i])
			fs = append(fs, s2[i])
		}
	}

	if len(s1) > len(s2) {
		for i := len(s1) - 1; i >= 0; i-- {
			if i > len(s2)-1 {
				fs = append(fs, s1[i])
			} else {
				fs = append(fs, s1[i])
				fs = append(fs, s2[i])
			}
		}

	}

	if len(s1) < len(s2) {
		for i := len(s2) - 1; i >= 0; i-- {
			if i > len(s1)-1 {
				fs = append(fs, s2[i])
			} else {
				fs = append(fs, s1[i])
				fs = append(fs, s2[i])
			}
		}
	}

	return fs
}
