package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

var clients_int chan bool

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()
	for i := 0; i < cap(clients_int); i++ {
		defer fmt.Println(<-clients_int)
	}
	fmt.Println("")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		go handleRequest(conn)
		/**/

	}

} //MAIN END

func handleRequest(conn net.Conn) {

	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// write data to response
	time := time.Now().Format(time.ANSIC)
	fmt.Printf("Server got: %v\n Received time: %v\n", string(buffer[:]), time)
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v\n", string(buffer[:]), time)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
	//<-clients_int
	clients_int <- true
}
