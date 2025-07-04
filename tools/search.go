package tools

import (
	"context"
	"fmt"
	"haos-mcp/melody"

	"github.com/mark3labs/mcp-go/mcp"
)

var SearchTool = mcp.NewTool("melody-search",
	mcp.WithDescription("Search for music using Melody CLI"),
	mcp.WithString("query",
		mcp.Required(),
		mcp.Description("Search query for music"),
	),
)

func SearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := request.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	output,err := melody.RunCmd(query, "search")

	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Search failed: %v", err)), nil
	}

	return mcp.NewToolResultText(*output), nil
}
