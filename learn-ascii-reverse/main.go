package main

import (
	"bufio"
	"fmt"
	"os"
	// "text/scanner"
)

func main() {
	fileCont, err := os.Open("text12.txt")
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
	fmt.Println(positions)
	fmt.Println(positionSp)
	// fmt.Println(sentence)
}
