package asciiart

import (
	"fmt"
	"strings"
)

// Func WordDistributer splits the string int a slice then sends eaach part to be printed.
func WordDistributer(input, bannerTemplate string, positions []int, ansi string, lenSub int) string {
	asciiMap := ReadBanner(bannerTemplate)
	artString := strings.Builder{}
	fmt.Printf("%q\n", input)

	art := GetArt(asciiMap, input, positions, ansi, lenSub)
	artString.WriteString(art)
	return artString.String()
}
