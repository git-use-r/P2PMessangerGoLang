package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

var serverConn net.Conn
var serverConfig tls.Config
var serverCert tls.Certificate

func InitializeServer() {

	var err1 error
	serverCert, err1 = GenerateCert()
	if err1 != nil {
		log.Fatalf("sertificate error: %v", err1)
		return
	}

	serverConfig = tls.Config{
		Certificates: []tls.Certificate{serverCert},
		MinVersion:   tls.VersionTLS13,
	}

}

func ServerRun(port string) {

	listiner, err1 := tls.Listen("tcp", port, &serverConfig)
	if err1 != nil {
		log.Fatal("err open tls session: ", err1)
	}

	fmt.Println("tls server open! port: " + port)

	var err2 error
	serverConn, err2 = listiner.Accept()
	if err2 != nil {
		log.Fatal("err open tls accept: ", err2)
	}

	fmt.Println("Client Connect")

	go BackReadMsgSER()
}

func BackReadMsgSER() {

	for {
		buf := make([]byte, 1024)

		n, err := serverConn.Read(buf)
		if err != nil {
			log.Fatal("error read msg: ", err)
		}

		WiriteArea(fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), string(buf[:n])), "remoteUserMsg")

		WiriteArea(fmt.Sprintf("get - [%s] addres - [%s]", string(buf[:n]), serverConn.RemoteAddr()), "logMenu")
	}

}

func WriteServer(msg string) {
	serverConn.Write([]byte(msg))
	WiriteArea(fmt.Sprintf("send - [%s] addres - [%s]", msg, serverConn.RemoteAddr()), "logMenu")
}
