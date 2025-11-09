package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerHotspotsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("hotspots_search",
		mcp.WithDescription("Search for Security Hotpots. <br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects. <br>When issue indexation is in progress returns 503 service unavailable HTTP code."),
		mcp.WithInputSchema[client.ApiHotspotsSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(hotspotsSearchHandler))
}

func hotspotsSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiHotspotsSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiHotspotsSearch(ctx, &params, authorizationHeader))
}

func registerHotspotsShow(s *server.MCPServer) {
	tool := mcp.NewTool("hotspots_show",
		mcp.WithDescription("Provides the details of a Security Hotspot."),
		mcp.WithInputSchema[client.ApiHotspotsShowParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(hotspotsShowHandler))
}

func hotspotsShowHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiHotspotsShowParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiHotspotsShow(ctx, &params, authorizationHeader))
}
