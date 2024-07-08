package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
	// checkIgnoreCmd := flag.NewFlagSet("check-ignore", flag.ExitOnError)
	// checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)
	// commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	// hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	// logCmd := flag.NewFlagSet("log", flag.ExitOnError)
	// lsFilesCmd := flag.NewFlagSet("ls-files", flag.ExitOnError)
	// lsTreeCmd := flag.NewFlagSet("ls-tree", flag.ExitOnError)
	// revParseCmd := flag.NewFlagSet("rev-parse", flag.ExitOnError)
	// rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)
	// showRefCmd := flag.NewFlagSet("show-ref", flag.ExitOnError)
	// statusCmd := flag.NewFlagSet("status", flag.ExitOnError)
	// tagCmd := flag.NewFlagSet("tag", flag.ExitOnError)

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
