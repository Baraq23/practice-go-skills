package output

import "fmt"

// Function StrValidate() checks if the input given by the user is within the the banner file range (ascii values from 32 to 126).
func StrValidate(str string) bool {
	for _, v := range str {
		if v >= 32 && v <= 126 || v == '\n' {
			continue
		} else {
			fmt.Printf("Unrecognised character: %v\n", v)
			return false
		}
	}
	return true
}
