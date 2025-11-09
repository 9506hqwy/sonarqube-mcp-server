package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectAnalysesCreateEvent(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_create_event",
		mcp.WithDescription("Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectAnalysesCreateEventParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectAnalysesCreateEventHandler))
}

func projectAnalysesCreateEventHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectAnalysesCreateEventParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectAnalysesCreateEvent(ctx, &params, authorizationHeader))
}

func registerProjectAnalysesDelete(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_delete",
		mcp.WithDescription("Delete a project analysis.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the project of the specified analysis</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectAnalysesDeleteParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectAnalysesDeleteHandler))
}

func projectAnalysesDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectAnalysesDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectAnalysesDelete(ctx, &params, authorizationHeader))
}

func registerProjectAnalysesDeleteEvent(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_delete_event",
		mcp.WithDescription("Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectAnalysesDeleteEventParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectAnalysesDeleteEventHandler))
}

func projectAnalysesDeleteEventHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectAnalysesDeleteEventParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectAnalysesDeleteEvent(ctx, &params, authorizationHeader))
}

func registerProjectAnalysesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_search",
		mcp.WithDescription("Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project. <br>For applications, it also requires 'Browse' permission on its child projects."),
		mcp.WithInputSchema[client.ApiProjectAnalysesSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectAnalysesSearchHandler))
}

func projectAnalysesSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectAnalysesSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectAnalysesSearch(ctx, &params, authorizationHeader))
}

func registerProjectAnalysesSetBaseline(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_set_baseline",
		mcp.WithDescription("Set an analysis as the baseline of the New Code Period on a project or a branch.<br/>This manually set baseline.<br/>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectAnalysesSetBaselineParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectAnalysesSetBaselineHandler))
}

func projectAnalysesSetBaselineHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectAnalysesSetBaselineParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectAnalysesSetBaseline(ctx, &params, authorizationHeader))
}

func registerProjectAnalysesUnsetBaseline(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_unset_baseline",
		mcp.WithDescription("Unset any manually-set New Code Period baseline on a project or a branch.<br/>Unsetting a manual baseline restores the use of the default new code period setting.<br/>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectAnalysesUnsetBaselineParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectAnalysesUnsetBaselineHandler))
}

func projectAnalysesUnsetBaselineHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectAnalysesUnsetBaselineParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectAnalysesUnsetBaseline(ctx, &params, authorizationHeader))
}

func registerProjectAnalysesUpdateEvent(s *server.MCPServer) {
	tool := mcp.NewTool("project_analyses_update_event",
		mcp.WithDescription("Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiProjectAnalysesUpdateEventParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectAnalysesUpdateEventHandler))
}

func projectAnalysesUpdateEventHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectAnalysesUpdateEventParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectAnalysesUpdateEvent(ctx, &params, authorizationHeader))
}
