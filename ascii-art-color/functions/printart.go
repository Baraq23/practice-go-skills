package asciiart


// Func GetArt() takes the map of ascii art and and string then returns the art in string form.
func GetArt(asciiMap map[rune][]string, input string, positions []int, ansi string, ansi2 string, lenSub int) string {
	store := [][]string{}
	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}

	storeColor := store
	if ansi != "" {
		storeColor = colorArt(store, positions, ansi, ansi2, lenSub)
	}

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
func colorArt(store [][]string, positions []int, ansi string, ansi2 string, lenSub int) [][]string {
	storeColored := [][]string{}
	
	pos := 0
	for i := range store {
		var charColor []string
		if len(positions) != 0 {
			if i >= positions[pos] && i < positions[pos]+lenSub {				
				charColor = colorWholeChar(ansi, store[i])
			} else {
				if ansi2 != "" {
					charColor = colorWholeChar(ansi2, store[i])
				} else {
					charColor = store[i]
				}
			}
			// Move to the next index of the positions
			if positions[pos]+lenSub-1 == i {
				if pos < len(positions)-1{
					pos++
				}
			}
		} else {
			charColor = colorWholeChar(ansi, store[i])
		}
		storeColored = append(storeColored, charColor)
	}
	return storeColored
}

// Func colorWholeCharacter() returns a slice of strings of olored character
func colorWholeChar(ansi string, charSlice []string) []string {
	newSlice := []string{}
	res := "\x1b[0m"
	for _, c := range charSlice {
		st := ansi + string(c) + res
		newSlice = append(newSlice, st)
	}
	return newSlice
}
