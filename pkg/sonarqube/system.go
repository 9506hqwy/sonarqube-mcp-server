package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerSystemChangeLogLevel(s *server.MCPServer) {
	tool := mcp.NewTool("system_change_log_level",
		mcp.WithDescription("Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission. "),
		mcp.WithString("level",
			mcp.Description("The new level. Be cautious: DEBUG, and even more TRACE, may have performance impacts. "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, systemChangeLogLevelHandler)
}

func systemChangeLogLevelHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSystemChangeLogLevel(request)
	return toResult(c.ApiSystemChangeLogLevel(ctx, &params, authorizationHeader))
}

func parseSystemChangeLogLevel(request mcp.CallToolRequest) client.ApiSystemChangeLogLevelParams {
	params := client.ApiSystemChangeLogLevelParams{}

	level := request.GetString("level", "")
	if level != "" {
		params.Level = level
	}

	return params
}

func registerSystemDbMigrationStatus(s *server.MCPServer) {
	tool := mcp.NewTool("system_db_migration_status",
		mcp.WithDescription("Display the database migration status of SonarQube.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul> "),
	)

	s.AddTool(tool, systemDbMigrationStatusHandler)
}

func systemDbMigrationStatusHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemDbMigrationStatus(ctx, authorizationHeader))
}

func registerSystemHealth(s *server.MCPServer) {
	tool := mcp.NewTool("system_health",
		mcp.WithDescription("Provide health status of SonarQube.<p>Although global health is calculated based on both application and search nodes, detailed information is returned only for application nodes.</p><p> <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p> "),
	)

	s.AddTool(tool, systemHealthHandler)
}

func systemHealthHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemHealth(ctx, authorizationHeader))
}

func registerSystemInfo(s *server.MCPServer) {
	tool := mcp.NewTool("system_info",
		mcp.WithDescription("Get detailed information about system configuration.<br/>Requires 'Administer' permissions. "),
	)

	s.AddTool(tool, systemInfoHandler)
}

func systemInfoHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemInfo(ctx, authorizationHeader))
}

func registerSystemLogs(s *server.MCPServer) {
	tool := mcp.NewTool("system_logs",
		mcp.WithDescription("Get system logs in plain-text format. Requires system administration permission. "),
		mcp.WithString("process",
			mcp.Description("Process to get logs from "),
		),
	)

	s.AddTool(tool, systemLogsHandler)
}

func systemLogsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSystemLogs(request)
	return toResult(c.ApiSystemLogs(ctx, &params, authorizationHeader))
}

func parseSystemLogs(request mcp.CallToolRequest) client.ApiSystemLogsParams {
	params := client.ApiSystemLogsParams{}

	process := request.GetString("process", "")
	if process != "" {
		params.Process = &process
	}

	return params
}

func registerSystemMigrateDb(s *server.MCPServer) {
	tool := mcp.NewTool("system_migrate_db",
		mcp.WithDescription("Migrate the database to match the current version of SonarQube.<br/>Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul> "),
	)

	s.AddTool(tool, systemMigrateDbHandler)
}

func systemMigrateDbHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemMigrateDb(ctx, authorizationHeader))
}

func registerSystemPing(s *server.MCPServer) {
	tool := mcp.NewTool("system_ping",
		mcp.WithDescription("Answers \"pong\" as plain-text "),
	)

	s.AddTool(tool, systemPingHandler)
}

func systemPingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemPing(ctx, authorizationHeader))
}

func registerSystemRestart(s *server.MCPServer) {
	tool := mcp.NewTool("system_restart",
		mcp.WithDescription("Restarts server. Requires 'Administer System' permission. Performs a full restart of the Web, Search and Compute Engine Servers processes. Does not reload sonar.properties. "),
	)

	s.AddTool(tool, systemRestartHandler)
}

func systemRestartHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemRestart(ctx, authorizationHeader))
}

func registerSystemStatus(s *server.MCPServer) {
	tool := mcp.NewTool("system_status",
		mcp.WithDescription("Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p> "),
	)

	s.AddTool(tool, systemStatusHandler)
}

func systemStatusHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemStatus(ctx, authorizationHeader))
}

func registerSystemUpgrades(s *server.MCPServer) {
	tool := mcp.NewTool("system_upgrades",
		mcp.WithDescription("Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response. "),
	)

	s.AddTool(tool, systemUpgradesHandler)
}

func systemUpgradesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSystemUpgrades(ctx, authorizationHeader))
}
