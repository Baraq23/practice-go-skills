package funcs

import (
	"fmt"
	"strings"
)



func Reverse(revfile string) (string, error) {
	reversedText := ""
	var err error
		asciiMapS, errMapS := ReadBanner("banners/standard.txt")
		if errMapS != nil {			
			return "", errMapS
		}
		asciiMapT, errMapT := ReadBanner("banners/thinkertoy.txt")
		if errMapT != nil {
			return "", errMapT
		}
		asciiMapSh, errMapSh := ReadBanner("banners/shadow.txt")
		if errMapSh != nil {
			return "", errMapSh
		}

		maps := []map[rune][]string{}
		maps = append(maps, asciiMapS, asciiMapSh, asciiMapT)

		fString := ""
		for _, m := range maps {
			reversedText, err = AsciiReverse(revfile, m)
			if err != nil {
				return "", err
			}
			fString+=reversedText
		}
		if len(fString) == 0 {
			return "", fmt.Errorf("file provided may contain incompleted or corrupted asciiart")
		}
		fString = strings.TrimSpace(fString)
		return fString, nil
}



// func AsciiReverse() if called when reverse flag is used
func AsciiReverse(revfile string, asciiMap map[rune][]string) (string, error) {
	lines, positions, err := GetEndPositionsAndLines(revfile)
	if err != nil {
		return "", err
	}

	strSlice := GetConvertedStrings(lines, positions, asciiMap)
	finalStr := strings.Join(strSlice, "\n")
	return finalStr, nil
}
