package asciiart

import "fmt"

// This function checks if the input given by the user is within the the banner file range (ascii values from 32 to 126).
func StrValidate(str string) bool {
	for _, v := range str {
		if v >= 32 && v <= 126 {
			continue
		} else {
			if v == '\n' {
				fmt.Println("'\\n' is not allowed.")
				return false
			}
			fmt.Printf("Unrecognised character: %v\n", string(v))
			return false
		}
	}
	return true
}
