package main

import "fmt"

var mapping = make(map[string]string)

func add() {
	mapping["name"] = "John"
}

func see() {
	add()
}

func main() {
	see()
	fmt.Println(mapping)
}
