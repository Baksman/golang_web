package main

import (
	"fmt"
	// "io"
	"time"
	"bufio"
	"log"
	"net"
)

func main(){
	li,err := net.Dial("tcp","localhost:8080")
	if err != nil{
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn,err := li.Accept()

	}
}