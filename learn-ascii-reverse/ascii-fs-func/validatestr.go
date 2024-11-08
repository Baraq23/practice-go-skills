package asciiart

import "fmt"

// This function checks if the input given by the user is within the the banner file range (ascii values from 32 to 126).
func StrValidate(str string) error {
	for _, v := range str {
		if v >= 32 && v <= 126 || v == '\n' {
			continue
		} else {
			return fmt.Errorf("Unrecognised character: %v\n", v)
		}
	}
	return nil
}
