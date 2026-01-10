package main

import (
	"errors"
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

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Error connecting to %s. Error: %v\n", serverAddress, err)
	}
	defer conn.Close()

	fmt.Printf("Connected to server at %s\n", serverAddress)

	randomInt := strconv.Itoa(rand.Intn(1000))
	if _, err := conn.Write([]byte("Hello from a different IP " + randomInt + "!\n")); err != nil {
		log.Fatal(err)
	}

	// @TODO evaluate message starting with $ as command
	// otherwise just type text in new file
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("[client_main] server closed connection")
				return
			}
			log.Println("[client_main] read failed: ", err.Error())
			return
		}

		fmt.Printf("Received msg: %s\n", string(buf[:n]))
	}
}
