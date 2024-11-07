package main

import (
	"fmt"
	"strings"

	"rev/functions"
)

func main() {
	lines, positions, err := funcs.GetEndPositionsAndLines("text.txt")
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
	strSlice := GetConvertedStrings(lines, positions, asciiMap)
	// fmt.Println(artStr)

	finalStr := strings.Join(strSlice, "\n")
	fmt.Println(finalStr)
}

// Func GetConvertedStrings() returns a slice of stings converted from ascii art to normal text
func GetConvertedStrings(lines [][]string, positions [][]int, asciiMap map[rune][]string) []string {
	sliceStr := []string{}

	for i, l := range lines {
		charSlice := funcs.GetCharacterSlice(l, positions[i])
		artStr := funcs.GetArt(charSlice)
		fmt.Println(artStr)
		convertedStr := funcs.ArtToString(asciiMap, charSlice)
		sliceStr = append(sliceStr, convertedStr)
	}
	return sliceStr
}
