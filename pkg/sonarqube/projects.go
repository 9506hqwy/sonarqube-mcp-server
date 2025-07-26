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
		mcp.WithString("analyzedBefore",
			mcp.Description("Filter the projects for which last analysis of any branch is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided."),
		),
		mcp.WithString("onProvisionedOnly",
			mcp.Description("Filter the projects that are provisioned"),
		),
		mcp.WithString("projects",
			mcp.Description("Comma-separated list of project keys"),
		),
		mcp.WithString("q",
			mcp.Description("Limit to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>"),
		),
		mcp.WithString("qualifiers",
			mcp.Description("Comma-separated list of component qualifiers. Filter the results with the specified qualifiers"),
		),
	)

	s.AddTool(tool, projectsBulkDeleteHandler)
}

func projectsBulkDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsBulkDelete(request)
	return toResult(c.ApiProjectsBulkDelete(ctx, &params, authorizationHeader))
}

func parseProjectsBulkDelete(request mcp.CallToolRequest) client.ApiProjectsBulkDeleteParams {
	params := client.ApiProjectsBulkDeleteParams{}

	analyzedBefore := request.GetString("analyzedBefore", "")
	if analyzedBefore != "" {
		params.AnalyzedBefore = &analyzedBefore
	}

	onProvisionedOnly := request.GetString("onProvisionedOnly", "")
	if onProvisionedOnly != "" {
		params.OnProvisionedOnly = &onProvisionedOnly
	}

	projects := request.GetString("projects", "")
	if projects != "" {
		params.Projects = &projects
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	qualifiers := request.GetString("qualifiers", "")
	if qualifiers != "" {
		params.Qualifiers = &qualifiers
	}

	return params
}

func registerProjectsCreate(s *server.MCPServer) {
	tool := mcp.NewTool("projects_create",
		mcp.WithDescription("Create a project.<br/>Requires 'Create Projects' permission"),
		mcp.WithString("mainBranch",
			mcp.Description("Key of the main branch of the project. If not provided, the default main branch key will be used."),
		),
		mcp.WithString("name",
			mcp.Description("Name of the project. If name is longer than 500, it is abbreviated."),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("Key of the project"),
			mcp.Required(),
		),
		mcp.WithString("visibility",
			mcp.Description("Whether the created project should be visible to everyone, or only specific user/groups.<br/>If no visibility is specified, the default project visibility will be used."),
		),
	)

	s.AddTool(tool, projectsCreateHandler)
}

func projectsCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsCreate(request)
	return toResult(c.ApiProjectsCreate(ctx, &params, authorizationHeader))
}

func parseProjectsCreate(request mcp.CallToolRequest) client.ApiProjectsCreateParams {
	params := client.ApiProjectsCreateParams{}

	mainBranch := request.GetString("mainBranch", "")
	if mainBranch != "" {
		params.MainBranch = &mainBranch
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	visibility := request.GetString("visibility", "")
	if visibility != "" {
		params.Visibility = &visibility
	}

	return params
}

func registerProjectsDelete(s *server.MCPServer) {
	tool := mcp.NewTool("projects_delete",
		mcp.WithDescription("Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project."),
		mcp.WithString("project",
			mcp.Description("Project key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectsDeleteHandler)
}

func projectsDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsDelete(request)
	return toResult(c.ApiProjectsDelete(ctx, &params, authorizationHeader))
}

func parseProjectsDelete(request mcp.CallToolRequest) client.ApiProjectsDeleteParams {
	params := client.ApiProjectsDeleteParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerProjectsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("projects_search",
		mcp.WithDescription("Search for projects or views to administrate them.<br>Requires 'Administer System' permission"),
		mcp.WithString("analyzedBefore",
			mcp.Description("Filter the projects for which the last analysis of all branches are older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided."),
		),
		mcp.WithString("onProvisionedOnly",
			mcp.Description("Filter the projects that are provisioned"),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("projects",
			mcp.Description("Comma-separated list of project keys"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>"),
		),
		mcp.WithString("qualifiers",
			mcp.Description("Comma-separated list of component qualifiers. Filter the results with the specified qualifiers"),
		),
	)

	s.AddTool(tool, projectsSearchHandler)
}

func projectsSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsSearch(request)
	return toResult(c.ApiProjectsSearch(ctx, &params, authorizationHeader))
}

func parseProjectsSearch(request mcp.CallToolRequest) client.ApiProjectsSearchParams {
	params := client.ApiProjectsSearchParams{}

	analyzedBefore := request.GetString("analyzedBefore", "")
	if analyzedBefore != "" {
		params.AnalyzedBefore = &analyzedBefore
	}

	onProvisionedOnly := request.GetString("onProvisionedOnly", "")
	if onProvisionedOnly != "" {
		params.OnProvisionedOnly = &onProvisionedOnly
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	projects := request.GetString("projects", "")
	if projects != "" {
		params.Projects = &projects
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	qualifiers := request.GetString("qualifiers", "")
	if qualifiers != "" {
		params.Qualifiers = &qualifiers
	}

	return params
}

func registerProjectsUpdateKey(s *server.MCPServer) {
	tool := mcp.NewTool("projects_update_key",
		mcp.WithDescription("Update a project all its sub-components keys.<br>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithString("from",
			mcp.Description("Project key"),
			mcp.Required(),
		),
		mcp.WithString("to",
			mcp.Description("New project key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectsUpdateKeyHandler)
}

func projectsUpdateKeyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsUpdateKey(request)
	return toResult(c.ApiProjectsUpdateKey(ctx, &params, authorizationHeader))
}

func parseProjectsUpdateKey(request mcp.CallToolRequest) client.ApiProjectsUpdateKeyParams {
	params := client.ApiProjectsUpdateKeyParams{}

	from := request.GetString("from", "")
	if from != "" {
		params.From = from
	}

	to := request.GetString("to", "")
	if to != "" {
		params.To = to
	}

	return params
}

func registerProjectsUpdateVisibility(s *server.MCPServer) {
	tool := mcp.NewTool("projects_update_visibility",
		mcp.WithDescription("Updates visibility of a project or view.<br>Requires 'Project administer' permission on the specified project or view"),
		mcp.WithString("project",
			mcp.Description("Project key"),
			mcp.Required(),
		),
		mcp.WithString("visibility",
			mcp.Description("New visibility"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectsUpdateVisibilityHandler)
}

func projectsUpdateVisibilityHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsUpdateVisibility(request)
	return toResult(c.ApiProjectsUpdateVisibility(ctx, &params, authorizationHeader))
}

func parseProjectsUpdateVisibility(request mcp.CallToolRequest) client.ApiProjectsUpdateVisibilityParams {
	params := client.ApiProjectsUpdateVisibilityParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	visibility := request.GetString("visibility", "")
	if visibility != "" {
		params.Visibility = visibility
	}

	return params
}
