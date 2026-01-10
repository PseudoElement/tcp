package main

import (
	"os/exec"
	"strings"
)

func execute(msg string) (outMsg string, err error) {
	command := strings.TrimPrefix(msg, "$ ")

	// Use shell to interpret pipes, redirects, etc.
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.Output()
	outMsg = string(out)

	return outMsg, err
}
