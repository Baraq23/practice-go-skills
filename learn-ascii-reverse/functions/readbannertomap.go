package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// This function ReadBanner() maps the ascii charracters to their corresponding art.
func ReadBanner(file string) (map[rune][]string, error) {
	validBannerTemp := ValidateBanner(file)
	if !validBannerTemp {
		return nil, fmt.Errorf("error: banner file '%v' not valid", file)		
	}
	asciiMap := make(map[rune][]string)
	bannerFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error: could not open banner file '%v'", file)
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
	return asciiMap, nil
}
