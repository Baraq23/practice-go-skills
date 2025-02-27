package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(0)
		return
	}
	args = args[1:]
	count := 0

	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			count++
		}
	}
	// fmt.Println(args)
	fmt.Println(count)
}
