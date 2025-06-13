package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerPluginsAvailable(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_available",
		mcp.WithDescription("Get the list of all the plugins available for installation on the SonarQube instance, sorted by plugin name.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: <ul><li>COMPATIBLE: plugin is compatible with current SonarQube instance.</li><li>INCOMPATIBLE: plugin is not compatible with current SonarQube instance.</li><li>REQUIRES_SYSTEM_UPGRADE: plugin requires SonarQube to be upgraded before being installed.</li><li>DEPS_REQUIRE_SYSTEM_UPGRADE: at least one plugin on which the plugin is dependent requires SonarQube to be upgraded.</li></ul>Require 'Administer System' permission. "),
	)

	s.AddTool(tool, pluginsAvailableHandler)
}

func pluginsAvailableHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPluginsAvailable(ctx, authorizationHeader))
}

func registerPluginsCancelAll(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_cancel_all",
		mcp.WithDescription("Cancels any operation pending on any plugin (install, update or uninstall)<br/>Requires user to be authenticated with Administer System permissions "),
	)

	s.AddTool(tool, pluginsCancelAllHandler)
}

func pluginsCancelAllHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPluginsCancelAll(ctx, authorizationHeader))
}

func registerPluginsInstall(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_install",
		mcp.WithDescription("Installs the latest version of a plugin specified by its key.<br/>Plugin information is retrieved from Update Center.<br/>Fails if used on commercial editions or plugin risk consent has not been accepted.<br/>Requires user to be authenticated with Administer System permissions "),
		mcp.WithString("key",
			mcp.Description("The key identifying the plugin to install "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, pluginsInstallHandler)
}

func pluginsInstallHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePluginsInstall(request)
	return toResult(c.ApiPluginsInstall(ctx, &params, authorizationHeader))
}

func parsePluginsInstall(request mcp.CallToolRequest) client.ApiPluginsInstallParams {
	params := client.ApiPluginsInstallParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}

func registerPluginsInstalled(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_installed",
		mcp.WithDescription("Get the list of all the plugins installed on the SonarQube instance, sorted by plugin name.<br/>Requires authentication. "),
		mcp.WithString("f",
			mcp.Description("Comma-separated list of the additional fields to be returned in response. No additional field is returned by default. Possible values are:<ul><li>category - category as defined in the Update Center. A connection to the Update Center is needed</li></ul> "),
		),
	)

	s.AddTool(tool, pluginsInstalledHandler)
}

func pluginsInstalledHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePluginsInstalled(request)
	return toResult(c.ApiPluginsInstalled(ctx, &params, authorizationHeader))
}

func parsePluginsInstalled(request mcp.CallToolRequest) client.ApiPluginsInstalledParams {
	params := client.ApiPluginsInstalledParams{}

	f := request.GetString("f", "")
	if f != "" {
		params.F = &f
	}

	return params
}

func registerPluginsPending(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_pending",
		mcp.WithDescription("Get the list of plugins which will either be installed or removed at the next startup of the SonarQube instance, sorted by plugin name.<br/>Require 'Administer System' permission. "),
	)

	s.AddTool(tool, pluginsPendingHandler)
}

func pluginsPendingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPluginsPending(ctx, authorizationHeader))
}

func registerPluginsUninstall(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_uninstall",
		mcp.WithDescription("Uninstalls the plugin specified by its key.<br/>Requires user to be authenticated with Administer System permissions. "),
		mcp.WithString("key",
			mcp.Description("The key identifying the plugin to uninstall "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, pluginsUninstallHandler)
}

func pluginsUninstallHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePluginsUninstall(request)
	return toResult(c.ApiPluginsUninstall(ctx, &params, authorizationHeader))
}

func parsePluginsUninstall(request mcp.CallToolRequest) client.ApiPluginsUninstallParams {
	params := client.ApiPluginsUninstallParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}

func registerPluginsUpdate(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_update",
		mcp.WithDescription("Updates a plugin specified by its key to the latest version compatible with the SonarQube instance.<br/>Plugin information is retrieved from Update Center.<br/>Requires user to be authenticated with Administer System permissions "),
		mcp.WithString("key",
			mcp.Description("The key identifying the plugin to update "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, pluginsUpdateHandler)
}

func pluginsUpdateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePluginsUpdate(request)
	return toResult(c.ApiPluginsUpdate(ctx, &params, authorizationHeader))
}

func parsePluginsUpdate(request mcp.CallToolRequest) client.ApiPluginsUpdateParams {
	params := client.ApiPluginsUpdateParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}

func registerPluginsUpdates(s *server.MCPServer) {
	tool := mcp.NewTool("plugins_updates",
		mcp.WithDescription("Lists plugins installed on the SonarQube instance for which at least one newer version is available, sorted by plugin name.<br/>Each newer version is listed, ordered from the oldest to the newest, with its own update/compatibility status.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: [COMPATIBLE, INCOMPATIBLE, REQUIRES_UPGRADE, DEPS_REQUIRE_UPGRADE].<br/>Require 'Administer System' permission. "),
	)

	s.AddTool(tool, pluginsUpdatesHandler)
}

func pluginsUpdatesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPluginsUpdates(ctx, authorizationHeader))
}
