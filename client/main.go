package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

const BUFSIZE = 2048

type Client struct {
	conn    net.Conn
	address string
}

func main() {

	c := NewClient()
	defer c.conn.Close()

	filename := flag.String("filename", "./download.jpeg", "")

	flag.Parse()

	c.SendFile(*filename)
}

func NewClient() Client {

	conn, err := net.Dial("tcp", ":3000")

	if err != nil {
		panic(err)
	}

	return Client{
		conn:    conn,
		address: conn.RemoteAddr().String(),
	}
}

func (c Client) SendFile(filename string) {
	fileHandle, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer fileHandle.Close()

	//buf := make([]byte, BUFSIZE)

	for {

		//nClient, errClient := fileHandle.Read(buf)

		//fmt.Printf("Loaded %d bytes from file: %v\n", nClient, buf[:10])

		nServer, errServer := io.CopyN(c.conn, fileHandle, BUFSIZE)

		fmt.Printf("Sent %d bytes to server\n", nServer)

		if errServer != nil {
			if errServer == io.EOF {
				fmt.Println("SUCCESS!")
				return
			}

			fmt.Printf("server err %v:", errServer)
		}

	}

}
