package generaterandomcharacter

import (
	"fmt"
	"math/rand"
)

// Function GenerateRandomCharacters generates three random characters added to a downloaded file.
func GenerateRandomCharacters() string {
	charset := "abcdefghijklmnopqrstuvwxyz"
	var randomCharacters []byte

	if len(charset) == 0 {
		fmt.Println("String to generate random characters from cannot be empty")
		return ""
	}

	for i := 0; i < 3; i++ {
		randomCharacter := charset[rand.Intn(len(charset))]
		randomCharacters = append(randomCharacters, randomCharacter)
	}
	return string(randomCharacters)
}
