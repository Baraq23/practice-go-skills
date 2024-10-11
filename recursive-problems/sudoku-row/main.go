package main

import (
	"fmt"
	"os"
	"strings"
	// "time"
)

func main() {
	

	if len(os.Args) != 2 {
		fmt.Println("missing row(s)")
		return
	}

	inputR := os.Args[1]
	if len(inputR) != 9 {
		fmt.Println("Incorrect number of value provided in the row")
		fmt.Println("Values length must equal 9")
		return
	}

	unsolved := strings.Split(inputR, "")
	if solve(unsolved, 0) {
		fmt.Println("row solved")
	} else {
		fmt.Println("No solution found")
	}

	// fmt.Println(unsolved)	

}

func solve(sdk []string, indx int) bool {
	// fmt.Println("IndeR:", indx)
	if indx == 9 {
		return true
	}
	if sdk[indx] == "." {
		for i := '1'; i <= '9'; i++ {
			sdk[indx] = string(i)
			// fmt.Println("Index:", indx)
			fmt.Println(sdk)
			// time.Sleep(2*time.Second)
			if isValid(sdk, indx, string(i)) {
				if solve(sdk, indx+1) {
					return true
				}
			}
			sdk[indx] = "."
		}
	} else {
		if solve(sdk, indx+1) {
			return true
		}
	}
	return false
}

func isValid(sdk []string, indx int, tv string) bool {
	for i := range sdk {
		if sdk[i] == tv && indx != i {
			return false
		}
	}
	return true
}


