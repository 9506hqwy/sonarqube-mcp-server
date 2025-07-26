package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerWebhooksCreate(s *server.MCPServer) {
	tool := mcp.NewTool("webhooks_create",
		mcp.WithDescription("Create a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithString("name",
			mcp.Description("Name displayed in the administration console of webhooks"),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("The key of the project that will own the webhook"),
		),
		mcp.WithString("secret",
			mcp.Description("If provided, secret will be used as the key to generate the HMAC hex (lowercase) digest value in the 'X-Sonar-Webhook-HMAC-SHA256' header"),
		),
		mcp.WithString("url",
			mcp.Description("Server endpoint that will receive the webhook payload, for example 'http://my_server/foo'. If HTTP Basic authentication is used, HTTPS is recommended to avoid man in the middle attacks. Example: 'https://myLogin:myPassword@my_server/foo'"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, webhooksCreateHandler)
}

func webhooksCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebhooksCreate(request)
	return toResult(c.ApiWebhooksCreate(ctx, &params, authorizationHeader))
}

func parseWebhooksCreate(request mcp.CallToolRequest) client.ApiWebhooksCreateParams {
	params := client.ApiWebhooksCreateParams{}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	secret := request.GetString("secret", "")
	if secret != "" {
		params.Secret = &secret
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerWebhooksDelete(s *server.MCPServer) {
	tool := mcp.NewTool("webhooks_delete",
		mcp.WithDescription("Delete a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithString("webhook",
			mcp.Description("The key of the webhook to be deleted, auto-generated value can be obtained through api/webhooks/create or api/webhooks/list"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, webhooksDeleteHandler)
}

func webhooksDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebhooksDelete(request)
	return toResult(c.ApiWebhooksDelete(ctx, &params, authorizationHeader))
}

func parseWebhooksDelete(request mcp.CallToolRequest) client.ApiWebhooksDeleteParams {
	params := client.ApiWebhooksDeleteParams{}

	webhook := request.GetString("webhook", "")
	if webhook != "" {
		params.Webhook = webhook
	}

	return params
}

func registerWebhooksDeliveries(s *server.MCPServer) {
	tool := mcp.NewTool("webhooks_deliveries",
		mcp.WithDescription("Get the recent deliveries for a specified project or Compute Engine task.<br/>Require 'Administer' permission on the related project.<br/>Note that additional information are returned by api/webhooks/delivery."),
		mcp.WithString("ceTaskId",
			mcp.Description("Id of the Compute Engine task"),
		),
		mcp.WithString("componentKey",
			mcp.Description("Key of the project"),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less than 500"),
		),
		mcp.WithString("webhook",
			mcp.Description("Key of the webhook that triggered those deliveries, auto-generated value that can be obtained through api/webhooks/create or api/webhooks/list"),
		),
	)

	s.AddTool(tool, webhooksDeliveriesHandler)
}

func webhooksDeliveriesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebhooksDeliveries(request)
	return toResult(c.ApiWebhooksDeliveries(ctx, &params, authorizationHeader))
}

func parseWebhooksDeliveries(request mcp.CallToolRequest) client.ApiWebhooksDeliveriesParams {
	params := client.ApiWebhooksDeliveriesParams{}

	ceTaskId := request.GetString("ceTaskId", "")
	if ceTaskId != "" {
		params.CeTaskId = &ceTaskId
	}

	componentKey := request.GetString("componentKey", "")
	if componentKey != "" {
		params.ComponentKey = &componentKey
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	webhook := request.GetString("webhook", "")
	if webhook != "" {
		params.Webhook = &webhook
	}

	return params
}

func registerWebhooksDelivery(s *server.MCPServer) {
	tool := mcp.NewTool("webhooks_delivery",
		mcp.WithDescription("Get a webhook delivery by its id.<br/>Require 'Administer System' permission.<br/>Note that additional information are returned by api/webhooks/delivery."),
		mcp.WithString("deliveryId",
			mcp.Description("Id of delivery"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, webhooksDeliveryHandler)
}

func webhooksDeliveryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebhooksDelivery(request)
	return toResult(c.ApiWebhooksDelivery(ctx, &params, authorizationHeader))
}

func parseWebhooksDelivery(request mcp.CallToolRequest) client.ApiWebhooksDeliveryParams {
	params := client.ApiWebhooksDeliveryParams{}

	deliveryId := request.GetString("deliveryId", "")
	if deliveryId != "" {
		params.DeliveryId = deliveryId
	}

	return params
}

func registerWebhooksList(s *server.MCPServer) {
	tool := mcp.NewTool("webhooks_list",
		mcp.WithDescription("Search for global webhooks or project webhooks. Webhooks are ordered by name.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithString("project",
			mcp.Description("Project key"),
		),
	)

	s.AddTool(tool, webhooksListHandler)
}

func webhooksListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebhooksList(request)
	return toResult(c.ApiWebhooksList(ctx, &params, authorizationHeader))
}

func parseWebhooksList(request mcp.CallToolRequest) client.ApiWebhooksListParams {
	params := client.ApiWebhooksListParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	return params
}

func registerWebhooksUpdate(s *server.MCPServer) {
	tool := mcp.NewTool("webhooks_update",
		mcp.WithDescription("Update a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission."),
		mcp.WithString("name",
			mcp.Description("new name of the webhook"),
			mcp.Required(),
		),
		mcp.WithString("secret",
			mcp.Description("If provided, secret will be used as the key to generate the HMAC hex (lowercase) digest value in the 'X-Sonar-Webhook-HMAC-SHA256' header. If blank, any secret previously configured will be removed. If not set, the secret will remain unchanged."),
		),
		mcp.WithString("url",
			mcp.Description("new url to be called by the webhook"),
			mcp.Required(),
		),
		mcp.WithString("webhook",
			mcp.Description("The key of the webhook to be updated, auto-generated value can be obtained through api/webhooks/create or api/webhooks/list"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, webhooksUpdateHandler)
}

func webhooksUpdateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebhooksUpdate(request)
	return toResult(c.ApiWebhooksUpdate(ctx, &params, authorizationHeader))
}

func parseWebhooksUpdate(request mcp.CallToolRequest) client.ApiWebhooksUpdateParams {
	params := client.ApiWebhooksUpdateParams{}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	secret := request.GetString("secret", "")
	if secret != "" {
		params.Secret = &secret
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	webhook := request.GetString("webhook", "")
	if webhook != "" {
		params.Webhook = webhook
	}

	return params
}
