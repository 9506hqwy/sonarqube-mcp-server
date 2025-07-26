package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerNotificationsAdd(s *server.MCPServer) {
	tool := mcp.NewTool("notifications_add",
		mcp.WithDescription("Add a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li> <li>System administration if a login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li></ul>"),
		mcp.WithString("channel",
			mcp.Description("Channel through which the notification is sent. For example, notifications can be sent by email."),
		),
		mcp.WithString("login",
			mcp.Description("User login"),
		),
		mcp.WithString("project",
			mcp.Description("Project key"),
		),
		mcp.WithString("type",
			mcp.Description("Notification type. Possible values are for:<ul> <li>Global notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, SQ-MyNewIssues</li> <li>Per project notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li></ul>"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, notificationsAddHandler)
}

func notificationsAddHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNotificationsAdd(request)
	return toResult(c.ApiNotificationsAdd(ctx, &params, authorizationHeader))
}

func parseNotificationsAdd(request mcp.CallToolRequest) client.ApiNotificationsAddParams {
	params := client.ApiNotificationsAddParams{}

	channel := request.GetString("channel", "")
	if channel != "" {
		params.Channel = &channel
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	_type := request.GetString("type", "")
	if _type != "" {
		params.Type = _type
	}

	return params
}

func registerNotificationsList(s *server.MCPServer) {
	tool := mcp.NewTool("notifications_list",
		mcp.WithDescription("List notifications of the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided</li> <li>System administration if a login is provided</li></ul>"),
		mcp.WithString("login",
			mcp.Description("User login"),
		),
	)

	s.AddTool(tool, notificationsListHandler)
}

func notificationsListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNotificationsList(request)
	return toResult(c.ApiNotificationsList(ctx, &params, authorizationHeader))
}

func parseNotificationsList(request mcp.CallToolRequest) client.ApiNotificationsListParams {
	params := client.ApiNotificationsListParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	return params
}

func registerNotificationsRemove(s *server.MCPServer) {
	tool := mcp.NewTool("notifications_remove",
		mcp.WithDescription("Remove a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided</li> <li>System administration if a login is provided</li></ul>"),
		mcp.WithString("channel",
			mcp.Description("Channel through which the notification is sent. For example, notifications can be sent by email."),
		),
		mcp.WithString("login",
			mcp.Description("User login"),
		),
		mcp.WithString("project",
			mcp.Description("Project key"),
		),
		mcp.WithString("type",
			mcp.Description("Notification type. Possible values are for:<ul> <li>Global notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, SQ-MyNewIssues</li> <li>Per project notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li></ul>"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, notificationsRemoveHandler)
}

func notificationsRemoveHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNotificationsRemove(request)
	return toResult(c.ApiNotificationsRemove(ctx, &params, authorizationHeader))
}

func parseNotificationsRemove(request mcp.CallToolRequest) client.ApiNotificationsRemoveParams {
	params := client.ApiNotificationsRemoveParams{}

	channel := request.GetString("channel", "")
	if channel != "" {
		params.Channel = &channel
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	_type := request.GetString("type", "")
	if _type != "" {
		params.Type = _type
	}

	return params
}
