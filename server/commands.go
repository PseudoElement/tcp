package main

type CommandToClient = string

type ServerInput = int

const (
	TUNNEL_ID ServerInput = 0
	COMMAND   ServerInput = 1
)
