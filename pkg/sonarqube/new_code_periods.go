package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerNewCodePeriodsList(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_list",
		mcp.WithDescription("List the New Code Periods for all branches in a project.<br>Requires the permission to browse the project "),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, newCodePeriodsListHandler)
}

func newCodePeriodsListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNewCodePeriodsList(request)
	return toResult(c.ApiNewCodePeriodsList(ctx, &params, authorizationHeader))
}

func parseNewCodePeriodsList(request mcp.CallToolRequest) client.ApiNewCodePeriodsListParams {
	params := client.ApiNewCodePeriodsListParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerNewCodePeriodsSet(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_set",
		mcp.WithDescription("Updates the setting for the New Code Period on different levels:<br><ul><li>Project key must be provided to update the value for a project</li><li>Both project and branch keys must be provided to update the value for a branch</li></ul>Requires one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights on the specified project to change the project setting</li></ul> "),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
		),
		mcp.WithString("type",
			mcp.Description("Type<br/>New code periods of the following types are allowed:<ul><li>SPECIFIC_ANALYSIS - can be set at branch level only</li><li>PREVIOUS_VERSION - can be set at any level (global, project, branch)</li><li>NUMBER_OF_DAYS - can be set at any level (global, project, branch)</li><li>REFERENCE_BRANCH - can only be set for projects and branches</li></ul> "),
			mcp.Required(),
		),
		mcp.WithString("value",
			mcp.Description("Value<br/>For each type, a different value is expected:<ul><li>the uuid of an analysis, when type is SPECIFIC_ANALYSIS</li><li>no value, when type is PREVIOUS_VERSION</li><li>a number, when type is NUMBER_OF_DAYS</li><li>a string, when type is REFERENCE_BRANCH</li></ul> "),
		),
	)

	s.AddTool(tool, newCodePeriodsSetHandler)
}

func newCodePeriodsSetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNewCodePeriodsSet(request)
	return toResult(c.ApiNewCodePeriodsSet(ctx, &params, authorizationHeader))
}

func parseNewCodePeriodsSet(request mcp.CallToolRequest) client.ApiNewCodePeriodsSetParams {
	params := client.ApiNewCodePeriodsSetParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	_type := request.GetString("type", "")
	if _type != "" {
		params.Type = _type
	}

	value := request.GetString("value", "")
	if value != "" {
		params.Value = &value
	}

	return params
}

func registerNewCodePeriodsShow(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_show",
		mcp.WithDescription("Shows a setting for the New Code Period.<br> If the component requested doesn't exist or if no new code period is set for it, a value is inherited from the project or from the global setting.Requires one of the following permissions if a component is specified: <ul><li>'Administer' rights on the specified component</li><li>'Execute analysis' rights on the specified component</li></ul> "),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
		),
	)

	s.AddTool(tool, newCodePeriodsShowHandler)
}

func newCodePeriodsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNewCodePeriodsShow(request)
	return toResult(c.ApiNewCodePeriodsShow(ctx, &params, authorizationHeader))
}

func parseNewCodePeriodsShow(request mcp.CallToolRequest) client.ApiNewCodePeriodsShowParams {
	params := client.ApiNewCodePeriodsShowParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	return params
}

func registerNewCodePeriodsUnset(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_unset",
		mcp.WithDescription("Unset the New Code Period setting for a branch, project or global.<br>Requires one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights for a specified component</li></ul> "),
		mcp.WithString("branch",
			mcp.Description("Branch key "),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
		),
	)

	s.AddTool(tool, newCodePeriodsUnsetHandler)
}

func newCodePeriodsUnsetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNewCodePeriodsUnset(request)
	return toResult(c.ApiNewCodePeriodsUnset(ctx, &params, authorizationHeader))
}

func parseNewCodePeriodsUnset(request mcp.CallToolRequest) client.ApiNewCodePeriodsUnsetParams {
	params := client.ApiNewCodePeriodsUnsetParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	return params
}
