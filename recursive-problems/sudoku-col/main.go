package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Incorrect number of collumns")
	}

	cols := os.Args[1:]
	
	
	if solve(cols, 0) {
		fmt.Println("solution found")
	} else {
		fmt.Println("No solutions found")
	}
}


func solve(col []string, indx int) bool {
	if indx == 9 {
		return true
	}

	if col[indx] == "." {
		for i := '1'; i <= '9'; i++ {
			col[indx] = string(i)
			fmt.Println(col)
			time.Sleep(2 * time.Second)
			if isValid(col, indx, string(i)) {
				if solve(col, indx+1) {
					return true
				}
			} else {
				col[indx] = "."
			}
		}
	} else {
		if solve(col, indx +1) {
			return true
		}
	}
	return false
}

func isValid(col []string, indx int, tv string) bool {
	for i := range col {
		if col[i] == tv && indx != i {
			return false
		}
	}
	return true
}


