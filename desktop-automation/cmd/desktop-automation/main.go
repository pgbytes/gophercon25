package main

import (
	"fmt"
	"os"

	"github.com/pgbytes/gophercon25/desktop-automation/internal/commands"
)

var (
	version = "v0.1.0"
)

func main() {
	// Initialize the root command
	rootCmd := commands.NewRootCmd()

	// Add version flag
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("desktop-automation {{.Version}}\n")

	// Execute the root command and handle errors gracefully
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
