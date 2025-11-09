package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectsBulkDelete(s *server.MCPServer) {
	tool := mcp.NewTool("projects_bulk_delete",
		mcp.WithDescription("Delete one or several projects.<br />Only the 1'000 first items in project filters are taken into account.<br />Requires 'Administer System' permission.<br />At least one parameter is required among analyzedBefore, projects and q"),
		mcp.WithInputSchema[client.ApiProjectsBulkDeleteParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsBulkDeleteHandler))
}

func projectsBulkDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectsBulkDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectsBulkDelete(ctx, &params, authorizationHeader))
}

func registerProjectsCreate(s *server.MCPServer) {
	tool := mcp.NewTool("projects_create",
		mcp.WithDescription("Create a project.<br/>Requires 'Create Projects' permission"),
		mcp.WithInputSchema[client.ApiProjectsCreateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsCreateHandler))
}

func projectsCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectsCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectsCreate(ctx, &params, authorizationHeader))
}

func registerProjectsDelete(s *server.MCPServer) {
	tool := mcp.NewTool("projects_delete",
		mcp.WithDescription("Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project."),
		mcp.WithInputSchema[client.ApiProjectsDeleteParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsDeleteHandler))
}

func projectsDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectsDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectsDelete(ctx, &params, authorizationHeader))
}

func registerProjectsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("projects_search",
		mcp.WithDescription("Search for projects or views to administrate them.<br>Requires 'Administer System' permission"),
		mcp.WithInputSchema[client.ApiProjectsSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsSearchHandler))
}

func projectsSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectsSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectsSearch(ctx, &params, authorizationHeader))
}

func registerProjectsUpdateKey(s *server.MCPServer) {
	tool := mcp.NewTool("projects_update_key",
		mcp.WithDescription("Update a project all its sub-components keys.<br>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectsUpdateKeyParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsUpdateKeyHandler))
}

func projectsUpdateKeyHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectsUpdateKeyParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectsUpdateKey(ctx, &params, authorizationHeader))
}

func registerProjectsUpdateVisibility(s *server.MCPServer) {
	tool := mcp.NewTool("projects_update_visibility",
		mcp.WithDescription("Updates visibility of a project or view.<br>Requires 'Project administer' permission on the specified project or view"),
		mcp.WithInputSchema[client.ApiProjectsUpdateVisibilityParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsUpdateVisibilityHandler))
}

func projectsUpdateVisibilityHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectsUpdateVisibilityParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectsUpdateVisibility(ctx, &params, authorizationHeader))
}
