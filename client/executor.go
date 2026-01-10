package main

import (
	"os/exec"
	"runtime"
	"strings"
)

func execute(msg string) (outMsg string, err error) {
	command := strings.TrimPrefix(msg, "$ ")

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-Command", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	out, err := cmd.Output()
	outMsg = string(out)

	return outMsg, err
}
