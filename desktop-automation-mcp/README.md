# Desktop Automation MCP Server

A Model Context Protocol (MCP) server that exposes desktop automation capabilities through mouse and keyboard control. Built using Go and the MCP-Go framework.

## Features

### Mouse Automation
- **mouse_move**: Move mouse cursor to specified coordinates instantly
- **mouse_smooth_move**: Move mouse cursor smoothly with customizable duration
- **mouse_click**: Click at specified coordinates
- **mouse_get_position**: Get current mouse cursor position

### Keyboard Automation
- **keyboard_type**: Type specified text
- **keyboard_type_with_delay**: Type text with customizable delay between keystrokes

## Installation

### Prerequisites
- Go 1.24 or later
- Task (task runner) - optional but recommended

### Build

Using Task:
```bash
task build
```

Using Go directly:
```bash
go build -o bin/desktop-automation-mcp ./cmd/mcp-server
```

## Usage

### As STDIO MCP Server

Run the server in STDIO mode (most common for MCP integration):

```bash
task run
# or
./bin/desktop-automation-mcp
```

### Integration with Claude Desktop

Add to your Claude Desktop configuration:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "desktop-automation": {
      "command": "go",
      "args": ["run", "/path/to/desktop-automation-mcp/cmd/mcp-server/main.go"],
      "env": {
        "LOG_LEVEL": "info"
      }
    }
  }
}
```

Or using the built binary:

```json
{
  "mcpServers": {
    "desktop-automation": {
      "command": "/path/to/desktop-automation-mcp/bin/desktop-automation-mcp"
    }
  }
}
```

## Development

### Available Tasks

```bash
task build    # Build the MCP server
task run      # Run the MCP server
task clean    # Clean build artifacts
task deps     # Download and tidy dependencies
task test     # Run tests
task lint     # Run linting tools
task dev      # Run in development mode with auto-restart
```

### Architecture

```
desktop-automation-mcp/
├── cmd/
│   └── mcp-server/          # Main entry point
│       └── main.go
├── internal/
│   └── automation/          # Desktop automation logic
│       ├── automation.go
│       ├── mouse.go
│       └── keyboard.go
├── bin/                     # Built binaries
├── go.mod
├── go.sum
├── Taskfile.yml
├── .gitignore
└── README.md
```

## Protocol Compliance

This server implements the Model Context Protocol (MCP) specification and uses STDIO transport for communication. It's compatible with any MCP client that supports STDIO transport.

## Security Considerations

- This server provides direct desktop automation capabilities
- Use with trusted clients only
- Consider running in restricted environments for production use
- Validate all coordinates and text inputs

## License

This project follows the same licensing as the parent GopherCon 25 project.
