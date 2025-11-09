package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerSettingsListDefinitions(s *server.MCPServer) {
	tool := mcp.NewTool("settings_list_definitions",
		mcp.WithDescription("List settings definitions.<br>Requires 'Browse' permission when a component is specified<br/>To access licensed settings, authentication is required<br/>To access secured settings, one of the following permissions is required: <ul><li>'Execute Analysis'</li><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>"),
		mcp.WithInputSchema[client.ApiSettingsListDefinitionsParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(settingsListDefinitionsHandler))
}

func settingsListDefinitionsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiSettingsListDefinitionsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSettingsListDefinitions(ctx, &params, authorizationHeader))
}

func registerSettingsReset(s *server.MCPServer) {
	tool := mcp.NewTool("settings_reset",
		mcp.WithDescription("Remove a setting value.<br>The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>"),
		mcp.WithInputSchema[client.ApiSettingsResetParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(settingsResetHandler))
}

func settingsResetHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiSettingsResetParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSettingsReset(ctx, &params, authorizationHeader))
}

func registerSettingsSet(s *server.MCPServer) {
	tool := mcp.NewTool("settings_set",
		mcp.WithDescription("Update a setting value.<br>Either 'value' or 'values' must be provided.<br> The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>"),
		mcp.WithInputSchema[client.ApiSettingsSetParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(settingsSetHandler))
}

func settingsSetHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiSettingsSetParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSettingsSet(ctx, &params, authorizationHeader))
}

func registerSettingsValues(s *server.MCPServer) {
	tool := mcp.NewTool("settings_values",
		mcp.WithDescription("List settings values.<br>If no value has been set for a setting, then the default value is returned.<br>The settings from conf/sonar.properties are excluded from results.<br>Requires 'Browse' or 'Execute Analysis' permission when a component is specified.<br/>Secured settings values are not returned by the endpoint.<br/>"),
		mcp.WithInputSchema[client.ApiSettingsValuesParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(settingsValuesHandler))
}

func settingsValuesHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiSettingsValuesParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSettingsValues(ctx, &params, authorizationHeader))
}
