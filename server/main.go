package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/pseudoelement/tcp/common"
)

func main() {
	envPath := common.GetEnvPath()
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading ../.env file. Error: %v\n", err)
	}

	serverPort := os.Getenv("SERVER_PORT")
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
