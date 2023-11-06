package main

/* author Samuel
 */

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Listener() {
	fmt.Print("listening on port 8080...\n")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection: %v", err)
		}
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	//buffer to accept incoming data
	buffer := make([]byte, 1024)
	content, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	//used to extract filenmameme form data received
	filename := buffer[:content]
	file, err := os.Create(string(filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, conn)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Recieved file succesfully: %s\n", filename)
}
