package funcs

import "strings"

// func AsciiReverse if called when reverse flag is used
func AsciiReverse(revfile string, asciiMap map[rune][]string) (string, error) {
	lines, positions, err := GetEndPositionsAndLines(revfile)
	if err != nil {
		// fmt.Println(err)
		return "", err
	}

	strSlice := GetConvertedStrings(lines, positions, asciiMap)
	finalStr := strings.Join(strSlice, "\n")
	return finalStr, nil
}
