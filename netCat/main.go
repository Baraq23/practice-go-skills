package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Server starts a TCP server and handles client connections
func Server(port string) {
	if port == "" {
		port = "8989"
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer ln.Close()
	fmt.Println("Server listening on port " + port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

// handleConnection manages a single client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Reader goroutine to handle incoming messages
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("Client: %s\n", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading from client: %v", err)
		}
	}()

	// Main loop for sending messages to the client
	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You can start typing messages:")
	for scanner.Scan() {
		message := scanner.Text() + "\n"
		_, err := writer.WriteString(message)
		if err != nil {
			log.Printf("Error writing to client: %v", err)
			break
		}
		writer.Flush()
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v", err)
	}
}

func main() {
	port := "8989" // Change this to use a different port
	Server(port)
}
