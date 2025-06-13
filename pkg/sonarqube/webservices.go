package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerWebservicesList(s *server.MCPServer) {
	tool := mcp.NewTool("webservices_list",
		mcp.WithDescription("List web services "),
		mcp.WithString("include_internals",
			mcp.Description("Include web services that are implemented for internal use only. Their forward-compatibility is not assured "),
		),
	)

	s.AddTool(tool, webservicesListHandler)
}

func webservicesListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebservicesList(request)
	return toResult(c.ApiWebservicesList(ctx, &params, authorizationHeader))
}

func parseWebservicesList(request mcp.CallToolRequest) client.ApiWebservicesListParams {
	params := client.ApiWebservicesListParams{}

	include_internals := request.GetString("include_internals", "")
	if include_internals != "" {
		params.IncludeInternals = &include_internals
	}

	return params
}

func registerWebservicesResponseExample(s *server.MCPServer) {
	tool := mcp.NewTool("webservices_response_example",
		mcp.WithDescription("Display web service response example "),
		mcp.WithString("action",
			mcp.Description("Action of the web service "),
			mcp.Required(),
		),
		mcp.WithString("controller",
			mcp.Description("Controller of the web service "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, webservicesResponseExampleHandler)
}

func webservicesResponseExampleHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWebservicesResponseExample(request)
	return toResult(c.ApiWebservicesResponseExample(ctx, &params, authorizationHeader))
}

func parseWebservicesResponseExample(request mcp.CallToolRequest) client.ApiWebservicesResponseExampleParams {
	params := client.ApiWebservicesResponseExampleParams{}

	action := request.GetString("action", "")
	if action != "" {
		params.Action = action
	}

	controller := request.GetString("controller", "")
	if controller != "" {
		params.Controller = controller
	}

	return params
}
