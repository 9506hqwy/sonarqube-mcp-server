package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerQualitygatesCopy(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_copy",
		mcp.WithDescription("Copy a Quality Gate.<br>Either 'sourceName' or 'id' must be provided. Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesCopyParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesCopyHandler))
}

func qualitygatesCopyHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesCopyParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesCopy(ctx, &params, authorizationHeader))
}

func registerQualitygatesCreate(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_create",
		mcp.WithDescription("Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesCreateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesCreateHandler))
}

func qualitygatesCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesCreate(ctx, &params, authorizationHeader))
}

func registerQualitygatesCreateCondition(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_create_condition",
		mcp.WithDescription("Add a new condition to a quality gate.<br>Either 'gateId' or 'gateName' must be provided. Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesCreateConditionParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesCreateConditionHandler))
}

func qualitygatesCreateConditionHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesCreateConditionParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesCreateCondition(ctx, &params, authorizationHeader))
}

func registerQualitygatesDeleteCondition(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_delete_condition",
		mcp.WithDescription("Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesDeleteConditionParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesDeleteConditionHandler))
}

func qualitygatesDeleteConditionHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesDeleteConditionParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesDeleteCondition(ctx, &params, authorizationHeader))
}

func registerQualitygatesDeselect(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_deselect",
		mcp.WithDescription("Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul>"),
		mcp.WithInputSchema[client.ApiQualitygatesDeselectParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesDeselectHandler))
}

func qualitygatesDeselectHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesDeselectParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesDeselect(ctx, &params, authorizationHeader))
}

func registerQualitygatesDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_destroy",
		mcp.WithDescription("Delete a Quality Gate.<br>Either 'id' or 'name' must be specified. Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesDestroyParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesDestroyHandler))
}

func qualitygatesDestroyHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesDestroyParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesDestroy(ctx, &params, authorizationHeader))
}

func registerQualitygatesGetByProject(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_get_by_project",
		mcp.WithDescription("Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiQualitygatesGetByProjectParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesGetByProjectHandler))
}

func qualitygatesGetByProjectHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesGetByProjectParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesGetByProject(ctx, &params, authorizationHeader))
}

func registerQualitygatesList(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_list",
		mcp.WithDescription("Get a list of quality gates"),
	)

	s.AddTool(tool, qualitygatesListHandler)
}

func qualitygatesListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesList(ctx, authorizationHeader))
}

func registerQualitygatesProjectStatus(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_project_status",
		mcp.WithDescription("Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided <br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li><li>'Execute Analysis' on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiQualitygatesProjectStatusParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesProjectStatusHandler))
}

func qualitygatesProjectStatusHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesProjectStatusParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesProjectStatus(ctx, &params, authorizationHeader))
}

func registerQualitygatesRename(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_rename",
		mcp.WithDescription("Rename a Quality Gate.<br>Either 'id' or 'currentName' must be specified. Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesRenameParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesRenameHandler))
}

func qualitygatesRenameHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesRenameParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesRename(ctx, &params, authorizationHeader))
}

func registerQualitygatesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_search",
		mcp.WithDescription("Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for the current user will be returned."),
		mcp.WithInputSchema[client.ApiQualitygatesSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesSearchHandler))
}

func qualitygatesSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesSearch(ctx, &params, authorizationHeader))
}

func registerQualitygatesSelect(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_select",
		mcp.WithDescription("Associate a project to a quality gate.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Gates'</li> <li>'Administer' right on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiQualitygatesSelectParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesSelectHandler))
}

func qualitygatesSelectHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesSelectParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesSelect(ctx, &params, authorizationHeader))
}

func registerQualitygatesSetAsDefault(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_set_as_default",
		mcp.WithDescription("Set a quality gate as the default quality gate.<br>Either 'id' or 'name' must be specified. Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesSetAsDefaultParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesSetAsDefaultHandler))
}

func qualitygatesSetAsDefaultHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesSetAsDefaultParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesSetAsDefault(ctx, &params, authorizationHeader))
}

func registerQualitygatesShow(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_show",
		mcp.WithDescription("Display the details of a quality gate"),
		mcp.WithInputSchema[client.ApiQualitygatesShowParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesShowHandler))
}

func qualitygatesShowHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesShowParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesShow(ctx, &params, authorizationHeader))
}

func registerQualitygatesUpdateCondition(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_update_condition",
		mcp.WithDescription("Update a condition attached to a quality gate.<br>Requires the 'Administer Quality Gates' permission."),
		mcp.WithInputSchema[client.ApiQualitygatesUpdateConditionParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualitygatesUpdateConditionHandler))
}

func qualitygatesUpdateConditionHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualitygatesUpdateConditionParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualitygatesUpdateCondition(ctx, &params, authorizationHeader))
}
