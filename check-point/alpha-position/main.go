package main

import "fmt"

func main() {
    fmt.Println(AlphaPosition('a'))
    fmt.Println(AlphaPosition('z'))
    fmt.Println(AlphaPosition('B'))
    fmt.Println(AlphaPosition('Z'))
    fmt.Println(AlphaPosition('0'))
    fmt.Println(AlphaPosition(' '))
}

func AlphaPosition(c rune) int {
	val := -1
	if c >= 'a' && c <= 'z' {
		val = (int(c) - (int('a')) + 1 )
	}
	if c >= 'A' && c <= 'Z' {
		val = (int(c) - (int('A')) + 1 )
	}
	return val
}