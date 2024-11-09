package asciiart

// This function stores the ascii art to then prints the ascii art.
func RearrangeArt(asciiMap map[rune][]string, input string) string {
	// 2d slice to store the Art.
	store := [][]string{}
	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}
	word := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(store); j++ {
			word += store[j][i]
		}
		if i != 7 {
			word += "\n"
		}
	}
	return word
}
