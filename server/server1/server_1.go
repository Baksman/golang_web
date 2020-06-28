package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}
		io.WriteString(conn, "Hello from yor world")
		fmt.Println(conn, "how is your day going")
		fmt.Fprintf(conn, "%v", "Well i hope you are okay")
		conn.Close()
	}
}
