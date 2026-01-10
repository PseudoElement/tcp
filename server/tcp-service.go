package main

import (
	"log"
	"net"
)

type TcpService struct {
	uniqueInt int
	listener  net.Listener

	Tunnels map[TunnelID]*TcpTunnel
}

func NewTcpService(l net.Listener) *TcpService {
	return &TcpService{
		listener:  l,
		uniqueInt: 0,
		Tunnels:   make(map[TunnelID]*TcpTunnel),
	}
}

func (t *TcpService) ListenConnectionRequests() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Fatal("[TcpService_ListenConnectionRequests] t.listener.Accept failed: ", err)
		}

		id := t.uniqueInt
		tunnel := NewTcpTunnel(conn, id)
		t.Tunnels[id] = tunnel
		t.uniqueInt++

		go tunnel.Run()
		go t.handleTunnelClosing(tunnel)

		log.Println("Tunnel with id ", id, " created.")
	}
}

func (t *TcpService) handleTunnelClosing(tunnel *TcpTunnel) {
	<-tunnel.CloseChan()
	delete(t.Tunnels, tunnel.Id())
	tunnel = nil
}
