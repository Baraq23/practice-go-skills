package main

import (
	"fmt"
	"os"
	"strings"
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

	fmt.Println(unsolved)
	

}


