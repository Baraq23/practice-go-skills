package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"rev/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}

	file := flag.String("reverse", "", "Name of 'txt' file containing the ascii art.")
	flag.Parse()

	if os.Args[1][0:2] != "--" {
		fmt.Println("error: incorrect flag format.")
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}

	if *file == "" {
		fmt.Println("error: file name name not give.")
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}


	lines, positions, err := funcs.GetEndPositionsAndLines(*file)
	if err != nil {
		fmt.Println(err)
		return
	}

	bannerTemplate := "banners/standard.txt"

	asciiMap, err := funcs.ReadBanner(bannerTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}
	strSlice := funcs.GetConvertedStrings(lines, positions, asciiMap)
	finalStr := strings.Join(strSlice, "\n")
	fmt.Println(finalStr)
}
