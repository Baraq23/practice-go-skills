package netcat

func CheckName(name string) bool {
	if name == "" {
		return false
	}
	for _, cl := range Clients {
		if cl.name == name {
			return false
		}
	}
	return true
}
