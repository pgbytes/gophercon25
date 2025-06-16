// Package automation provides wrappers for the robotgo library
// to simplify desktop automation tasks
package automation

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

// Keyboard represents keyboard automation functionality
type Keyboard struct{}

// NewKeyboard creates a new keyboard automation instance
func NewKeyboard() *Keyboard {
	return &Keyboard{}
}

// Type simulates typing the given text
func (k *Keyboard) Type(text string) error {
	// Implementation will use robotgo
	return nil
}

// Hotkey simulates pressing a keyboard shortcut
func (k *Keyboard) Hotkey(keys ...string) error {
	// Implementation will use robotgo
	return nil
}

// TypeString types the given text
func (k *Keyboard) TypeString(text string) error {
	if text == "" {
		return fmt.Errorf("cannot type an empty string")
	}

	robotgo.TypeStr(text)
	return nil
}

// TypeStringWithDelay types the given text with a delay between keystrokes
func (k *Keyboard) TypeStringWithDelay(text string, delayMs int) error {
	if text == "" {
		return fmt.Errorf("cannot type an empty string")
	}

	if delayMs <= 0 {
		// If no delay or invalid delay, use regular typing
		return k.TypeString(text)
	}

	// Type each character with delay
	for _, char := range text {
		robotgo.TypeStr(string(char))
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}

	return nil
}
