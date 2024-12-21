package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"learn.zone01kisumu.ke/git/barraotieno/ascii-art-web-export-file/serverhandlers"
)

func main() {
	if !(len(os.Args) >= 1 && len(os.Args) < 3) {
		fmt.Println("Usage: go run . [:PORT] OR Usage: go run .")
		return
	}

	port := ""

	if len(os.Args) == 1 {
		port = "9000"
	}

	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs"))))

	http.HandleFunc("/", serverhandlers.HomeHandler)
	http.HandleFunc("/ascii-art", serverhandlers.AsciiArtHandler)
	http.HandleFunc("/about.html", serverhandlers.AboutHandler)
	http.HandleFunc("/download", serverhandlers.DownloadAsciiArtHandler)

	portNum, err := strconv.Atoi(port)
	if err != nil {
		fmt.Printf("Cannot convert %v to integer: %v\n", port, err)
	}

	// Check if portNum is within a valid range
	if portNum < 1024 || portNum > 65535 {
		fmt.Println("Port to running the server should be within the range 1024 to 65535")
		return
	}
	port = ":" + strconv.Itoa(portNum)
	log.Printf("Server running on http://localhost%v\n", port)
	http.ListenAndServe(port, nil)
}
