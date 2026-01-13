package main

import (
	"log"
	"net"

	"github.com/pseudoelement/tcp/common"
)

func main() {
	serverPort, _ := common.GetServerPortAndIp(false)
	serverAddress := ":" + serverPort

	l, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("Server listening on address", serverAddress)

	tcpService := NewTcpService(l)
	inputReader := NewInputReader(tcpService.Tunnels)

	go inputReader.Run()
	tcpService.ListenConnectionRequests()
}
