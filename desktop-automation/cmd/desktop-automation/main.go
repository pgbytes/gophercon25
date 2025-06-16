package main

import (
	"fmt"
	"os"

	"github.com/pgbytes/gophercon25/desktop-automation/internal/commands"
)

func main() {
	// Initialize and execute the root command
	rootCmd := commands.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
