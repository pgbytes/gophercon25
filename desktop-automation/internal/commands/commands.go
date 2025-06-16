// Package commands implements the CLI commands for desktop automation
package commands

import (
	"fmt"

	"github.com/pgbytes/gophercon25/desktop-automation/internal/automation"
	"github.com/spf13/cobra"
)

// NewRootCmd creates the root command for the desktop-automation CLI
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "desktop-automation",
		Short: "Beautiful Desktop Automation CLI",
		Long:  `A command-line interface for automating desktop interactions using Go.`,
	}

	// Add subcommands
	rootCmd.AddCommand(newClickCmd())
	rootCmd.AddCommand(newTypeCmd())
	rootCmd.AddCommand(newMoveCmd())

	return rootCmd
}

// newClickCmd creates the click command
func newClickCmd() *cobra.Command {
	clickCmd := &cobra.Command{
		Use:   "click x y",
		Short: "Click at a specific screen coordinate",
		Long:  `Click mouse at the specified X and Y coordinates on your screen.`,
		Example: `  # Click at position (100, 200)
  desktop-automation click 100 200`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var x, y int
			if _, err := fmt.Sscanf(args[0], "%d", &x); err != nil {
				return fmt.Errorf("invalid x coordinate: %s", args[0])
			}
			if _, err := fmt.Sscanf(args[1], "%d", &y); err != nil {
				return fmt.Errorf("invalid y coordinate: %s", args[1])
			}

			// Implementation using automation package
			mouse := automation.NewMouse()
			if err := mouse.MoveTo(x, y); err != nil {
				return fmt.Errorf("failed to move mouse: %w", err)
			}
			if err := mouse.Click(); err != nil {
				return fmt.Errorf("failed to click: %w", err)
			}

			fmt.Printf("Clicked at coordinates (%d, %d)\n", x, y)
			return nil
		},
	}

	return clickCmd
}

// newTypeCmd creates the type command
func newTypeCmd() *cobra.Command {
	typeCmd := &cobra.Command{
		Use:   "type \"text\"",
		Short: "Type text at current cursor position",
		Long:  `Simulate keyboard typing of the specified text at the current cursor position.`,
		Example: `  # Type "Hello, World!"
  desktop-automation type "Hello, World!"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			text := args[0]

			// Implementation using automation package
			keyboard := automation.NewKeyboard()
			if err := keyboard.Type(text); err != nil {
				return fmt.Errorf("failed to type text: %w", err)
			}

			fmt.Printf("Typed: %s\n", text)
			return nil
		},
	}

	return typeCmd
}

// newMoveCmd creates the move command
func newMoveCmd() *cobra.Command {
	moveCmd := &cobra.Command{
		Use:   "move x y",
		Short: "Move the mouse cursor to coordinates",
		Long:  `Move the mouse cursor to the specified X and Y coordinates on your screen.`,
		Example: `  # Move cursor to position (500, 300)
  desktop-automation move 500 300`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var x, y int
			if _, err := fmt.Sscanf(args[0], "%d", &x); err != nil {
				return fmt.Errorf("invalid x coordinate: %s", args[0])
			}
			if _, err := fmt.Sscanf(args[1], "%d", &y); err != nil {
				return fmt.Errorf("invalid y coordinate: %s", args[1])
			}

			// Implementation using automation package
			mouse := automation.NewMouse()
			if err := mouse.MoveTo(x, y); err != nil {
				return fmt.Errorf("failed to move mouse: %w", err)
			}

			fmt.Printf("Moved cursor to coordinates (%d, %d)\n", x, y)
			return nil
		},
	}

	return moveCmd
}
