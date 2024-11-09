package funcs

import (
	"strings"
)

// Func ArtToString() recieves a 2D slice containing art strings and converts it to a normal string(returns the string)
func ArtToString(asciiMap map[rune][]string, charSlice [][]string) string {
	wordStr := strings.Builder{}

	for _, char := range charSlice {
		for k, v := range asciiMap {
			if DeepEqual(v, char) {
				wordStr.WriteString(string(k))

			}
		}
	}
	return wordStr.String()
}

// Func DeepEqual() checks if the two slices containst the same content
func DeepEqual(v, c []string) bool {
	for i := 0; i < 8; i++ {
		if v[i] != c[i] {
			return false
		}
	}
	return true
}
