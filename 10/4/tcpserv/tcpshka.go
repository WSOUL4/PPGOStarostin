package main

import (
	"fmt"
	"log"
	"net"

	//"os"
	"time"
	//"tcpserver"
	"crypto/rand"
	"crypto/tls"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

var clients_int chan bool

func main() {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")

	if err != nil {
		log.Fatal(err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAnyClientCert}
	config.Rand = rand.Reader

	listen, err := tls.Listen(TYPE, HOST+":"+PORT, &config) //СОЗДАЁМ СЕРВЕР ЧЕРЕЗ ТЛС ,КОТОРЫЙ ПРИНИМАЕТ КОНФИГ С КЛЮЧОМ И СЕРТИФИКАТОМ
	if err != nil {
		log.Fatal(err)
		return
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
			return
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
