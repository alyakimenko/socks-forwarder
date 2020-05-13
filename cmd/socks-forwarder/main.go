package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const version = "0.01"

func main() {
	config := parseFlags()

	if *config.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	switch strings.ToLower(*config.LogLevel) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		panic("unsupport logging level")
	}
}
