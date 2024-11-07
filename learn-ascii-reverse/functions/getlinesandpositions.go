package funcs

import (
	"bufio"
	"os"
)

// GetEndPositions() returns slices of indexes from which each graphic end
func GetEndPositions(file string) ([]string, []int) {
	fileCont, err := os.Open(file)
	if err != nil {
		// fmt.Println("unable to open file")
		return []string{}, []int{}
	}
	defer fileCont.Close()

	scanner := bufio.NewScanner(fileCont)

	lines := []string{}
	for i := 0; i < 8; i++ {
		scanner.Scan()
		line := scanner.Text()
		lines = append(lines, line)
	}
	space := false
	count := 0
	positions := []int{}
	for j := 0; j < len(lines[0]); j++ {
		for i := 0; i < 8; i++ {
			if lines[i][j] != ' ' {
				space = false
				count = 0
				break
			}

			if i == 7 {
				space = true
				if count == 0 {
					positions = append(positions, j)
				}
				if space {
					count++
				}
			}
			if count == 6 {
				count = 0
			}
		}
	}
	return lines, positions
}
