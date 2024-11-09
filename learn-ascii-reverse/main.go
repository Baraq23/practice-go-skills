package main

import (
	"flag"
	"fmt"
	"os"

	// "strings"

	asciiart "rev/ascii-fs-func"
	"rev/functions"
)

func main() {
	revfile := flag.String("reverse", "", "Name of 'txt' file containing the ascii art.")
	bannerTempl := flag.String("banner", "standard", "Name of banner template to be used.")
	flag.Parse()

	// check if the length of the argument after the flags is only 1.
	if flag.NArg() > 1 {
		fmt.Println("error: wrong input format.")
		fmt.Println("Usage: go run . [OPTION] [SOME TEXT]\n\nEX: go run . --banner=<banner-template> \"some text.\" \n[OPTIONS]: <shadow, thinkertoy, standard>")
		return
	}

	if flag.NArg() == 1 {
		input := flag.Arg(0)

		nl := asciiart.FormatSpChar(input)
		if newLinesOnly(nl){
			for i := 0; i < len(nl); i++ {
				fmt.Println()
			}
			return
		}

		if *bannerTempl != "shadow" && *bannerTempl != "thinkertoy" && *bannerTempl != "standard" {
			fmt.Println("error: undefined banner template")
			fmt.Println("Usage: go run . [OPTION] [TEXT]\n\nEX: go run . --output=<fileName> --banner=<bannerTemplate>\n[OPTIONS]: <shadow, thinkertoy, standard>")
			return
		}

		bannerStyle := "banners/"+*bannerTempl+".txt"
		asciiArt, err := asciiart.GetAsciiArt(input, bannerStyle)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(asciiArt)
	}

	
	if *revfile != "" {
		if len(os.Args) != 2 {
			fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
			return
		}
		if os.Args[1][0:2] != "--" {
			fmt.Println("error: incorrect flag format.")
			fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
			return
		}
		reversedText, err := funcs.Reverse(*revfile)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
			return
		}
		fmt.Println(reversedText)
	}
}


func newLinesOnly(s string) bool {
	for _, nl := range s {
		if nl != '\n' {
			return false
		}
	}
	return true
}