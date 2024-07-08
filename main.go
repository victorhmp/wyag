package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("A subcommand is required")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "init":
		initCmd.Parse(os.Args[2:])
		fmt.Println(initCmd.Args())
	default:
		fmt.Println("Unknown subcommand")
		os.Exit(1)
	}
}
