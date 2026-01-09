package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"

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

	println("Server address: ", serverAddress)

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Error connecting to %s. Error: %v\n", serverAddress, err)
	}
	defer conn.Close()

	fmt.Printf("Connected to server at %s\n", serverAddress)

	randomInt := strconv.Itoa(rand.Intn(1000))
	// Example: send a message
	if _, err := conn.Write([]byte("Hello from a different IP " + randomInt + "!\n")); err != nil {
		log.Fatal(err)
	}

	// @TODO add infinite loop with reading incoming messages
	// Read response (optional)
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s\n", string(buf[:n]))
}
