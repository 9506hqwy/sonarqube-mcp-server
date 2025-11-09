package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerMetricsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("metrics_search",
		mcp.WithDescription("Search for metrics"),
		mcp.WithInputSchema[client.ApiMetricsSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(metricsSearchHandler))
}

func metricsSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiMetricsSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiMetricsSearch(ctx, &params, authorizationHeader))
}

func registerMetricsTypes(s *server.MCPServer) {
	tool := mcp.NewTool("metrics_types",
		mcp.WithDescription("List all available metric types."),
	)

	s.AddTool(tool, metricsTypesHandler)
}

func metricsTypesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiMetricsTypes(ctx, authorizationHeader))
}
