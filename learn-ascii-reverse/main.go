package main

import (
	"fmt"
	"rev/functions"
)

func main() {
	// GetEndPositions("text.txt")
	// GetEndPositions("text12.txt")
	lines, positions := funcs.GetEndPositions("tt.txt")
	if len(positions) == 0 {
		fmt.Println("error: not able to read file or file is empty")
		return
	}
	charSlice := funcs.GetCharacterSlice(lines, positions)

	// fmt.Println(positions)
	// fmt.Println(len(lines))
	// fmt.Println(len(charSlice))

	artStr := funcs.GetArt(charSlice)
	fmt.Println(artStr)

	bannerTemplate := "banners/standard.txt"
	asciiMap := funcs.ReadBanner(bannerTemplate)

	convertedStr := funcs.ArtToString(asciiMap, charSlice)

	fmt.Println(convertedStr)
}
