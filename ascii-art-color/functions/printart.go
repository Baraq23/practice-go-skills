package asciiart

import "fmt"

// Func GetArt() takes the map of ascii art and and string then returns the art in string form.
func GetArt(asciiMap map[rune][]string, input string, positions []int, ansi string, lenSub int) string {
	store := [][]string{}
	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}

	storeColor := store
	if ansi != "" {
		storeColor = colorArt(store, positions, ansi, lenSub)
	}
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
	
	pos := 0

	fmt.Println("positions:", positions)
	fmt.Println("lenof subS:", lenSub)
	for i := range store {
		charColor := []string{}
		if len(positions) != 0 {
			if i >= positions[pos] && i < positions[pos]+lenSub {				
				charColor = colorWholeChar(ansi, store[i])
			} else {
				charColor = store[i]
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
