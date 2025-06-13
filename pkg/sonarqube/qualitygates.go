package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerQualitygatesCopy(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_copy",
		mcp.WithDescription("Copy a Quality Gate.<br>Either 'sourceName' or 'id' must be provided. Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("id",
			mcp.Description("The ID of the source quality gate. This parameter is deprecated. Use 'sourceName' instead. "),
		),
		mcp.WithString("name",
			mcp.Description("The name of the quality gate to create "),
			mcp.Required(),
		),
		mcp.WithString("sourceName",
			mcp.Description("The name of the quality gate to copy "),
		),
	)

	s.AddTool(tool, qualitygatesCopyHandler)
}

func qualitygatesCopyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesCopy(request)
	return toResult(c.ApiQualitygatesCopy(ctx, &params, authorizationHeader))
}

func parseQualitygatesCopy(request mcp.CallToolRequest) client.ApiQualitygatesCopyParams {
	params := client.ApiQualitygatesCopyParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	sourceName := request.GetString("sourceName", "")
	if sourceName != "" {
		params.SourceName = &sourceName
	}

	return params
}

func registerQualitygatesCreate(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_create",
		mcp.WithDescription("Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("name",
			mcp.Description("The name of the quality gate to create "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualitygatesCreateHandler)
}

func qualitygatesCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesCreate(request)
	return toResult(c.ApiQualitygatesCreate(ctx, &params, authorizationHeader))
}

func parseQualitygatesCreate(request mcp.CallToolRequest) client.ApiQualitygatesCreateParams {
	params := client.ApiQualitygatesCreateParams{}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}

func registerQualitygatesCreateCondition(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_create_condition",
		mcp.WithDescription("Add a new condition to a quality gate.<br>Either 'gateId' or 'gateName' must be provided. Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("error",
			mcp.Description("Condition error threshold "),
			mcp.Required(),
		),
		mcp.WithString("gateId",
			mcp.Description("ID of the quality gate. This parameter is deprecated. Use 'gateName' instead. "),
		),
		mcp.WithString("gateName",
			mcp.Description("Name of the quality gate "),
		),
		mcp.WithString("metric",
			mcp.Description("Condition metric.<br/> Only metric of the following types are allowed:<ul><li>INT</li><li>MILLISEC</li><li>RATING</li><li>WORK_DUR</li><li>FLOAT</li><li>PERCENT</li><li>LEVEL</li></ul>Following metrics are forbidden:<ul><li>security_hotspots</li><li>alert_status</li><li>new_security_hotspots</li></ul> "),
			mcp.Required(),
		),
		mcp.WithString("op",
			mcp.Description("Condition operator:<br/><ul><li>LT = is lower than</li><li>GT = is greater than</li></ul> "),
		),
	)

	s.AddTool(tool, qualitygatesCreateConditionHandler)
}

func qualitygatesCreateConditionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesCreateCondition(request)
	return toResult(c.ApiQualitygatesCreateCondition(ctx, &params, authorizationHeader))
}

func parseQualitygatesCreateCondition(request mcp.CallToolRequest) client.ApiQualitygatesCreateConditionParams {
	params := client.ApiQualitygatesCreateConditionParams{}

	error := request.GetString("error", "")
	if error != "" {
		params.Error = error
	}

	gateId := request.GetString("gateId", "")
	if gateId != "" {
		params.GateId = &gateId
	}

	gateName := request.GetString("gateName", "")
	if gateName != "" {
		params.GateName = &gateName
	}

	metric := request.GetString("metric", "")
	if metric != "" {
		params.Metric = metric
	}

	op := request.GetString("op", "")
	if op != "" {
		params.Op = &op
	}

	return params
}

func registerQualitygatesDeleteCondition(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_delete_condition",
		mcp.WithDescription("Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("id",
			mcp.Description("Condition UUID "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualitygatesDeleteConditionHandler)
}

func qualitygatesDeleteConditionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesDeleteCondition(request)
	return toResult(c.ApiQualitygatesDeleteCondition(ctx, &params, authorizationHeader))
}

func parseQualitygatesDeleteCondition(request mcp.CallToolRequest) client.ApiQualitygatesDeleteConditionParams {
	params := client.ApiQualitygatesDeleteConditionParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = id
	}

	return params
}

func registerQualitygatesDeselect(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_deselect",
		mcp.WithDescription("Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul> "),
		mcp.WithString("projectKey",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualitygatesDeselectHandler)
}

func qualitygatesDeselectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesDeselect(request)
	return toResult(c.ApiQualitygatesDeselect(ctx, &params, authorizationHeader))
}

func parseQualitygatesDeselect(request mcp.CallToolRequest) client.ApiQualitygatesDeselectParams {
	params := client.ApiQualitygatesDeselectParams{}

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = projectKey
	}

	return params
}

func registerQualitygatesDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_destroy",
		mcp.WithDescription("Delete a Quality Gate.<br>Either 'id' or 'name' must be specified. Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("id",
			mcp.Description("ID of the quality gate to delete. This parameter is deprecated. Use 'name' instead. "),
		),
		mcp.WithString("name",
			mcp.Description("Name of the quality gate to delete "),
		),
	)

	s.AddTool(tool, qualitygatesDestroyHandler)
}

func qualitygatesDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesDestroy(request)
	return toResult(c.ApiQualitygatesDestroy(ctx, &params, authorizationHeader))
}

func parseQualitygatesDestroy(request mcp.CallToolRequest) client.ApiQualitygatesDestroyParams {
	params := client.ApiQualitygatesDestroyParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	return params
}

func registerQualitygatesGetByProject(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_get_by_project",
		mcp.WithDescription("Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul> "),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualitygatesGetByProjectHandler)
}

func qualitygatesGetByProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesGetByProject(request)
	return toResult(c.ApiQualitygatesGetByProject(ctx, &params, authorizationHeader))
}

func parseQualitygatesGetByProject(request mcp.CallToolRequest) client.ApiQualitygatesGetByProjectParams {
	params := client.ApiQualitygatesGetByProjectParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerQualitygatesList(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_list",
		mcp.WithDescription("Get a list of quality gates "),
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
		mcp.WithDescription("Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided <br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li><li>'Execute Analysis' on the specified project</li></ul> "),
		mcp.WithString("analysisId",
			mcp.Description("Analysis id "),
		),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
		),
		mcp.WithString("projectId",
			mcp.Description("Project UUID. Doesn't work with branches or pull requests "),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key "),
		),
		mcp.WithString("pullRequest",
			mcp.Description("Pull request id "),
		),
	)

	s.AddTool(tool, qualitygatesProjectStatusHandler)
}

func qualitygatesProjectStatusHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesProjectStatus(request)
	return toResult(c.ApiQualitygatesProjectStatus(ctx, &params, authorizationHeader))
}

func parseQualitygatesProjectStatus(request mcp.CallToolRequest) client.ApiQualitygatesProjectStatusParams {
	params := client.ApiQualitygatesProjectStatusParams{}

	analysisId := request.GetString("analysisId", "")
	if analysisId != "" {
		params.AnalysisId = &analysisId
	}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	projectId := request.GetString("projectId", "")
	if projectId != "" {
		params.ProjectId = &projectId
	}

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = &projectKey
	}

	pullRequest := request.GetString("pullRequest", "")
	if pullRequest != "" {
		params.PullRequest = &pullRequest
	}

	return params
}

func registerQualitygatesRename(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_rename",
		mcp.WithDescription("Rename a Quality Gate.<br>Either 'id' or 'currentName' must be specified. Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("currentName",
			mcp.Description("Current name of the quality gate "),
		),
		mcp.WithString("id",
			mcp.Description("ID of the quality gate to rename. This parameter is deprecated. Use 'currentName' instead. "),
		),
		mcp.WithString("name",
			mcp.Description("New name of the quality gate "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualitygatesRenameHandler)
}

func qualitygatesRenameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesRename(request)
	return toResult(c.ApiQualitygatesRename(ctx, &params, authorizationHeader))
}

func parseQualitygatesRename(request mcp.CallToolRequest) client.ApiQualitygatesRenameParams {
	params := client.ApiQualitygatesRenameParams{}

	currentName := request.GetString("currentName", "")
	if currentName != "" {
		params.CurrentName = &currentName
	}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}

func registerQualitygatesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_search",
		mcp.WithDescription("Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for the current user will be returned. "),
		mcp.WithString("gateId",
			mcp.Description("Quality Gate ID. This parameter is deprecated. Use 'gateName' instead. "),
		),
		mcp.WithString("gateName",
			mcp.Description("Quality Gate name "),
		),
		mcp.WithString("page",
			mcp.Description("Page number "),
		),
		mcp.WithString("pageSize",
			mcp.Description("Page size "),
		),
		mcp.WithString("query",
			mcp.Description("To search for projects containing this string. If this parameter is set, \"selected\" is set to \"all\". "),
		),
		mcp.WithString("selected",
			mcp.Description("Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all). "),
		),
	)

	s.AddTool(tool, qualitygatesSearchHandler)
}

func qualitygatesSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesSearch(request)
	return toResult(c.ApiQualitygatesSearch(ctx, &params, authorizationHeader))
}

func parseQualitygatesSearch(request mcp.CallToolRequest) client.ApiQualitygatesSearchParams {
	params := client.ApiQualitygatesSearchParams{}

	gateId := request.GetString("gateId", "")
	if gateId != "" {
		params.GateId = &gateId
	}

	gateName := request.GetString("gateName", "")
	if gateName != "" {
		params.GateName = &gateName
	}

	page := request.GetString("page", "")
	if page != "" {
		params.Page = &page
	}

	pageSize := request.GetString("pageSize", "")
	if pageSize != "" {
		params.PageSize = &pageSize
	}

	query := request.GetString("query", "")
	if query != "" {
		params.Query = &query
	}

	selected := request.GetString("selected", "")
	if selected != "" {
		params.Selected = &selected
	}

	return params
}

func registerQualitygatesSelect(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_select",
		mcp.WithDescription("Associate a project to a quality gate.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Gates'</li> <li>'Administer' right on the specified project</li></ul> "),
		mcp.WithString("gateId",
			mcp.Description("Quality gate ID. This parameter is deprecated. Use 'gateName' instead. "),
		),
		mcp.WithString("gateName",
			mcp.Description("Name of the quality gate "),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualitygatesSelectHandler)
}

func qualitygatesSelectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesSelect(request)
	return toResult(c.ApiQualitygatesSelect(ctx, &params, authorizationHeader))
}

func parseQualitygatesSelect(request mcp.CallToolRequest) client.ApiQualitygatesSelectParams {
	params := client.ApiQualitygatesSelectParams{}

	gateId := request.GetString("gateId", "")
	if gateId != "" {
		params.GateId = &gateId
	}

	gateName := request.GetString("gateName", "")
	if gateName != "" {
		params.GateName = &gateName
	}

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = projectKey
	}

	return params
}

func registerQualitygatesSetAsDefault(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_set_as_default",
		mcp.WithDescription("Set a quality gate as the default quality gate.<br>Either 'id' or 'name' must be specified. Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("id",
			mcp.Description("ID of the quality gate to set as default. This parameter is deprecated. Use 'name' instead. "),
		),
		mcp.WithString("name",
			mcp.Description("Name of the quality gate to set as default "),
		),
	)

	s.AddTool(tool, qualitygatesSetAsDefaultHandler)
}

func qualitygatesSetAsDefaultHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesSetAsDefault(request)
	return toResult(c.ApiQualitygatesSetAsDefault(ctx, &params, authorizationHeader))
}

func parseQualitygatesSetAsDefault(request mcp.CallToolRequest) client.ApiQualitygatesSetAsDefaultParams {
	params := client.ApiQualitygatesSetAsDefaultParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	return params
}

func registerQualitygatesShow(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_show",
		mcp.WithDescription("Display the details of a quality gate "),
		mcp.WithString("id",
			mcp.Description("ID of the quality gate. Either id or name must be set "),
		),
		mcp.WithString("name",
			mcp.Description("Name of the quality gate. Either id or name must be set "),
		),
	)

	s.AddTool(tool, qualitygatesShowHandler)
}

func qualitygatesShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesShow(request)
	return toResult(c.ApiQualitygatesShow(ctx, &params, authorizationHeader))
}

func parseQualitygatesShow(request mcp.CallToolRequest) client.ApiQualitygatesShowParams {
	params := client.ApiQualitygatesShowParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	return params
}

func registerQualitygatesUpdateCondition(s *server.MCPServer) {
	tool := mcp.NewTool("qualitygates_update_condition",
		mcp.WithDescription("Update a condition attached to a quality gate.<br>Requires the 'Administer Quality Gates' permission. "),
		mcp.WithString("error",
			mcp.Description("Condition error threshold "),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("Condition ID "),
			mcp.Required(),
		),
		mcp.WithString("metric",
			mcp.Description("Condition metric.<br/> Only metric of the following types are allowed:<ul><li>INT</li><li>MILLISEC</li><li>RATING</li><li>WORK_DUR</li><li>FLOAT</li><li>PERCENT</li><li>LEVEL</li></ul>Following metrics are forbidden:<ul><li>security_hotspots</li><li>alert_status</li><li>new_security_hotspots</li></ul> "),
			mcp.Required(),
		),
		mcp.WithString("op",
			mcp.Description("Condition operator:<br/><ul><li>LT = is lower than</li><li>GT = is greater than</li></ul> "),
		),
	)

	s.AddTool(tool, qualitygatesUpdateConditionHandler)
}

func qualitygatesUpdateConditionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualitygatesUpdateCondition(request)
	return toResult(c.ApiQualitygatesUpdateCondition(ctx, &params, authorizationHeader))
}

func parseQualitygatesUpdateCondition(request mcp.CallToolRequest) client.ApiQualitygatesUpdateConditionParams {
	params := client.ApiQualitygatesUpdateConditionParams{}

	error := request.GetString("error", "")
	if error != "" {
		params.Error = error
	}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = id
	}

	metric := request.GetString("metric", "")
	if metric != "" {
		params.Metric = metric
	}

	op := request.GetString("op", "")
	if op != "" {
		params.Op = &op
	}

	return params
}
