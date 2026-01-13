package common

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetServerPortAndIp(fromClient bool) (port string, serverIp string) {
	for _, arg := range os.Args {
		p, hasPort := strings.CutPrefix(arg, "--port=")
		ip, hasIp := strings.CutPrefix(arg, "--ip=")
		if hasPort {
			port = p
		}
		if hasIp {
			serverIp = ip
		}
	}

	if port != "" && serverIp != "" {
		return port, serverIp
	}

	godotenv.Load(".env")

	p := os.Getenv("SERVER_PORT")
	ip := os.Getenv("SERVER_IP")

	if port == "" && p != "" {
		port = p
	} else {
		port = "8228"
	}
	if serverIp == "" && ip != "" {
		serverIp = ip
	} else {
		serverIp = "82.146.32.19"
	}

	return port, serverIp
}
