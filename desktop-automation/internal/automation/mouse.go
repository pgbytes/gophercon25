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

// MoveTo moves the mouse to the specified coordinates
func (m *Mouse) MoveTo(x, y int) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("invalid coordinates: x=%d, y=%d (must be non-negative)", x, y)
	}

	robotgo.MoveMouse(x, y)
	return nil
}

// Click performs a mouse click at the specified coordinates
func (m *Mouse) Click(x, y int) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("invalid coordinates: x=%d, y=%d (must be non-negative)", x, y)
	}

	// Move to position first
	if err := m.MoveTo(x, y); err != nil {
		return err
	}

	// Perform the click
	robotgo.Click()
	return nil
}

// GetPosition returns the current mouse cursor position
func (m *Mouse) GetPosition() (int, int) {
	x, y := robotgo.GetMousePos()
	return x, y
}
