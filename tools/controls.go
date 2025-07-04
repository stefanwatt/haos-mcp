package tools

import (
	"context"
	"fmt"
	"haos-mcp/melody"
	"os/exec"
	"slices"

	"github.com/mark3labs/mcp-go/mcp"
)

var ALLOWED_VERBS = []string{"resume", "pause", "next", "prev"}

var ControlsTool = mcp.NewTool("melody-controls",
	mcp.WithDescription("Control music using Melody CLI"),
	mcp.WithString("verb",
		mcp.Required(),
		mcp.Description("the verb to use. Possible values: resume,pause,next,prev"),
	),
)

func ControlsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	verb, err := request.RequireString("verb")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if !slices.Contains(ALLOWED_VERBS, verb) {
		return mcp.NewToolResultError(fmt.Sprintf("verb now allowd: %s", verb)), nil
	}
	scriptPath := melody.GetBasePath()
	cmd := exec.Command("python", scriptPath, verb)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Applying controls failed: %v", err)), nil
	}

	return mcp.NewToolResultText(string(output)), nil
}
