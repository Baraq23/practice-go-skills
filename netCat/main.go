package main

import (
	"fmt"
	"os"

	netcat "netcat/functions"
)

func main() {
	args := os.Args
	var port string
	if len(args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(args) == 2 {
		port = args[1]
	}
	netcat.Server(port)
}
