// Package commands implements the CLI commands for desktop automation
package commands

import (
	"fmt"
	"strconv"

	"github.com/pgbytes/gophercon25/desktop-automation/internal/automation"
	"github.com/spf13/cobra"
)

// newMoveCmd creates the move command
func newMoveCmd() *cobra.Command {
	var (
		smooth   bool
		duration float64
	)

	moveCmd := &cobra.Command{
		Use:   "move x y",
		Short: "Move the mouse cursor to specific screen coordinates",
		Long:  `Move the mouse cursor to the specified X and Y coordinates on your screen, either instantly or smoothly.`,
		Example: `  # Move instantly to position (800, 600)
  desktop-automation move 800 600
  
  # Move smoothly to position (800, 600) over 5 seconds
  desktop-automation move --smooth --duration 5.0 800 600`,
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

			// Get current mouse position before moving
			mouse := automation.NewMouse()
			currentX, currentY := mouse.GetPosition()
			fmt.Printf("Current mouse position: (%d, %d)\n", currentX, currentY)
			fmt.Printf("Target position: (%d, %d)\n", x, y)
			fmt.Println("Moving...")

			// Perform the movement based on the smooth flag
			if smooth {
				if err := mouse.SmoothMove(x, y, duration); err != nil {
					return fmt.Errorf("failed to move smoothly: %w", err)
				}
			} else {
				if err := mouse.Move(x, y); err != nil {
					return fmt.Errorf("failed to move: %w", err)
				}
			}

			// Get final position to confirm
			finalX, finalY := mouse.GetPosition()
			fmt.Printf("Final position: (%d, %d)\n", finalX, finalY)

			return nil
		},
	}

	// Add flags for smooth movement
	moveCmd.Flags().BoolVar(&smooth, "smooth", false, "Move the mouse smoothly (animated)")
	moveCmd.Flags().Float64Var(&duration, "duration", 1.0, "Duration in seconds for smooth movement (only applied with --smooth)")

	return moveCmd
}
