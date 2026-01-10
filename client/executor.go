package main

import (
	"os/exec"
	"strings"
)

func execute(msg string) (outMsg string, err error) {
	command := strings.TrimPrefix(msg, "$ ")
	cmdSegments := strings.Split(command, " ")

	cmd := exec.Command(cmdSegments[0], cmdSegments[1:]...)
	out, err := cmd.Output()
	outMsg = string(out)

	return outMsg, err
}
