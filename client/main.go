package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("error connecting to the tcp server")
	}
	defer conn.Close()

	log.Print("connected to tcp server")
}
