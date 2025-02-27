package main

import (
	"fmt"

)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}



func PrintMemory(arr [10]byte) {
    count := 0
	for i:= range arr {
		if arr[i] == 0 {
			fmt.Print("00")
			count++
		} else {
			fmt.Print(GetHex(rune(arr[i])))
			count++
		}
		if count == 4 {
			fmt.Println()
			count = 0
		} else {
			if i < len(arr) - 1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
}


func GetHex(r rune) string {
	hex := ""
	h := "0123456789abcdef"

	for r != 0 {
		mod := r%16
		hex=string(h[mod])+hex
		r = r/16
	}
	return hex
}

// Output

// $ go run . | cat -e
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
// $
