package funcs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// Function checkBanner() uses the crypto package to check the validity of the banner
func ValidateBanner(f string) (bool, error) {
	var hashMap map[string]string = map[string]string{
		"banners/standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
		"banners/shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
		"banners/thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
	}

	fileBytes, err := os.ReadFile(f)
	if err != nil {
		return false, fmt.Errorf("error reading file '%v'", f)
	}

	hasher := sha256.New()
	hasher.Write(fileBytes)
	hashInBytes := hasher.Sum(nil)
	fHash := hex.EncodeToString(hashInBytes)

	if fHash != hashMap[f] {
		return false, fmt.Errorf("error: file '%v' is corupted", f)
	}
	return true, nil
}
