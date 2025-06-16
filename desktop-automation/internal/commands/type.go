// Package commands implements the CLI commands for desktop automation
package commands

import (
	"fmt"

	"github.com/pgbytes/gophercon25/desktop-automation/internal/automation"
	"github.com/spf13/cobra"
)

// newTypeCmd creates the type command
func newTypeCmd() *cobra.Command {
	var delayMs int

	typeCmd := &cobra.Command{
		Use:   "type [text]",
		Short: "Type text on the keyboard",
		Long:  `Simulate keyboard typing of the specified text.`,
		Example: `  # Type "Hello World!"
  desktop-automation type 'Hello World!'
  
  # Type with delay between keystrokes
  desktop-automation type --delay=50 'Slow typing!'`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			text := args[0]

			// Create keyboard automation instance
			keyboard := automation.NewKeyboard()

			var err error
			if delayMs > 0 {
				err = keyboard.TypeStringWithDelay(text, delayMs)
			} else {
				err = keyboard.TypeString(text)
			}

			if err != nil {
				return err
			}

			// Show success message with character count
			fmt.Printf("Successfully typed %d characters\n", len(text))
			return nil
		},
	}

	// Add delay flag
	typeCmd.Flags().IntVar(&delayMs, "delay", 0, "Delay between keystrokes in milliseconds")

	return typeCmd
}
