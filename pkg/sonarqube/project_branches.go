package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectBranchesDelete(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_delete",
		mcp.WithDescription("Delete a non-main branch of a project or application.<br/>Requires 'Administer' rights on the specified project or application."),
		mcp.WithInputSchema[client.ApiProjectBranchesDeleteParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBranchesDeleteHandler))
}

func projectBranchesDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBranchesDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBranchesDelete(ctx, &params, authorizationHeader))
}

func registerProjectBranchesList(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_list",
		mcp.WithDescription("List the branches of a project or application.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project or application."),
		mcp.WithInputSchema[client.ApiProjectBranchesListParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBranchesListHandler))
}

func projectBranchesListHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBranchesListParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBranchesList(ctx, &params, authorizationHeader))
}

func registerProjectBranchesRename(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_rename",
		mcp.WithDescription("Rename the main branch of a project or application.<br/>Requires 'Administer' permission on the specified project or application."),
		mcp.WithInputSchema[client.ApiProjectBranchesRenameParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBranchesRenameHandler))
}

func projectBranchesRenameHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBranchesRenameParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBranchesRename(ctx, &params, authorizationHeader))
}

func registerProjectBranchesSetAutomaticDeletionProtection(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_set_automatic_deletion_protection",
		mcp.WithDescription("Protect a specific branch from automatic deletion. Protection can't be disabled for the main branch.<br/>Requires 'Administer' permission on the specified project or application."),
		mcp.WithInputSchema[client.ApiProjectBranchesSetAutomaticDeletionProtectionParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBranchesSetAutomaticDeletionProtectionHandler))
}

func projectBranchesSetAutomaticDeletionProtectionHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBranchesSetAutomaticDeletionProtectionParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBranchesSetAutomaticDeletionProtection(ctx, &params, authorizationHeader))
}
