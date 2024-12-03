package netcat

import "net"

func Chat(conn net.Conn, msg string) {
	for _, cl := range Clients {
		if cl.conn != conn {
			cl.conn.Write([]byte(msg))
		}
	}
}
