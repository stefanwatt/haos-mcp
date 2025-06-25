package tools

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mark3labs/mcp-go/mcp"
)

var PlayTool = mcp.NewTool("melody-play",
	mcp.WithDescription("Play music using Melody CLI"),
	mcp.WithNumber("index",
		mcp.Required(),
		mcp.Description("Index of the track to play from search results"),
	),
)

func PlayHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	index, err := request.RequireInt("index")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	homeDir, err := os.UserHomeDir()
	scriptPath := filepath.Join(homeDir, "Projects", "Melody.CLI", "melody_cli.py")
	cmd := exec.Command("python", scriptPath, "play", fmt.Sprintf("%d", index))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Playback failed: %v", err)), nil
	}

	return mcp.NewToolResultText(string(output)), nil
}
