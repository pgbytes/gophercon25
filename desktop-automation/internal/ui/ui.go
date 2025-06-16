// Package ui provides TUI components using Bubble Tea
package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// KeyMap defines keybindings
type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Quit   key.Binding
}

// DefaultKeyMap returns the default keybindings
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

// Model represents the TUI application state
type Model struct {
	list    list.Model
	help    help.Model
	spinner spinner.Model
	keys    KeyMap
	loading bool
}

// NewModel creates a new TUI model
func NewModel() Model {
	keys := DefaultKeyMap()
	helpModel := help.New()
	spinnerModel := spinner.New()
	spinnerModel.Spinner = spinner.Dot

	return Model{
		help:    helpModel,
		keys:    keys,
		spinner: spinnerModel,
		loading: false,
	}
}

// Init initializes the TUI model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles user input and events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the TUI
func (m Model) View() string {
	// Basic view implementation
	return "Desktop Automation TUI"
}
