package asciiart

import "strings"

// Func getSubstringPositions() returns a slice of intergers containing indexes where substring starts
func GetSubStringPositions(s, sb string) []int {
	positions := []int{}
	start := 0

	for start < len(s) {
		pos := strings.Index(s[start:], sb)
		if pos == -1 {
			break
		}
		positions = append(positions, pos+start)
		start += pos + len(sb)
	}
	return positions
}
