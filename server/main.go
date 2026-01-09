package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/pseudoelement/tcp/common"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf(
		"Serving connection from remote %s, local %s\n",
		conn.RemoteAddr().String(),
		conn.LocalAddr().String(),
	)

	// @TODO add loop to read/write many times
	reader := bufio.NewReader(conn)
	buf, _, err := reader.ReadLine()
	// nr, err := conn.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		panic("conn.Read failed: " + err.Error())
	}
	println("Read ", len(buf), " bytes from client.")
	println("Msg from client", string(buf))

	nw, err := conn.Write([]byte("Congrats. Virus working!"))
	if err != nil {
		panic("conn.Write failed: " + err.Error())
	}
	println("Written ", nw, " bytes to client.")
}

func main() {
	envPath := common.GetEnvPath()
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading ../.env file. Error: %v\n", err)
	}

	serverIp := os.Getenv("SERVER_IP")
	serverPort := os.Getenv("SERVER_PORT")
	serverAddress := serverIp + ":" + serverPort

	// Listen on TCP port 8000 on all available IP addresses
	l, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("Server listening on address ", serverAddress)

	for {
		// Wait for a connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}
