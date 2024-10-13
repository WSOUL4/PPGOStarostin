package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	var pmes string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Mes to serv:")
	pmes, _ = reader.ReadString('\n')

	//fmt.Println("Mes to serv:")
	//fmt.Scanf("%s\n", &pmes)
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err.Error())
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP(TYPE, nil, tcpServer)

	if err != nil {
		log.Fatal(err.Error())
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte(pmes))
	if err != nil {
		log.Fatal(err.Error())
		fmt.Println("Write data failed:", err.Error())
		os.Exit(1)
	}
	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		log.Fatal(err.Error())
		fmt.Println("Read data failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Received message:", string(received))
	conn.Close()

}
