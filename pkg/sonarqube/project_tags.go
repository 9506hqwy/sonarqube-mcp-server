package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectTagsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("project_tags_search",
		mcp.WithDescription("Search tags"),
		mcp.WithInputSchema[client.ApiProjectTagsSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectTagsSearchHandler))
}

func projectTagsSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectTagsSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectTagsSearch(ctx, &params, authorizationHeader))
}

func registerProjectTagsSet(s *server.MCPServer) {
	tool := mcp.NewTool("project_tags_set",
		mcp.WithDescription("Set tags on a project.<br>Requires the following permission: 'Administer' rights on the specified project"),
		mcp.WithInputSchema[client.ApiProjectTagsSetParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectTagsSetHandler))
}

func projectTagsSetHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectTagsSetParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectTagsSet(ctx, &params, authorizationHeader))
}
