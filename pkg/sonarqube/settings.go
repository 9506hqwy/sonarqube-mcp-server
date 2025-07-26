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
		mcp.WithString("component",
			mcp.Description("Component key"),
		),
	)

	s.AddTool(tool, settingsListDefinitionsHandler)
}

func settingsListDefinitionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSettingsListDefinitions(request)
	return toResult(c.ApiSettingsListDefinitions(ctx, &params, authorizationHeader))
}

func parseSettingsListDefinitions(request mcp.CallToolRequest) client.ApiSettingsListDefinitionsParams {
	params := client.ApiSettingsListDefinitionsParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = &component
	}

	return params
}

func registerSettingsReset(s *server.MCPServer) {
	tool := mcp.NewTool("settings_reset",
		mcp.WithDescription("Remove a setting value.<br>The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>"),
		mcp.WithString("component",
			mcp.Description("Component key"),
		),
		mcp.WithString("keys",
			mcp.Description("Comma-separated list of keys"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, settingsResetHandler)
}

func settingsResetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSettingsReset(request)
	return toResult(c.ApiSettingsReset(ctx, &params, authorizationHeader))
}

func parseSettingsReset(request mcp.CallToolRequest) client.ApiSettingsResetParams {
	params := client.ApiSettingsResetParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = &component
	}

	keys := request.GetString("keys", "")
	if keys != "" {
		params.Keys = keys
	}

	return params
}

func registerSettingsSet(s *server.MCPServer) {
	tool := mcp.NewTool("settings_set",
		mcp.WithDescription("Update a setting value.<br>Either 'value' or 'values' must be provided.<br> The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>"),
		mcp.WithString("component",
			mcp.Description("Component key"),
		),
		mcp.WithString("fieldValues",
			mcp.Description("Setting field values. To set several values, the parameter must be called once for each value."),
		),
		mcp.WithString("key",
			mcp.Description("Setting key"),
			mcp.Required(),
		),
		mcp.WithString("value",
			mcp.Description("Setting value. To reset a value, please use the reset web service."),
		),
		mcp.WithString("values",
			mcp.Description("Setting multi value. To set several values, the parameter must be called once for each value."),
		),
	)

	s.AddTool(tool, settingsSetHandler)
}

func settingsSetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSettingsSet(request)
	return toResult(c.ApiSettingsSet(ctx, &params, authorizationHeader))
}

func parseSettingsSet(request mcp.CallToolRequest) client.ApiSettingsSetParams {
	params := client.ApiSettingsSetParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = &component
	}

	fieldValues := request.GetString("fieldValues", "")
	if fieldValues != "" {
		params.FieldValues = &fieldValues
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	value := request.GetString("value", "")
	if value != "" {
		params.Value = &value
	}

	values := request.GetString("values", "")
	if values != "" {
		params.Values = &values
	}

	return params
}

func registerSettingsValues(s *server.MCPServer) {
	tool := mcp.NewTool("settings_values",
		mcp.WithDescription("List settings values.<br>If no value has been set for a setting, then the default value is returned.<br>The settings from conf/sonar.properties are excluded from results.<br>Requires 'Browse' or 'Execute Analysis' permission when a component is specified.<br/>Secured settings values are not returned by the endpoint.<br/>"),
		mcp.WithString("component",
			mcp.Description("Component key"),
		),
		mcp.WithString("keys",
			mcp.Description("List of setting keys"),
		),
	)

	s.AddTool(tool, settingsValuesHandler)
}

func settingsValuesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSettingsValues(request)
	return toResult(c.ApiSettingsValues(ctx, &params, authorizationHeader))
}

func parseSettingsValues(request mcp.CallToolRequest) client.ApiSettingsValuesParams {
	params := client.ApiSettingsValuesParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = &component
	}

	keys := request.GetString("keys", "")
	if keys != "" {
		params.Keys = &keys
	}

	return params
}
