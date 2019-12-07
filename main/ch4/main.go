package main

import (
	"fmt"
)

func main() {
	cmd := parseCmd()
	if cmd.version {
		fmt.Println("version 0.0.1")
	} else if cmd.help {
		printUsage()
	} else {
		startJvm(cmd)
	}
}
