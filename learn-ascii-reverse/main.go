package main

import (
	"bufio"
	"fmt"
	"os"
	// "text/scanner"
)

func main() {
	printSl("text.txt")
	printSl("text12.txt")
	printSl("text123C.txt")
}

// printSl() prints slices of indexes from which each graphic end
func printSl(file string) {
	fileCont, err := os.Open(file)
	if err != nil {
		fmt.Println("unable to open file")
		return
	}
	defer fileCont.Close()

	scanner := bufio.NewScanner(fileCont)

	sentence := []string{}

	for i := 0; i < 8; i++ {
		scanner.Scan()
		line := scanner.Text()
		sentence = append(sentence, line)
	}

	space := false
	count := 0
	positions := []int{}
	positionSp := []int{}
	for j := 0; j < len(sentence[0]); j++ {
		for i := 0; i < 8; i++ {
			if sentence[i][j] != ' ' {
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
				positionSp = append(positionSp, j+1)
				count = 0
			}
		}
	}
	fmt.Println("Indexes: ", positions)
	fmt.Println("SpaceBar:", positionSp)
	fmt.Println()
}
