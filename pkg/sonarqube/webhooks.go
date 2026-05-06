package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerWebhooksCreate(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiWebhooksCreateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webhooks_create",
		mcp.WithDescription("Create a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webhooksCreateHandler))
}

func webhooksCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebhooksCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebhooksCreate(ctx, &params, authorizationHeader))
}

func registerWebhooksDelete(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiWebhooksDeleteParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webhooks_delete",
		mcp.WithDescription("Delete a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webhooksDeleteHandler))
}

func webhooksDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebhooksDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebhooksDelete(ctx, &params, authorizationHeader))
}

func registerWebhooksDeliveries(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiWebhooksDeliveriesParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webhooks_deliveries",
		mcp.WithDescription("Get the recent deliveries for a specified project or Compute Engine task.<br/>Require 'Administer' permission on the related project.<br/>Note that additional information are returned by api/webhooks/delivery."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webhooksDeliveriesHandler))
}

func webhooksDeliveriesHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebhooksDeliveriesParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebhooksDeliveries(ctx, &params, authorizationHeader))
}

func registerWebhooksDelivery(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiWebhooksDeliveryParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webhooks_delivery",
		mcp.WithDescription("Get a webhook delivery by its id.<br/>Require 'Administer System' permission.<br/>Note that additional information are returned by api/webhooks/delivery."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webhooksDeliveryHandler))
}

func webhooksDeliveryHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebhooksDeliveryParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebhooksDelivery(ctx, &params, authorizationHeader))
}

func registerWebhooksList(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiWebhooksListParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webhooks_list",
		mcp.WithDescription("Search for global webhooks or project webhooks. Webhooks are ordered by name.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webhooksListHandler))
}

func webhooksListHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebhooksListParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebhooksList(ctx, &params, authorizationHeader))
}

func registerWebhooksUpdate(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiWebhooksUpdateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webhooks_update",
		mcp.WithDescription("Update a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webhooksUpdateHandler))
}

func webhooksUpdateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebhooksUpdateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebhooksUpdate(ctx, &params, authorizationHeader))
}
