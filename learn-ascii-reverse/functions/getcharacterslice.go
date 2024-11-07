package funcs

// Func GetCharacterSlice() returns 3D slice containing the individual ascii art graphic characters
func GetCharacterSlice(lines []string, positions []int) [][]string {
	artCharacters := [][]string{}
	character := []string{}
	for i, pos := range positions {
		prevPos := 0
		if i > 0 {
			prevPos = positions[i-1]+1
		}
		for j := range lines {
			charPart := ""
			if len(positions)-1 == i {
				charPart = lines[j][prevPos:]
			} else {
				charPart = lines[j][prevPos : pos+1]
			}
			character = append(character, charPart)
		}
		artCharacters = append(artCharacters, character)
		character = []string{}
	}
	return artCharacters
}