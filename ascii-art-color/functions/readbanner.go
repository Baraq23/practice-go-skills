package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

// This function ReadBanner() maps the ascii charracters to their corresponding art.
func ReadBanner(file string) map[rune][]string {
	bannerOk := ValidateBanner(file)
	if !bannerOk {
		fmt.Printf("%v file is corrupt\n", file)
		os.Exit(1)
	}

	asciiMap := make(map[rune][]string)
	bannerFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: Could not open %v file\n", file)
		os.Exit(1)
	}
	defer bannerFile.Close()

	scanner := bufio.NewScanner(bannerFile)

	for i := 32; i <= 126; i++ {
		runeValue := i
		scanner.Scan()
		var art []string
		for i := 0; i < 8; i++ {
			scanner.Scan()
			line := scanner.Text()
			art = append(art, line)
		}
		asciiMap[rune(runeValue)] = art
	}
	return asciiMap
}


