package main

import (
	"fmt"
	"os"
)

const version = "0.01"

func main() {
	config := parseFlags()

	if *config.Version {
		fmt.Println(version)
		os.Exit(0)
	}
}
