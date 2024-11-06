package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	asciiart "asciiart/functions"
)

func main() {
	errorMessege := "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println(errorMessege)
		return
	}

	color := flag.String("color", "", "rgb color code")
	flag.Parse()

	if *color == "" {
		fmt.Println("Color not given.")
		fmt.Println(errorMessege)
		return
	}

	str, subS, tempFlag, ansi, ansi2 := "", "", "", "", ""
	var err error
	if flag.NFlag() == 0 {
		if len(os.Args) < 2 || len(os.Args) > 3 {
			fmt.Println(errorMessege)
			return
		}
		if len(os.Args) == 2 {
			str = os.Args[1]
		} else {
			str = os.Args[1]
			tempFlag = os.Args[2]
		}
	} else {
		if os.Args[1][:8] != "--color=" {
			fmt.Println(errorMessege)
			return
		}
		str, subS, tempFlag = asciiart.GetInputs(os.Args)
		if str == "" {
			fmt.Println(errorMessege)
			return
		}
		ansi, ansi2, err = asciiart.GetAnsi(*color)
		if err != nil {
			fmt.Println(err)
			fmt.Println(errorMessege)
			return
		}
	}

	validateStr := asciiart.StrValidate(str)
	if !validateStr {
		return
	}

	positions := []int{}
	lenSub := 0
	if subS != "" {
		validateSubStr := asciiart.StrValidate(subS)
		if !validateSubStr {
			return
		}
		if !strings.Contains(str, subS) {
			fmt.Println(fmt.Errorf("Error: Substring is not contained in the main string"))
			return
		}
		lenSub = len(subS)
		positions = asciiart.GetSubStringPositions(str, subS)
	}

	bannerTemplate := filepath.Join("banners", "standard.txt")

	if tempFlag != "" {

		if tempFlag != "-sh" && tempFlag != "-th" && tempFlag != "-st" {
			fmt.Printf("Error: \"%v\" is not a flag_name. Usage: go run . \"<string>\" <flag_name> (flags: -sh, -st, -th)\n", tempFlag)
			return
		}

		// Handling Flags
		switch tempFlag {
		case "-sh":
			bannerTemplate = filepath.Join("banners", "shadow.txt")
		case "-th":
			bannerTemplate = filepath.Join("banners", "thinkertoy.txt")
		default:
			bannerTemplate = filepath.Join("banners", "standard.txt")
		}
	}

	artStr := asciiart.WordDistributer(str, bannerTemplate, positions, ansi,ansi2, lenSub)

	fmt.Println(artStr)
}
