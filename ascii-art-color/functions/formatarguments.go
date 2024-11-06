package asciiart

// func getInputs() returns the string, substring and bannerfile flag
func GetInputs(sl []string) (string, string, string) {
	str := ""
	subS := ""
	flag := ""

	// go run . --color="rgb(200,34,50)" "to" "some string goes here" "-sh"
	if len(sl) == 5 {
		str = sl[3]
		subS = sl[2]
		flag = sl[4]
	}

	// go run . --color="rgb(200,34,50)" "to" "some string goes here"
	if len(sl) == 4 {
		if sl[3][0] == '-' {
			flag = sl[3]
			str = sl[2]
		} else {
			str = sl[3]
			subS = sl[2]
		}
	}

	// go run . --color="rgb(200,34,50)" "some string goes here"
	if len(sl) == 3 {
		str = sl[2]
	}
	return str, subS, flag
}