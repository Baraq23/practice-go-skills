package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// GetEndPositions() returns slices of indexes from which each graphic end
func GetEndPositionsAndLines(file string) ([][]string, [][]int, error) {
	fileCont, err := os.Open(file)
	if err != nil {
		// fmt.Println("unable to open file")
		return nil, nil, fmt.Errorf("error: unable to open the file '%v'", file)
	}
	defer fileCont.Close()

	scanner := bufio.NewScanner(fileCont)

	allLines := [][]string{}
	lines := []string{}
	count := 0
	totallines := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		count++
		totallines++
		// fmt.Println("line", totallines)
		if count == 8 {
			allLines = append(allLines, lines)
			count = 0
			lines = []string{}
		}
	}

	fmt.Println("totallines", totallines)

	if totallines%8 != 0 {
		return nil, nil, fmt.Errorf("files with new lines are not allowed!")
	}
	fmt.Println("All lines length:", len(allLines))

	allPositions := [][]int{}
	for _, l := range allLines {
		positions := GetPositions(l)
		allPositions = append(allPositions, positions)
	}
	fmt.Println(allPositions)
	return allLines, allPositions, nil
}

func GetPositions(lines []string) []int {
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
	fmt.Println(positions)
	return positions
}
