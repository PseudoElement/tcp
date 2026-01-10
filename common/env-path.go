package common

import (
	"os"
	"runtime"
	"strings"
)

func GetEnvPath() string {
	pwd, _ := os.Getwd()

	var splitter string = "/"
	if runtime.GOOS == "windows" {
		splitter = "\\"
	}

	segments := strings.Split(pwd, splitter)

	absolutePath := make([]string, 0, len(segments))
	appRootFound := false
	for _, segment := range segments {
		absolutePath = append(absolutePath, segment)
		if segment == "tcp" {
			appRootFound = true
			break
		}
	}

	if !appRootFound {
		panic("Repository root dir should be called \"tcp\". Otherwise needs to change implementation of getEnvPath() function.")
	}

	envPath := strings.Join(absolutePath, splitter) + splitter + ".env"

	return envPath
}
