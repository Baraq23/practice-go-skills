package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [Name-of-file.txt]")
		fmt.Println("Example: go run . <fileName.txt>")
		return
	}
	inputFile := os.Args[1]
	if ext := filepath.Ext(inputFile); ext != ".txt" {
		fmt.Printf("only files with .txt extension are accepted")
		return
	}

	file, err := os.Open(inputFile)
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

	numberofants, err := strconv.Atoi(fileSlice[0])
	if err != nil {
		fmt.Println("ERROR: invalid number of Ants")
		return
	}
	fmt.Printf("Number of ants: %t\n", numberofants)

	connections, points, startRoom, endRoom := getslices(fileSlice)
	fmt.Println(connections, points, startRoom, endRoom)
	allPaths, err := DFS(connections, points, startRoom, endRoom)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print all paths
	fmt.Println("All possible paths from O to E:")
	for _, path := range allPaths {
		fmt.Println(path)
	}
}

func getslices(l []string) ([][]string, [][]string, []string, []string) {
	connections := [][]string{}
	points := [][]string{}
	startRoom := []string{}
	endRoom := []string{}

	for i, line := range l {
		if strings.Contains(line, "-") {
			c := strings.Split(line, "-")
			connections = append(connections, c)
		}
		if strings.Contains(line, " ") {
			p := strings.Split(line, " ")
			points = append(points, p)
		}
		if strings.Contains(line, "##start") && i+1 != len(l) {
			startRoom = strings.Split(l[i+1], " ")
		}
		if strings.Contains(line, "##end") && i+1 != len(l) {
			endRoom = strings.Split(l[i+1], " ")
		}
	}
	return connections, points, startRoom, endRoom
}
