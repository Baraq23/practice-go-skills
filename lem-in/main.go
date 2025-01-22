package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error Opening file: ", err)
		return
	}
	defer file.Close()

	fileSlice := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileSlice = append(fileSlice, line)
	}
		
}
