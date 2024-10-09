package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("missing row(s)")
		return
	}
	sdkRows := os.Args[1:]
	for i :=range sdkRows {
		if len(sdkRows[i]) != 9 {
			fmt.Println("Incorrect number of columns at row number:", i+1)
			return 
		}
	}

	var unsolved [][]string
	for _, s := range sdkRows {
		unsolved = append(unsolved, strings.Split(s, ""))
	}

	printSdk(unsolved)

}

func printSdk(sdk [][] string) {
	for _, v := range sdk {
		fmt.Println(v)
	}
}