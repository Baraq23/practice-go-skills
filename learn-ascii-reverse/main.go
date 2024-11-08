package main

import (
	"flag"
	"fmt"
	"os"
	// "strings"

	"rev/functions"
)

func main() {

	revfile := flag.String("reverse", "", "Name of 'txt' file containing the ascii art.")
	// outPutFile := flag.String("output", "", "Name of 'txt' file to be created.")
	// bannerTempl := flag.String("banner", "standard", "Name of banner template to be used.")
	flag.Parse()

	// check if the length of the argument after the flags is only 1.
	if flag.NArg() > 1 {
		fmt.Println("error: too many arguments.")
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --output=<fileName> \"some text.\"")
		return
	}

	// Handling banner templates
	// if *bannerTempl != "shadow" && *bannerTempl != "thinkertoy" && *bannerTempl != "standard" {
	// 	fmt.Println("error: undefined banner template")
	// 	fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --output=<fileName> --banner=<bannerTemplate")
	// 	return
	// }
	// bannerTemplate := "banners/" + *bannerTempl + ".txt"

	// asciiMap, errMap := funcs.ReadBanner(bannerTemplate)
	// if errMap != nil {
	// 	fmt.Println(errMap)
	// 	return
	// }

	if os.Args[1][0:2] != "--" {
		fmt.Println("error: incorrect flag format.")
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}

	// Handling flags
	// if *outPutFile == "" {
	// 	fmt.Println("error: output file name name not give.")
	// 	fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName> --output=<fileName> \"some text.\"")
	// 	return
	// }

	
	if *revfile != "" {
		reversedText,err := Reverse(*revfile)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
			return
		}
		fmt.Println(reversedText)		
	}
}

func Reverse(revfile string) (string, error) {

	reversedText := ""
	var err error
		asciiMapS, errMapS := funcs.ReadBanner("banners/standard.txt")
		if errMapS != nil {			
			return "", errMapS
		}
		asciiMapT, errMapT := funcs.ReadBanner("banners/thinkertoy.txt")
		if errMapT != nil {
			return "", errMapT
		}
		asciiMapSh, errMapSh := funcs.ReadBanner("banners/shadow.txt")
		if errMapSh != nil {
			return "", errMapSh
		}

		maps := []map[rune][]string{}
		maps = append(maps, asciiMapS, asciiMapSh, asciiMapT)

		for _, m := range maps {
			reversedText, err = funcs.AsciiReverse(revfile, m)
			if err != nil {
				return "", err
			}
			if len(reversedText) != 0 {
				break
			}
		}		
		return reversedText, nil
}
