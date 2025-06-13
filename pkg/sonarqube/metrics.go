package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerMetricsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("metrics_search",
		mcp.WithDescription("Search for metrics "),
		mcp.WithString("p",
			mcp.Description("1-based page number "),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500 "),
		),
	)

	s.AddTool(tool, metricsSearchHandler)
}

func metricsSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseMetricsSearch(request)
	return toResult(c.ApiMetricsSearch(ctx, &params, authorizationHeader))
}

func parseMetricsSearch(request mcp.CallToolRequest) client.ApiMetricsSearchParams {
	params := client.ApiMetricsSearchParams{}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	return params
}

func registerMetricsTypes(s *server.MCPServer) {
	tool := mcp.NewTool("metrics_types",
		mcp.WithDescription("List all available metric types. "),
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
