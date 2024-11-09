package asciiart

import (
	"rev/functions"
	"strings"
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
	artSlice := []string{}

	words := strings.Split(input, "\n")
	// newLine := false
	art := ""
	nl := ""
	for _, word := range words {
		if word == "" {
			nl+="\n"
			continue
		} else {
			// fmt.Println(len(nl))
			art = nl + RearrangeArt(asciiMap, word)
			nl = ""
		}
		artSlice = append(artSlice, art)
	}
	artString := strings.Join(artSlice, "\n")
	artString+=nl
	return artString, nil
}
