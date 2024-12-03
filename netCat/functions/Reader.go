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
	var client Client
	client.conn = conn
	entry, err := Readfile("./src/logo.txt")
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	conn.Write([]byte(entry))
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Write([]byte("Connection Denied"))
		return
	}
	name = strings.TrimSpace(name)
	client.name = name
	ent := fmt.Sprintf("%v has joined the chat\n", name)
	ext := fmt.Sprintf("%v has left the chat\n", name)
	Chat(conn, ent)
	Clients = append(Clients, client)

	if len(History) != 0 {
		history := strings.Join(History, "")
		conn.Write([]byte("\nRecent Chats\n" + history + "\nNew Chats\n"))
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			Chat(conn, ext)
			Clients = Remove(Clients, conn)
			conn.Close()
			break
		}
		text := fmt.Sprintf("[%v][%v]: %v", time.Now().Format("2006-01-02 15:04:05"), client.name, message)
		History = append(History, text)
		Chat(client.conn, text)
	}
}
