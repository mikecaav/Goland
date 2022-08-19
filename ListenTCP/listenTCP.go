package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func handleConnection(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(message)
	conn.Close()
}

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		log.Println("Connected to: "+CONN_TYPE, CONN_HOST+":"+CONN_PORT)
		if err != nil {
			log.Fatal("Error: " + err.Error())
		}
		go handleConnection(conn)
	}
}
