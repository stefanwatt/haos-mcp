package main

import (
	"log"

	melody "haos-mcp/melody"
	tools "haos-mcp/tools"

	"github.com/mark3labs/mcp-go/server"
)

var port = "10400"

func main() {
	daemonCmd, err := melody.StartMelodyDaemon()
	if err != nil {
		log.Printf("Warning: Could not start Melody daemon: %v", err)
	} else {
		defer func() {
			if daemonCmd != nil && daemonCmd.Process != nil {
				daemonCmd.Process.Kill()
			}
		}()
	}
	mcpServer := server.NewMCPServer(
		"mcp-go",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	mcpServer.AddTool(tools.PlayTool, tools.PlayHandler)
	mcpServer.AddTool(tools.SearchTool, tools.SearchHandler)

	sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL("http://192.168.178.64:"+port))
	log.Printf("Starting SSE server on localhost:" + port)
	if err := sseServer.Start(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
