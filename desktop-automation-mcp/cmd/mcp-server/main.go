package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/pgbytes/desktop-automation-mcp/internal/automation"
)

func main() {
	// Create MCP server with desktop automation capabilities
	s := server.NewMCPServer("Desktop Automation MCP", "1.0.0",
		server.WithToolCapabilities(true),
	)

	// Initialize automation components
	mouse := automation.NewMouse()
	keyboard := automation.NewKeyboard()

	// Add mouse tools
	addMouseTools(s, mouse)

	// Add keyboard tools
	addKeyboardTools(s, keyboard)

	// Start STDIO server
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// addMouseTools adds mouse automation tools to the server
func addMouseTools(s *server.MCPServer, mouse *automation.Mouse) {
	// Mouse move tool
	s.AddTool(
		mcp.NewTool("mouse_move",
			mcp.WithDescription("Move mouse cursor to specified coordinates"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("X coordinate")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Y coordinate")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			x, err := req.RequireInt("x")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid x coordinate: %v", err)), nil
			}

			y, err := req.RequireInt("y")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid y coordinate: %v", err)), nil
			}

			if err := mouse.Move(int(x), int(y)); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to move mouse: %v", err)), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Mouse moved to (%d, %d)", x, y)), nil
		},
	)

	// Mouse smooth move tool
	s.AddTool(
		mcp.NewTool("mouse_smooth_move",
			mcp.WithDescription("Move mouse cursor smoothly to specified coordinates"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("X coordinate")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Y coordinate")),
			mcp.WithNumber("duration", mcp.DefaultNumber(1.0), mcp.Description("Duration in seconds")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			x, err := req.RequireInt("x")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid x coordinate: %v", err)), nil
			}

			y, err := req.RequireInt("y")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid y coordinate: %v", err)), nil
			}

			duration := 1.0
			if d := req.GetFloat("duration", 1.0); d != 1.0 {
				duration = d
			}

			if err := mouse.SmoothMove(int(x), int(y), duration); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to smooth move mouse: %v", err)), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Mouse smoothly moved to (%d, %d) over %.1fs", x, y, duration)), nil
		},
	)

	// Mouse click tool
	s.AddTool(
		mcp.NewTool("mouse_click",
			mcp.WithDescription("Click at specified coordinates"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("X coordinate")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Y coordinate")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			x, err := req.RequireInt("x")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid x coordinate: %v", err)), nil
			}

			y, err := req.RequireInt("y")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid y coordinate: %v", err)), nil
			}

			if err := mouse.Click(int(x), int(y)); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to click mouse: %v", err)), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Clicked at (%d, %d)", x, y)), nil
		},
	)

	// Get mouse position tool
	s.AddTool(
		mcp.NewTool("mouse_get_position",
			mcp.WithDescription("Get current mouse cursor position"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			x, y := mouse.GetPosition()
			result := map[string]interface{}{
				"x": x,
				"y": y,
			}
			jsonData, err := json.Marshal(result)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal position data: %w", err)
			}
			return mcp.NewToolResultText(string(jsonData)), nil
		},
	)
}

// addKeyboardTools adds keyboard automation tools to the server
func addKeyboardTools(s *server.MCPServer, keyboard *automation.Keyboard) {
	// Type text tool
	s.AddTool(
		mcp.NewTool("keyboard_type",
			mcp.WithDescription("Type the specified text"),
			mcp.WithString("text", mcp.Required(), mcp.Description("Text to type")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			text, err := req.RequireString("text")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid text: %v", err)), nil
			}

			if err := keyboard.TypeString(text); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to type text: %v", err)), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Typed: %s", text)), nil
		},
	)

	// Type text with delay tool
	s.AddTool(
		mcp.NewTool("keyboard_type_with_delay",
			mcp.WithDescription("Type text with delay between keystrokes"),
			mcp.WithString("text", mcp.Required(), mcp.Description("Text to type")),
			mcp.WithNumber("delay_ms", mcp.DefaultNumber(100), mcp.Description("Delay between keystrokes in milliseconds")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			text, err := req.RequireString("text")
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid text: %v", err)), nil
			}

			delayMs := req.GetInt("delay_ms", 100)

			if err := keyboard.TypeStringWithDelay(text, delayMs); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to type text with delay: %v", err)), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Typed with %dms delay: %s", delayMs, text)), nil
		},
	)
}
