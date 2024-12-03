package netcat

import "net"

func Remove(clients []Client, conn net.Conn) (res []Client) {
	for _, cl := range clients {
		if cl.conn != conn {
			res = append(res, cl)
		}
	}
	return
}
