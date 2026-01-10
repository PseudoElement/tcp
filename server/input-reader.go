package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type InputReader struct {
	tunnels            map[TunnelID]*TcpTunnel
	targetTunnelId     TunnelID
	currentServerInput ServerInput
}

func NewInputReader(tunnels map[TunnelID]*TcpTunnel) *InputReader {
	return &InputReader{
		tunnels:            tunnels,
		targetTunnelId:     -1,
		currentServerInput: TUNNEL_ID,
	}
}

func (i *InputReader) Run() {
	log.Println("Input tunnel id:")

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input := in.Text()

		if input == ":qa" {
			break
		}

		if i.currentServerInput == TUNNEL_ID {
			id, err := strconv.Atoi(input)
			if err != nil {
				log.Printf("Id %v is invalid.\n", input)
				continue
			}

			if _, ok := i.tunnels[id]; !ok {
				log.Printf("Tunnel with id %d not found.\n", id)
				continue
			} else {
				log.Printf("Target tunnel with id %d set.\n", id)
				log.Println("Input command to client:")

				i.targetTunnelId = id
				i.currentServerInput = COMMAND
			}
		} else {
			if input == ":tunnel" {
				i.currentServerInput = TUNNEL_ID
				log.Println("Input tunnel id:")
			} else {
				if tunnel, ok := i.tunnels[i.targetTunnelId]; ok {
					tunnel.CommandToClientChan() <- input
				} else {
					log.Printf("Tunnel with id %d not found.\n", i.targetTunnelId)
					i.currentServerInput = TUNNEL_ID
					log.Println("Input tunnel id:")
				}

			}
		}

	}
}
