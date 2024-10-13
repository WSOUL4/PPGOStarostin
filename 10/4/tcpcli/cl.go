package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")

	if err != nil {
		log.Fatal(err)
	}
	var pmes string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Mes to serv:")
	pmes, _ = reader.ReadString('\n')

	//fmt.Println("Mes to serv:")
	//fmt.Scanf("%s\n", &pmes)
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true} //ТУТ ПРИКОЛ , НАШ СЕРТИФИКАТ НИФИГА НЕ НОМРАЛЬНЫЙ ПОЭТОМУ ПРОВЕРКУ ОТКЛЮЧИТЬ
	conn, err := tls.Dial(TYPE, HOST+":"+PORT, &config)
	if err != nil {
		log.Fatal(err)
		return
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
	//conn.Close()
	defer conn.Close()
}
