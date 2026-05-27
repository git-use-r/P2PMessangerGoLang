package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

var clientConn net.Conn
var clientConfig tls.Config
var clientCert tls.Certificate

func InitializeClient() {

	var err1 error
	clientCert, err1 = GenerateCert()
	if err1 != nil {
		log.Fatalf("sertificate error: %v", err1)
		return
	}

	clientConfig = tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		MinVersion:         tls.VersionTLS13,
		InsecureSkipVerify: true,
	}

}

func ClientRun(addr string) {

	var err error
	clientConn, err = tls.Dial("tcp", addr, &clientConfig)
	if err != nil {
		log.Fatalf("sertificate error: %v", err)
		return
	}

	log.Println("TLS connect!")

	go BackReadMsgCL()
}

func BackReadMsgCL() {

	for {
		buf := make([]byte, 1024)

		n, err := clientConn.Read(buf)
		if err != nil {
			log.Fatal("error read msg: ", err)
		}

		WiriteArea(fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), string(buf[:n])), "remoteUserMsg")

		WiriteArea(fmt.Sprintf("get - [%s] addres - [%s]", string(buf[:n]), clientConn.RemoteAddr()), "logMenu")
	}

}

func WriteClient(msg string) {
	clientConn.Write([]byte(msg))
	WiriteArea(fmt.Sprintf("send - [%s] addres - [%s]", msg, clientConn.RemoteAddr()), "logMenu")
}
