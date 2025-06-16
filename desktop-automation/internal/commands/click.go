// Package commands implements the CLI commands for desktop automation
package commands

import (
	"fmt"
	"strconv"

	"github.com/pgbytes/gophercon25/desktop-automation/internal/automation"
	"github.com/spf13/cobra"
)

// newClickCmd creates the click command
func newClickCmd() *cobra.Command {
	clickCmd := &cobra.Command{
		Use:   "click x y",
		Short: "Click at a specific screen coordinate",
		Long:  `Click mouse at the specified X and Y coordinates on your screen.`,
		Example: `  # Click at position (500, 300)
  desktop-automation click 500 300`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Parse and validate x coordinate
			x, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid x coordinate: %s (must be an integer)", args[0])
			}
			if x < 0 {
				return fmt.Errorf("x coordinate must be non-negative, got: %d", x)
			}

			// Parse and validate y coordinate
			y, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid y coordinate: %s (must be an integer)", args[1])
			}
			if y < 0 {
				return fmt.Errorf("y coordinate must be non-negative, got: %d", y)
			}

			// Get current mouse position before clicking
			mouse := automation.NewMouse()
			currentX, currentY := mouse.GetPosition()
			fmt.Printf("Current mouse position: (%d, %d)\n", currentX, currentY)

			// Perform the click
			if err := mouse.Click(x, y); err != nil {
				return fmt.Errorf("failed to click: %w", err)
			}

			fmt.Printf("Successfully clicked at coordinates (%d, %d)\n", x, y)
			return nil
		},
	}

	return clickCmd
}
