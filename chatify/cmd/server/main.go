package main

import "gopher-playground/chatify/pkg/chatify"

func main() {
	server := chatify.NewServer()
	server.Run()
}
