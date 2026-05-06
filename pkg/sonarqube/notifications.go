package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerNotificationsAdd(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiNotificationsAddParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("notifications_add",
		mcp.WithDescription("Add a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li> <li>System administration if a login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li></ul>"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(notificationsAddHandler))
}

func notificationsAddHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiNotificationsAddParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiNotificationsAdd(ctx, &params, authorizationHeader))
}

func registerNotificationsList(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiNotificationsListParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("notifications_list",
		mcp.WithDescription("List notifications of the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided</li> <li>System administration if a login is provided</li></ul>"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(notificationsListHandler))
}

func notificationsListHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiNotificationsListParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiNotificationsList(ctx, &params, authorizationHeader))
}

func registerNotificationsRemove(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiNotificationsRemoveParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("notifications_remove",
		mcp.WithDescription("Remove a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided</li> <li>System administration if a login is provided</li></ul>"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(notificationsRemoveHandler))
}

func notificationsRemoveHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiNotificationsRemoveParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiNotificationsRemove(ctx, &params, authorizationHeader))
}
