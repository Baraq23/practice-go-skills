package funcs

import "fmt"

// Func GetConvertedStrings() returns a slice of stings converted from ascii art to normal text
func GetConvertedStrings(lines [][]string, positions [][]int, asciiMap map[rune][]string) []string {
	sliceStr := []string{}

	for i, l := range lines {
		charSlice := GetCharacterSlice(l, positions[i])
		artStr := GetArt(charSlice)
		fmt.Println(artStr)
		convertedStr := ArtToString(asciiMap, charSlice)
		sliceStr = append(sliceStr, convertedStr)
	}
	return sliceStr
}
