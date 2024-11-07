package funcs

// func getArt() prints the art fron the 2d slice
func GetArt(store [][]string) string {
	word := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(store); j++ {
			word += store[j][i]
		}
		if i < 7 {
			word += "\n"
		}
	}
	return word
}
