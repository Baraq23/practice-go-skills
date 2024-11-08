package asciiart

import (
	"fmt"
	"strings"
	"rev/functions"
)

// This function is the starting point of the program.
// Banner file and input string has already been validated in main.go.
func GetAsciiArt(input, bannerTemplate string) (string, error) {
	if err := StrValidate(input); err != nil {
		return "", err
	}
	input = FormatSpChar(input)
	asciiMap, err := funcs.ReadBanner(bannerTemplate)
	if err != nil {
		return "", err
	}
	//Art string will be used during testing of the program.
	artString := ""
	words := strings.Split(input, "\n")
	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			art := PrintArt(asciiMap, word)
			artString = art
		}
	}
	// This return value is used for testing purposes.
	return artString, nil
}
