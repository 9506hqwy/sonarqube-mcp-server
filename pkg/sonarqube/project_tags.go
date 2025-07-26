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
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 100"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to tags that contain the supplied string."),
		),
	)

	s.AddTool(tool, projectTagsSearchHandler)
}

func projectTagsSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectTagsSearch(request)
	return toResult(c.ApiProjectTagsSearch(ctx, &params, authorizationHeader))
}

func parseProjectTagsSearch(request mcp.CallToolRequest) client.ApiProjectTagsSearchParams {
	params := client.ApiProjectTagsSearchParams{}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	return params
}

func registerProjectTagsSet(s *server.MCPServer) {
	tool := mcp.NewTool("project_tags_set",
		mcp.WithDescription("Set tags on a project.<br>Requires the following permission: 'Administer' rights on the specified project"),
		mcp.WithString("project",
			mcp.Description("Project key"),
			mcp.Required(),
		),
		mcp.WithString("tags",
			mcp.Description("Comma-separated list of tags"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectTagsSetHandler)
}

func projectTagsSetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectTagsSet(request)
	return toResult(c.ApiProjectTagsSet(ctx, &params, authorizationHeader))
}

func parseProjectTagsSet(request mcp.CallToolRequest) client.ApiProjectTagsSetParams {
	params := client.ApiProjectTagsSetParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	tags := request.GetString("tags", "")
	if tags != "" {
		params.Tags = tags
	}

	return params
}
