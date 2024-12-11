package netcat

import (
	"bufio"
	"fmt"
	"strings"
)

func ChangeName(client Client) ([]Client, Client, error) {
	count := 0
	client.conn.Write([]byte("[ENTER NEW NAME]: "))
	name, err := bufio.NewReader(client.conn).ReadString('\n')
	if err != nil {
		return Clients, client, fmt.Errorf("connection closed")
	}
	name = strings.TrimSpace(name)

	for !CheckName(name) {
		if count == 5 {
			return Clients, client, fmt.Errorf("connection closed")
		}
		client.conn.Write([]byte("Invalid name try another name\n[ENTER NEW NAME]: ]"))
		count++
	}

	for i := 0; i < len(Clients); i++ {
		if Clients[i].conn == client.conn {
			Clients[i].name = name
			client.name = name
		}
	}
	return Clients, client, nil
}
