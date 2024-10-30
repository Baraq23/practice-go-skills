package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	asciiart "asciiart/functions"
)

type Rgb struct {
	R int
	G int
	B int
}

func main() {
	errorMessege := "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""
	if len(os.Args) < 3 || len(os.Args) > 5 {
		fmt.Println(errorMessege)
		return
	}

	if os.Args[1][:2] != "--" {
		fmt.Println(errorMessege)
		return
	}

	color := flag.String("color", "", "rgb color code")
	flag.Parse()

	str, subS, tempFlag := getInputs(os.Args)

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
		positions = getSubStringPositions(str, subS)
	}

	ansi, err := getAnsi(*color)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errorMessege)
		return
	}
	// str = asciiart.FormatSpChar(str)
	// subS = asciiart.FormatSpChar(subS)

	bannerTemplate := filepath.Join("banners", "standard.txt")
	// arguments := os.Args[0:]

	// Code checks if the length of the argument is exactly two and the second one is not a flag.
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

	// asciiart.ValidateBanner(bannerTemplate)
	artStr := asciiart.WordDistributer(str, bannerTemplate, positions, ansi, lenSub)

	fmt.Println(artStr)
}

func getSubStringPositions(s, sb string) []int {
	// res := "\x1b[0m"
	positions := []int{}
	start := 0

	// geting position indexes of substring
	for start < len(s) {
		pos := strings.Index(s[start:], sb)
		if pos == -1 {
			break
		}
		positions = append(positions, pos+start)
		start += pos + len(sb)
	}
	return positions
}

// func getAnsi() take sin the RGB color code provide and returns a formated ansi format
func getAnsi(s string) (string, error) {
	var rgb Rgb
	var err error

	// common colors
	ansiMap := make(map[string]string)
	ansiMap["red"] = "\033[38;2;255;0;0m"
	ansiMap["green"] = "\033[38;2;0;255;0m"
	ansiMap["blue"] = "\033[38;2;0;0;255m"
	ansiMap["black"] = "\033[38;2;0;0;0m"
	ansiMap["white"] = "\033[38;2;255;255;255m"
	ansiMap["magenta"] = "\033[38;2;255;0;255m"
	ansiMap["cyan"] = "\033[38;2;0;255;255m"
	ansiMap["yellow"] = "\033[38;2;255;255;0m"
	ansiMap["grey"] = "\033[38;2;128;128;128m"
	ansiMap["teal"] = "\033[38;2;0;128;128m"
	ansiMap["navy"] = "\033[38;2;0;0;128m"
	ansiMap["purple"] = "\033[38;2;128;0;128m"
	ansiMap["olive"] = "\033[38;2;128;128;0m"
	ansiMap["maroon"] = "\033[38;2;128;0;0m"
	ansiMap["silver"] = "\033[38;2;192;192;192m"
	ansiMap["orange"] = "\033[38;2;255;165;0m"
	ansiMap["lime"] = "\033[38;2;50;205;50m"
	ansiMap["brown"] = "\033[38;2;165;42;42m"
	ansiMap["darkgreen"] = "\033[38;2;0;100;0m"
	ansiMap["gold"] = "\033[38;2;255;215;0m"

	ansi := ""
	if s[:3] == "rgb" {
		rgb.R, rgb.G, rgb.B, err = getRgbInt(s)
		if err != nil {
			return "", fmt.Errorf("Invalid RGB color code")
		}
		ansi = fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B)
	} else if _, ok := ansiMap[s]; ok {
		ansi = ansiMap[s]
	} else {
		return "", fmt.Errorf("Color given is not defined")
	}

	return ansi, nil
}

// Func getRgbInt() returns the provided rgb color code integers
func getRgbInt(s string) (int, int, int, error) {
	bracket := false
	intSlice := []string{}
	colorNum := ""
	for _, v := range s {
		if v == '(' {
			bracket = true
			continue
		}
		if v == ')' && colorNum != "" {
			intSlice = append(intSlice, colorNum)
			break
		}
		if v == ',' && colorNum != "" {
			intSlice = append(intSlice, colorNum)
			colorNum = ""
			continue
		}
		if bracket {
			if v >= '0' && v <= '9' {
				colorNum += string(v)
			}
		}
	}
	if len(intSlice) != 3 {
		return 0, 0, 0, fmt.Errorf("Invalid RGB values")
	}

	r, _ := strconv.Atoi(intSlice[0])
	g, _ := strconv.Atoi(intSlice[1])
	b, _ := strconv.Atoi(intSlice[2])

	return r, g, b, nil
}

// func getInputs() returns the string, substring and bannerfile flag
func getInputs(sl []string) (string, string, string) {
	str := ""
	subS := ""
	flag := ""

	// go run . --color=[200,34,50] "to" "tommorrow come to my place" "-sh"
	if len(sl) == 5 {
		str = sl[3]
		subS = sl[2]
		flag = sl[4]
	}

	// go run . --color=[200,34,50] "to" "tommorrow come to my place"
	if len(sl) == 4 {
		if sl[3][0] == '-' {
			flag = sl[3]
			str = sl[2]
		} else {
			str = sl[3]
			subS = sl[2]
		}
	}

	// go run . --color=[200,34,50] "tommorrow come to my place"
	if len(sl) == 3 {
		str = sl[2]
	}
	return str, subS, flag
}
