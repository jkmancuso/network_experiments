package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const BUFSIZE = 2048

type Server struct {
	listener net.Listener
	address  string
}

func main() {

	s := NewServer()
	log.Fatal(s.StartServer())
}

func (s Server) StartServer() error {

	var conn net.Conn
	var err error

	defer s.listener.Close()

	for {

		conn, err = s.listener.Accept()

		if err != nil {
			return err
		}

		fmt.Println("New conn")

		go s.handleConn(conn)

	}
}

func (s Server) handleConn(conn net.Conn) {
	defer conn.Close()

	fileHandle, err := os.Create("upload.jpeg")

	if err != err {
		panic(err)
	}

	for {
		n, err := io.CopyN(fileHandle, conn, BUFSIZE)

		fmt.Printf("Got %d bytes:\n", n)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func NewServer() Server {

	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		panic(err)
	}

	return Server{
		listener: listener,
		address:  listener.Addr().String(),
	}
}
