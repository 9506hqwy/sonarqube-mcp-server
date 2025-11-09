package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerMeasuresComponent(s *server.MCPServer) {
	tool := mcp.NewTool("measures_component",
		mcp.WithDescription("Return component with specified measures.<br>Requires the following permission: 'Browse' on the project of specified component."),
		mcp.WithInputSchema[client.ApiMeasuresComponentParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(measuresComponentHandler))
}

func measuresComponentHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiMeasuresComponentParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiMeasuresComponent(ctx, &params, authorizationHeader))
}

func registerMeasuresComponentTree(s *server.MCPServer) {
	tool := mcp.NewTool("measures_component_tree",
		mcp.WithDescription("Navigate through components based on the chosen strategy with specified measures.<br>Requires the following permission: 'Browse' on the specified project.<br>For applications, it also requires 'Browse' permission on its child projects. <br>When limiting search with the q parameter, directories are not returned."),
		mcp.WithInputSchema[client.ApiMeasuresComponentTreeParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(measuresComponentTreeHandler))
}

func measuresComponentTreeHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiMeasuresComponentTreeParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiMeasuresComponentTree(ctx, &params, authorizationHeader))
}

func registerMeasuresSearchHistory(s *server.MCPServer) {
	tool := mcp.NewTool("measures_search_history",
		mcp.WithDescription("Search measures history of a component.<br>Measures are ordered chronologically.<br>Pagination applies to the number of measures for each metric.<br>Requires the following permission: 'Browse' on the specified component. <br>For applications, it also requires 'Browse' permission on its child projects."),
		mcp.WithInputSchema[client.ApiMeasuresSearchHistoryParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(measuresSearchHistoryHandler))
}

func measuresSearchHistoryHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiMeasuresSearchHistoryParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiMeasuresSearchHistory(ctx, &params, authorizationHeader))
}
