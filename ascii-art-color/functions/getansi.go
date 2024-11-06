package asciiart

import (
	"fmt"
	"strconv"
	"strings"
)

type Rgb struct {
	R int
	G int
	B int
}

// func getAnsi() takes in the RGB color code provide the color name, then returns a formated ansi format
func GetAnsi(s string) (string, string, error) {
	var rgb Rgb
	var err error

	// defined colors
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
	ansi2 := ""
	if strings.Contains(s, "rgb") {
		count := 0
		for _, c := range s {
			if c == 'r' {
				count++
			}
		}
		if count > 2 {
			return "", "", fmt.Errorf("only 2 rgb instances are allowed")
		}
		if len(s) <= 3 {
			return "", "", fmt.Errorf("invalid RGB color code")
		}
		if s[3] != '(' {
			return "", "", fmt.Errorf("invalid RGB color code")
		}
		rgb.R, rgb.G, rgb.B, err = getRgbInt(s)
		if err != nil {
			return "", "", err
		}
		ansi = fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B)
		// if second rgb
		ns := s[3:]		
		if strings.Contains(ns, "rgb") {
			ind := strings.Index(ns, "rgb")
			s = ns[ind:]
			if len(s) <= 3 {
				return "", "", fmt.Errorf("invalid RGB color code")
			}
			if s[3] != '(' {
				return "", "", fmt.Errorf("invalid RGB color code")
			}
			rgb.R, rgb.G, rgb.B, err = getRgbInt(s)
			if err != nil {
				return "", "", err
			}
			ansi2 = fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B)
		} else {
			return "", "", fmt.Errorf("invalid RGB color code")
		}
	} else {
		colors := strings.Split(s , ",")
		if len(colors) == 1 {
			if _, ok := ansiMap[colors[0]]; ok {
				ansi = ansiMap[colors[0]]
			} else {
				return "", "", fmt.Errorf("error: color given is not defined")
			}
		} else if len(colors) == 2 {
			 _, ok := ansiMap[colors[0]]
			 _, yes := ansiMap[colors[1]]
			if ok && yes {
				ansi = ansiMap[colors[0]]
				ansi2 = ansiMap[colors[1]]
			} else {
				return "", "", fmt.Errorf("error: color given is not defined")
			}
	} else if len(colors) > 2 {
		return "", "", fmt.Errorf("error: only 2 defined colors allowed")
	}  else {
		return "", "", fmt.Errorf("error: color given is not defined")
	}
}
	return ansi, ansi2, nil
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
		return 0, 0, 0, fmt.Errorf("invalid rgb values")
	}

	r, _ := strconv.Atoi(intSlice[0])
	g, _ := strconv.Atoi(intSlice[1])
	b, _ := strconv.Atoi(intSlice[2])

	if r > 255 || g > 255 || b > 255 {
		return 0, 0, 0, fmt.Errorf("error: color code values should not be more than 255")
	}

	return r, g, b, nil
}
