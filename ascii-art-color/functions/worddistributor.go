package asciiart

import "strings"

// Func WordDistributer splits the string int a slice then sends eaach part to be printed.
func WordDistributer(input, bannerTemplate string, positions []int, ansi string, ansi2 string, lenSub int) string {
	asciiMap := ReadBanner(bannerTemplate)
	artString := strings.Builder{}

	art := GetArt(asciiMap, input, positions, ansi, ansi2, lenSub)
	artString.WriteString(art)
	return artString.String()
}
