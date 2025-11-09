package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectLinksCreate(s *server.MCPServer) {
	tool := mcp.NewTool("project_links_create",
		mcp.WithDescription("Create a new project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithInputSchema[client.ApiProjectLinksCreateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectLinksCreateHandler))
}

func projectLinksCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectLinksCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectLinksCreate(ctx, &params, authorizationHeader))
}

func registerProjectLinksDelete(s *server.MCPServer) {
	tool := mcp.NewTool("project_links_delete",
		mcp.WithDescription("Delete existing project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithInputSchema[client.ApiProjectLinksDeleteParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectLinksDeleteHandler))
}

func projectLinksDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectLinksDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectLinksDelete(ctx, &params, authorizationHeader))
}

func registerProjectLinksSearch(s *server.MCPServer) {
	tool := mcp.NewTool("project_links_search",
		mcp.WithDescription("List links of a project.<br>The 'projectId' or 'projectKey' must be provided.<br>Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectLinksSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectLinksSearchHandler))
}

func projectLinksSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectLinksSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectLinksSearch(ctx, &params, authorizationHeader))
}
