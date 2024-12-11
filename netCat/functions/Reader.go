package netcat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

var History []string

func Reader(conn net.Conn) {
	count := 0
	var client Client
	client.conn = conn
	entry, err := Readfile("./src/logo.txt")
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("Internal server error: %v", err)))
		conn.Close()
		return
	}

	conn.Write([]byte(entry))
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Write([]byte("Connection Denied"))
		conn.Close()
		return
	}

	name = strings.TrimSpace(name)
	for !CheckName(name) {
		if count == 5 {
			conn.Write([]byte("Connection Denied"))
			conn.Close()
			return
		}
		conn.Write([]byte("Invalid name try another name\n[ENTER YOUR NAME]: ]"))
		count++
	}
	ent := fmt.Sprintf("%v has joined the chat\n", name)
	client.name = name
	Chat(conn, ent)
	Clients = append(Clients, client)

	if len(History) != 0 {
		history := strings.Join(History, "")
		conn.Write([]byte(history))
	}

	for {
		ext := fmt.Sprintf("%v has left the chat\n", client.name)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			Chat(conn, ext)
			Clients = Remove(Clients, conn)
			conn.Close()
			break
		}

		message = strings.TrimSpace(message)
		if message != "" {
			if strings.Contains(message, "--cn") {
				old := client.name
				Clients, client, err = ChangeName(client)
				if err != nil {
					client.conn.Write([]byte("Connection Denied"))
					client.conn.Close()
					Chat(client.conn, ext)
					Clients = Remove(Clients, client.conn)
					break
				}
				cn := fmt.Sprintf("%v changed name to %v\n", old, client.name)
				Chat(conn, cn)
				continue
			}

			text := fmt.Sprintf("[%v][%v]: %v\n", time.Now().Format("2006-01-02 15:04:05"), client.name, message)
			History = append(History, text)
			conn.Write([]byte("\033[F\033[K"))
			conn.Write([]byte(text))
			Chat(client.conn, text)
		}
	}
}
