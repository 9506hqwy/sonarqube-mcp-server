package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerServerVersion(s *server.MCPServer) {
	tool := mcp.NewTool("server_version",
		mcp.WithDescription("Version of SonarQube in plain text "),
	)

	s.AddTool(tool, serverVersionHandler)
}

func serverVersionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiServerVersion(ctx, authorizationHeader))
}
