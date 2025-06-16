// Package automation provides wrappers for the robotgo library
// to simplify desktop automation tasks
package automation

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Mouse represents mouse automation functionality
type Mouse struct{}

// NewMouse creates a new mouse automation instance
func NewMouse() *Mouse {
	return &Mouse{}
}

// Move moves the mouse instantly to the specified coordinates
func (m *Mouse) Move(x, y int) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("invalid coordinates: x=%d, y=%d (must be non-negative)", x, y)
	}

	robotgo.Move(x, y)
	return nil
}

// SmoothMove moves the mouse smoothly to the specified coordinates over the given duration
func (m *Mouse) SmoothMove(x, y int, duration float64) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("invalid coordinates: x=%d, y=%d (must be non-negative)", x, y)
	}

	if duration <= 0 {
		return fmt.Errorf("invalid duration: %f (must be positive)", duration)
	}

	robotgo.MoveSmooth(x, y, duration, 1.0)
	return nil
}

// MoveTo moves the mouse to the specified coordinates (legacy method)
func (m *Mouse) MoveTo(x, y int) error {
	return m.Move(x, y)
}

// Click performs a mouse click at the specified coordinates
func (m *Mouse) Click(x, y int) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("invalid coordinates: x=%d, y=%d (must be non-negative)", x, y)
	}

	// Move to position first
	if err := m.Move(x, y); err != nil {
		return err
	}

	// Perform the click
	robotgo.Click()
	return nil
}

// GetPosition returns the current mouse cursor position
func (m *Mouse) GetPosition() (int, int) {
	x, y := robotgo.Location()
	return x, y
}
