package asciiart

import (
	"strings"
)

func FormatSpChar(str string) string {
	str = strings.ReplaceAll(str, "\\n", "\n")
	str = strings.ReplaceAll(str, "\\t", "    ")
	return str
}
