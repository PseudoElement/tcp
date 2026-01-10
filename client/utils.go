package main

import "strings"

func isCommand(serverMsg string) bool {
	return strings.HasPrefix(serverMsg, "$ ")
}
