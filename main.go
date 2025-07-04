package main

import (
	"flag"
	"log"

	melody "haos-mcp/melody"
	tools "haos-mcp/tools"

	"github.com/mark3labs/mcp-go/server"
)

var (
	port = "10420"
	ip   = "192.168.178.81"
)

func main() {
	// useStdio := flag.Bool("stdio", false, "Use stdio transport instead of SSE")
	flag.Parse()
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
	mcpServer.AddTool(tools.ControlsTool, tools.ControlsHandler)
	mcpServer.AddTool(tools.QueueTool, tools.QueueHandler)
	mcpServer.AddTool(tools.QueuePlayTool, tools.QueuePlayHandler)

	// if *useStdio {
		log.Printf("Starting stdio server")
		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatalf("Stdio server error: %v", err)
		}
	// } else {
	// 	sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL("http://"+ip+":"+port))
	// 	log.Printf("Starting SSE server on localhost:" + port)
	// 	if err := sseServer.Start(":" + port); err != nil {
	// 		log.Fatalf("SSE server error: %v", err)
	// 	}
	// }
}
