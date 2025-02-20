package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"forum/handlers"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: go run . [PORT]")
	}
	portStr := ":8080"
	if len(args) == 2 {
		portStr = args[1]
		if strings.Contains(portStr, ":") {
			portStr = portStr[1:]
		}
		portVal, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Println("Invalid port number input.")
			return
		}
		if portVal < 1024 || portVal > 65535 {
			fmt.Printf("Port :%d is not available. Available ports are between :1023 and :65536\n", portVal)
			return
		}
		portStr = ":"+portStr
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r)
	})
	log.Printf("Server running on: http://localhost%s\n", portStr)
	http.ListenAndServe(portStr, nil)
}
