package main

import (
	"log"
	"time"

	"github.com/pseudoelement/tcp/common"
)

func main() {
	serverPort, serverIp := common.GetServerPortAndIp(true)
	serverAddress := serverIp + ":" + serverPort

	connectWithRetries(serverAddress)
}

func connectWithRetries(serverAddress string) {
	tunnel := NewTcpTunnelClient()

	err := tunnel.Connect(serverAddress)
	if err != nil {
		log.Println("[connectWithRetries] tunnel.Connect failed. Error:", err)
		time.Sleep(time.Second * 2)

		connectWithRetries(serverAddress)
		return
	}

	err = tunnel.Run()
	if err != nil {
		log.Println("[connectWithRetries] tunnel.Run failed. Error:", err)
	}
	time.Sleep(time.Second * 2)

	connectWithRetries(serverAddress)
}
