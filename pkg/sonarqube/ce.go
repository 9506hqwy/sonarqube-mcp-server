package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerCeActivity(s *server.MCPServer) {
	tool := mcp.NewTool("ce_activity",
		mcp.WithDescription("Search for tasks.<br> Either componentId or component can be provided, but not both.<br> Requires the system administration permission, or project administration permission if componentId or component is set."),
		mcp.WithString("component",
			mcp.Description("Key of the component (project) to filter on"),
		),
		mcp.WithString("componentId",
			mcp.Description("Id of the component (project) to filter on"),
		),
		mcp.WithString("maxExecutedAt",
			mcp.Description("Maximum date of end of task processing (inclusive)"),
		),
		mcp.WithString("minSubmittedAt",
			mcp.Description("Minimum date of task submission (inclusive)"),
		),
		mcp.WithString("onlyCurrents",
			mcp.Description("Filter on the last tasks (only the most recent finished task by project)"),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 1000"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li><li>task ids that are exactly the same as the supplied string</li></ul>Must not be set together with componentId"),
		),
		mcp.WithString("status",
			mcp.Description("Comma separated list of task statuses"),
		),
		mcp.WithString("type",
			mcp.Description("Task type"),
		),
	)

	s.AddTool(tool, ceActivityHandler)
}

func ceActivityHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseCeActivity(request)
	return toResult(c.ApiCeActivity(ctx, &params, authorizationHeader))
}

func parseCeActivity(request mcp.CallToolRequest) client.ApiCeActivityParams {
	params := client.ApiCeActivityParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = &component
	}

	componentId := request.GetString("componentId", "")
	if componentId != "" {
		params.ComponentId = &componentId
	}

	maxExecutedAt := request.GetString("maxExecutedAt", "")
	if maxExecutedAt != "" {
		params.MaxExecutedAt = &maxExecutedAt
	}

	minSubmittedAt := request.GetString("minSubmittedAt", "")
	if minSubmittedAt != "" {
		params.MinSubmittedAt = &minSubmittedAt
	}

	onlyCurrents := request.GetString("onlyCurrents", "")
	if onlyCurrents != "" {
		params.OnlyCurrents = &onlyCurrents
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	status := request.GetString("status", "")
	if status != "" {
		params.Status = &status
	}

	_type := request.GetString("type", "")
	if _type != "" {
		params.Type = &_type
	}

	return params
}

func registerCeActivityStatus(s *server.MCPServer) {
	tool := mcp.NewTool("ce_activity_status",
		mcp.WithDescription("Returns CE activity related metrics.<br>Requires 'Administer System' permission or 'Administer' rights on the specified project."),
		mcp.WithString("component",
			mcp.Description("Key of the component (project) to filter on"),
		),
		mcp.WithString("componentId",
			mcp.Description("Id of the component (project) to filter on"),
		),
	)

	s.AddTool(tool, ceActivityStatusHandler)
}

func ceActivityStatusHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseCeActivityStatus(request)
	return toResult(c.ApiCeActivityStatus(ctx, &params, authorizationHeader))
}

func parseCeActivityStatus(request mcp.CallToolRequest) client.ApiCeActivityStatusParams {
	params := client.ApiCeActivityStatusParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = &component
	}

	componentId := request.GetString("componentId", "")
	if componentId != "" {
		params.ComponentId = &componentId
	}

	return params
}

func registerCeComponent(s *server.MCPServer) {
	tool := mcp.NewTool("ce_component",
		mcp.WithDescription("Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component."),
		mcp.WithString("component",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, ceComponentHandler)
}

func ceComponentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseCeComponent(request)
	return toResult(c.ApiCeComponent(ctx, &params, authorizationHeader))
}

func parseCeComponent(request mcp.CallToolRequest) client.ApiCeComponentParams {
	params := client.ApiCeComponentParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = component
	}

	return params
}

func registerCeTask(s *server.MCPServer) {
	tool := mcp.NewTool("ce_task",
		mcp.WithDescription("Give Compute Engine task details such as type, status, duration and associated component.<br />Requires 'Administer System' or 'Execute Analysis' permission.<br/>Since 6.1, field \"logs\" is deprecated and its value is always false."),
		mcp.WithString("additionalFields",
			mcp.Description("Comma-separated list of the optional fields to be returned in response."),
		),
		mcp.WithString("id",
			mcp.Description("Id of task"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, ceTaskHandler)
}

func ceTaskHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseCeTask(request)
	return toResult(c.ApiCeTask(ctx, &params, authorizationHeader))
}

func parseCeTask(request mcp.CallToolRequest) client.ApiCeTaskParams {
	params := client.ApiCeTaskParams{}

	additionalFields := request.GetString("additionalFields", "")
	if additionalFields != "" {
		params.AdditionalFields = &additionalFields
	}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = id
	}

	return params
}
