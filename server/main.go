package main

import (
	"log"
	"net"
)

func main() {
	srv, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("error starting tcp server: ", err)
	}
	defer srv.Close()

	log.Println("tcp server started on port 8080")
	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Println("connection error: ", err)
			continue
		}

		log.Print("new connection: ", conn)
	}
}
