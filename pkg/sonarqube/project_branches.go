package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectBranchesDelete(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_delete",
		mcp.WithDescription("Delete a non-main branch of a project or application.<br/>Requires 'Administer' rights on the specified project or application. "),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectBranchesDeleteHandler)
}

func projectBranchesDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBranchesDelete(request)
	return toResult(c.ApiProjectBranchesDelete(ctx, &params, authorizationHeader))
}

func parseProjectBranchesDelete(request mcp.CallToolRequest) client.ApiProjectBranchesDeleteParams {
	params := client.ApiProjectBranchesDeleteParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerProjectBranchesList(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_list",
		mcp.WithDescription("List the branches of a project or application.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project or application. "),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectBranchesListHandler)
}

func projectBranchesListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBranchesList(request)
	return toResult(c.ApiProjectBranchesList(ctx, &params, authorizationHeader))
}

func parseProjectBranchesList(request mcp.CallToolRequest) client.ApiProjectBranchesListParams {
	params := client.ApiProjectBranchesListParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerProjectBranchesRename(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_rename",
		mcp.WithDescription("Rename the main branch of a project or application.<br/>Requires 'Administer' permission on the specified project or application. "),
		mcp.WithString("name",
			mcp.Description("New name of the main branch "),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectBranchesRenameHandler)
}

func projectBranchesRenameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBranchesRename(request)
	return toResult(c.ApiProjectBranchesRename(ctx, &params, authorizationHeader))
}

func parseProjectBranchesRename(request mcp.CallToolRequest) client.ApiProjectBranchesRenameParams {
	params := client.ApiProjectBranchesRenameParams{}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerProjectBranchesSetAutomaticDeletionProtection(s *server.MCPServer) {
	tool := mcp.NewTool("project_branches_set_automatic_deletion_protection",
		mcp.WithDescription("Protect a specific branch from automatic deletion. Protection can't be disabled for the main branch.<br/>Requires 'Administer' permission on the specified project or application. "),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
		mcp.WithString("value",
			mcp.Description("Sets whether the branch should be protected from automatic deletion when it hasn't been analyzed for a set period of time. "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectBranchesSetAutomaticDeletionProtectionHandler)
}

func projectBranchesSetAutomaticDeletionProtectionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBranchesSetAutomaticDeletionProtection(request)
	return toResult(c.ApiProjectBranchesSetAutomaticDeletionProtection(ctx, &params, authorizationHeader))
}

func parseProjectBranchesSetAutomaticDeletionProtection(request mcp.CallToolRequest) client.ApiProjectBranchesSetAutomaticDeletionProtectionParams {
	params := client.ApiProjectBranchesSetAutomaticDeletionProtectionParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	value := request.GetString("value", "")
	if value != "" {
		params.Value = value
	}

	return params
}
