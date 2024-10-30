package asciiart

// Func GetArt() takes the map of ascii art and and string then returns the art in string form.
func GetArt(asciiMap map[rune][]string, input string, positions []int, ansi string, lenSub int) string {
	store := [][]string{}
	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}

	storeColor := colorArt(store, positions, ansi, lenSub)
	// fmt.Println("store color",storeColor)

	// Hold the word in a variable.
	artSentence := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(storeColor); j++ {
			artSentence += storeColor[j][i]
		}
		if i < 7 {
			artSentence += "\n"
		}
	}
	return artSentence
}

// func colorArt() assigns colors to the art at the desired indexes given by the positions
func colorArt(store [][]string, positions []int, ansi string, lenSub int) [][]string {
	storeColored := [][]string{}
	res := "\x1b[0m"
	pos := 0

	for i := range store {
		charColor := []string{}
		for _, c := range store[i] {
			if len(positions) != 0 {
				if i >= positions[pos] && i < positions[pos]+lenSub {
					charColor = append(charColor, ansi+string(c)+res)
				} else {
					charColor = append(charColor, string(c))
				}
				// Move to the next index of the positions
				if positions[pos] == i {
					if pos < len(positions)-1 {
						pos++
					}
				}
			} else {
				charColor = append(charColor, ansi+string(c)+res)
			}
		}
		storeColored = append(storeColored, charColor)
	}
	return storeColored
}
