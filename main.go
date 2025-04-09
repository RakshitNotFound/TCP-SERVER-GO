package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Processing the request")
	time.Sleep(8 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHELLO, WORLD\r\n"))
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server is running on port 1729...")

	for {
		log.Println("Waiting for client to connect")
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Client connected")

		go do(conn)
	}
}
