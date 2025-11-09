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
		mcp.WithInputSchema[client.ApiCeActivityParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ceActivityHandler))
}

func ceActivityHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiCeActivityParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiCeActivity(ctx, &params, authorizationHeader))
}

func registerCeActivityStatus(s *server.MCPServer) {
	tool := mcp.NewTool("ce_activity_status",
		mcp.WithDescription("Returns CE activity related metrics.<br>Requires 'Administer System' permission or 'Administer' rights on the specified project."),
		mcp.WithInputSchema[client.ApiCeActivityStatusParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ceActivityStatusHandler))
}

func ceActivityStatusHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiCeActivityStatusParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiCeActivityStatus(ctx, &params, authorizationHeader))
}

func registerCeComponent(s *server.MCPServer) {
	tool := mcp.NewTool("ce_component",
		mcp.WithDescription("Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component."),
		mcp.WithInputSchema[client.ApiCeComponentParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ceComponentHandler))
}

func ceComponentHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiCeComponentParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiCeComponent(ctx, &params, authorizationHeader))
}

func registerCeTask(s *server.MCPServer) {
	tool := mcp.NewTool("ce_task",
		mcp.WithDescription("Give Compute Engine task details such as type, status, duration and associated component.<br />Requires 'Administer System' or 'Execute Analysis' permission.<br/>Since 6.1, field \"logs\" is deprecated and its value is always false."),
		mcp.WithInputSchema[client.ApiCeTaskParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ceTaskHandler))
}

func ceTaskHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiCeTaskParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiCeTask(ctx, &params, authorizationHeader))
}
