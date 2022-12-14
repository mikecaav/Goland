package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer func() {
		log.Print("Closing connections")
		listener.Close()
	}()
	log.Println("Listening on " + CONN_HOST)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting", err.Error())
		}
		log.Print(conn)
		break
	}
}
