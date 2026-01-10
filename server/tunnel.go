package main

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/pseudoelement/tcp/common"
)

type TcpTunnel struct {
	clientConn          net.Conn
	commandToClientChan chan CommandToClient
	closeChan           chan struct{}
	id                  int
}

func NewTcpTunnel(conn net.Conn, id int) *TcpTunnel {
	log.Printf(
		"Opened tcp tunnel with remote %s, local %s\n",
		conn.RemoteAddr().String(),
		conn.LocalAddr().String(),
	)

	return &TcpTunnel{
		clientConn:          conn,
		closeChan:           make(chan struct{}),
		commandToClientChan: make(chan CommandToClient),
		id:                  id,
	}
}

func (t *TcpTunnel) Id() int {
	return t.id
}

func (t *TcpTunnel) Close() error {
	t.closeChan <- struct{}{}
	return t.clientConn.Close()
}

func (t *TcpTunnel) CommandToClientChan() chan<- CommandToClient {
	return t.commandToClientChan
}

func (t *TcpTunnel) CloseChan() <-chan struct{} {
	return t.closeChan
}

func (t *TcpTunnel) Run() {
	defer t.Close()

	go t.writeMessagesToClient()
	t.readDataFromClient()
}

func (t *TcpTunnel) readDataFromClient() {
	reader := bufio.NewReader(t.clientConn)
	for {
		buf, err := reader.ReadBytes(common.END_OF_MSG)
		if err != nil {
			if err == io.EOF {
				log.Println("[TcpTunnel_readDataFromClient] client closed connection")
				return
			}
			log.Println("[TcpTunnel_readDataFromClient] read failed: ", err.Error())
			return
		}

		log.Println("Data from client: ", string(buf))
	}
}

func (t *TcpTunnel) writeMessagesToClient() {
	for {
		select {
		case <-t.closeChan:
			return
		case commandToClient := <-t.commandToClientChan:
			_, err := t.clientConn.Write([]byte(commandToClient + "\n"))
			if err != nil {
				println("[TcpTunnel_writeMessagesToClient] t.clientConn.Write failed: " + err.Error())
			}

			log.Printf("\"%s\" sent to client", commandToClient)
		}
	}
}
