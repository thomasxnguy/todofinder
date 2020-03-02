package main

import (
	"os"
	"path/filepath"
	"fmt"
	"flag"
	"github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/cmd/search"
	"github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/cmd/server"
)

// Structure containing todofinder modes and run methods.
var (
	mode = []struct {
		Name        string
		Description string
		Run         func([]string) error
	}{
		{
			search.CommandName, search.Description, search.Run,
		},
		{
			server.CommandName, server.Description, server.Run,
		},
	}
)

// Usage will display todofinder usage message.
func Usage() {
	fmt.Printf("todofinder\n")
	fmt.Printf("usage: %s <command>\n", filepath.Base(os.Args[0]))
}

// PrintDefaults prints out the default flags for each commands.
func PrintDefaults() {
	fmt.Printf("The commands are:\n")
	for _, m := range mode {
		if m.Name == search.CommandName {
			fmt.Printf("   [%s] - %s (*default)\n", search.CommandName, search.Description)
		} else {
			fmt.Printf("   %s - %s\n", server.CommandName, server.Description)
		}
	}
}

func main() {
	var (
		cmd     func([]string) error
		options []string
	)
	// Command mode is provided
	if len(os.Args) >= 2 {
		options = os.Args[2:]
		for i := range mode {
			if os.Args[1] == mode[i].Name {
				cmd = mode[i].Run
			}
		}
	}
	// Use default command (search)
	if cmd == nil {
		options = os.Args[1:]
		if e := search.OptionCheck(options); e != nil {
			Usage()
			PrintDefaults()
			fmt.Printf("\n")
			search.Usage()
			search.PrintDefaults(flag.ExitOnError)
			os.Exit(1)
		}
		cmd = search.Run
	}
	// Execute command
	if e := cmd(options); e != nil {
		fmt.Printf("%v\n", e)
		os.Exit(1)
	}
}
