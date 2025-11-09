package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerNewCodePeriodsList(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_list",
		mcp.WithDescription("List the New Code Periods for all branches in a project.<br>Requires the permission to browse the project"),
		mcp.WithInputSchema[client.ApiNewCodePeriodsListParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newCodePeriodsListHandler))
}

func newCodePeriodsListHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiNewCodePeriodsListParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiNewCodePeriodsList(ctx, &params, authorizationHeader))
}

func registerNewCodePeriodsSet(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_set",
		mcp.WithDescription("Updates the setting for the New Code Period on different levels:<br><ul><li>Project key must be provided to update the value for a project</li><li>Both project and branch keys must be provided to update the value for a branch</li></ul>Requires one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights on the specified project to change the project setting</li></ul>"),
		mcp.WithInputSchema[client.ApiNewCodePeriodsSetParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newCodePeriodsSetHandler))
}

func newCodePeriodsSetHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiNewCodePeriodsSetParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiNewCodePeriodsSet(ctx, &params, authorizationHeader))
}

func registerNewCodePeriodsShow(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_show",
		mcp.WithDescription("Shows a setting for the New Code Period.<br> If the component requested doesn't exist or if no new code period is set for it, a value is inherited from the project or from the global setting.Requires one of the following permissions if a component is specified: <ul><li>'Administer' rights on the specified component</li><li>'Execute analysis' rights on the specified component</li></ul>"),
		mcp.WithInputSchema[client.ApiNewCodePeriodsShowParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newCodePeriodsShowHandler))
}

func newCodePeriodsShowHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiNewCodePeriodsShowParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiNewCodePeriodsShow(ctx, &params, authorizationHeader))
}

func registerNewCodePeriodsUnset(s *server.MCPServer) {
	tool := mcp.NewTool("new_code_periods_unset",
		mcp.WithDescription("Unset the New Code Period setting for a branch, project or global.<br>Requires one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights for a specified component</li></ul>"),
		mcp.WithInputSchema[client.ApiNewCodePeriodsUnsetParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newCodePeriodsUnsetHandler))
}

func newCodePeriodsUnsetHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiNewCodePeriodsUnsetParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiNewCodePeriodsUnset(ctx, &params, authorizationHeader))
}
