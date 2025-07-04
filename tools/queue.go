package tools

import (
	"context"
	"fmt"
	"haos-mcp/melody"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
)

var QueuePlayTool = mcp.NewTool("melody-queue-play",
	mcp.WithDescription("Play a track from the playback queue using Melody CLI"),
	mcp.WithNumber("index",
		mcp.Required(),
		mcp.Description("Index of the track to play from queue"),
	),
)

func QueuePlayHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	index, err := request.RequireInt("index")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	scriptPath := melody.GetBasePath()
	cmd := exec.Command("python", scriptPath, "queueplay", fmt.Sprintf("%d", index))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Playback failed: %v", err)), nil
	}

	return mcp.NewToolResultText(string(output)), nil
}


var QueueTool = mcp.NewTool("melody-queue",
	mcp.WithDescription("Access the playback queue using Melody CLI. Index can be used with the queue-play tool"),
)


func QueueHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	scriptPath := melody.GetBasePath()
	cmd := exec.Command("python", scriptPath, "queue", )
	output, err := cmd.CombinedOutput()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error getting queue : %v", err)), nil
	}
	return mcp.NewToolResultText(string(output)), nil
}
