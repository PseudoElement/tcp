package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pseudoelement/tcp/common"
)

func main() {
	envPath := common.GetEnvPath()
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading ../.env file. Error: %v\n", err)
	}

	serverIp := os.Getenv("SERVER_IP")
	serverPort := os.Getenv("SERVER_PORT")
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
	}

	err = tunnel.Run()
	if err != nil {
		log.Println("[connectWithRetries] tunnel.Run failed. Error:", err)
	}
	time.Sleep(time.Second * 2)

	connectWithRetries(serverAddress)
}
