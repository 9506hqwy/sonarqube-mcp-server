package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectAnalysesCreateEvent(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_create_event",
		mcp.WithDescription("Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul> "),
		mcp.WithString("analysis",
			mcp.Description("Analysis key "),
			mcp.Required(),
		),
		mcp.WithString("category",
			mcp.Description("Category "),
		),
		mcp.WithString("name",
			mcp.Description("Name "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectAnalysesCreateEventHandler)
}

func projectAnalysesCreateEventHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectAnalysesCreateEvent(request)
	return toResult(c.ApiProjectAnalysesCreateEvent(ctx, &params, authorizationHeader))
}

func parseProjectAnalysesCreateEvent(request mcp.CallToolRequest) client.ApiProjectAnalysesCreateEventParams {
	params := client.ApiProjectAnalysesCreateEventParams{}

	analysis := request.GetString("analysis", "")
	if analysis != "" {
		params.Analysis = analysis
	}

	category := request.GetString("category", "")
	if category != "" {
		params.Category = &category
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}

func registerProjectAnalysesDelete(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_delete",
		mcp.WithDescription("Delete a project analysis.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the project of the specified analysis</li></ul> "),
		mcp.WithString("analysis",
			mcp.Description("Analysis key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectAnalysesDeleteHandler)
}

func projectAnalysesDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectAnalysesDelete(request)
	return toResult(c.ApiProjectAnalysesDelete(ctx, &params, authorizationHeader))
}

func parseProjectAnalysesDelete(request mcp.CallToolRequest) client.ApiProjectAnalysesDeleteParams {
	params := client.ApiProjectAnalysesDeleteParams{}

	analysis := request.GetString("analysis", "")
	if analysis != "" {
		params.Analysis = analysis
	}

	return params
}

func registerProjectAnalysesDeleteEvent(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_delete_event",
		mcp.WithDescription("Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul> "),
		mcp.WithString("event",
			mcp.Description("Event key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectAnalysesDeleteEventHandler)
}

func projectAnalysesDeleteEventHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectAnalysesDeleteEvent(request)
	return toResult(c.ApiProjectAnalysesDeleteEvent(ctx, &params, authorizationHeader))
}

func parseProjectAnalysesDeleteEvent(request mcp.CallToolRequest) client.ApiProjectAnalysesDeleteEventParams {
	params := client.ApiProjectAnalysesDeleteEventParams{}

	event := request.GetString("event", "")
	if event != "" {
		params.Event = event
	}

	return params
}

func registerProjectAnalysesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_search",
		mcp.WithDescription("Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project. <br>For applications, it also requires 'Browse' permission on its child projects. "),
		mcp.WithString("category",
			mcp.Description("Event category. Filter analyses that have at least one event of the category specified. "),
		),
		mcp.WithString("from",
			mcp.Description("Filter analyses created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided "),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number "),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500 "),
		),
		mcp.WithString("to",
			mcp.Description("Filter analyses created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided "),
		),
	)

	s.AddTool(tool, projectAnalysesSearchHandler)
}

func projectAnalysesSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectAnalysesSearch(request)
	return toResult(c.ApiProjectAnalysesSearch(ctx, &params, authorizationHeader))
}

func parseProjectAnalysesSearch(request mcp.CallToolRequest) client.ApiProjectAnalysesSearchParams {
	params := client.ApiProjectAnalysesSearchParams{}

	category := request.GetString("category", "")
	if category != "" {
		params.Category = &category
	}

	from := request.GetString("from", "")
	if from != "" {
		params.From = &from
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	to := request.GetString("to", "")
	if to != "" {
		params.To = &to
	}

	return params
}

func registerProjectAnalysesSetBaseline(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_set_baseline",
		mcp.WithDescription("Set an analysis as the baseline of the New Code Period on a project or a branch.<br/>This manually set baseline.<br/>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul> "),
		mcp.WithString("analysis",
			mcp.Description("Analysis key "),
			mcp.Required(),
		),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectAnalysesSetBaselineHandler)
}

func projectAnalysesSetBaselineHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectAnalysesSetBaseline(request)
	return toResult(c.ApiProjectAnalysesSetBaseline(ctx, &params, authorizationHeader))
}

func parseProjectAnalysesSetBaseline(request mcp.CallToolRequest) client.ApiProjectAnalysesSetBaselineParams {
	params := client.ApiProjectAnalysesSetBaselineParams{}

	analysis := request.GetString("analysis", "")
	if analysis != "" {
		params.Analysis = analysis
	}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerProjectAnalysesUnsetBaseline(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_unset_baseline",
		mcp.WithDescription("Unset any manually-set New Code Period baseline on a project or a branch.<br/>Unsetting a manual baseline restores the use of the default new code period setting.<br/>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul> "),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectAnalysesUnsetBaselineHandler)
}

func projectAnalysesUnsetBaselineHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectAnalysesUnsetBaseline(request)
	return toResult(c.ApiProjectAnalysesUnsetBaseline(ctx, &params, authorizationHeader))
}

func parseProjectAnalysesUnsetBaseline(request mcp.CallToolRequest) client.ApiProjectAnalysesUnsetBaselineParams {
	params := client.ApiProjectAnalysesUnsetBaselineParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerProjectAnalysesUpdateEvent(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_update_event",
		mcp.WithDescription("Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul> "),
		mcp.WithString("event",
			mcp.Description("Event key "),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("New name "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectAnalysesUpdateEventHandler)
}

func projectAnalysesUpdateEventHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectAnalysesUpdateEvent(request)
	return toResult(c.ApiProjectAnalysesUpdateEvent(ctx, &params, authorizationHeader))
}

func parseProjectAnalysesUpdateEvent(request mcp.CallToolRequest) client.ApiProjectAnalysesUpdateEventParams {
	params := client.ApiProjectAnalysesUpdateEventParams{}

	event := request.GetString("event", "")
	if event != "" {
		params.Event = event
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}
