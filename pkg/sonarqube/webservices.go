package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerWebservicesList(s *server.MCPServer) {
	tool := mcp.NewTool("webservices_list",
		mcp.WithDescription("List web services"),
		mcp.WithInputSchema[client.ApiWebservicesListParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webservicesListHandler))
}

func webservicesListHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebservicesListParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebservicesList(ctx, &params, authorizationHeader))
}

func registerWebservicesResponseExample(s *server.MCPServer) {
	tool := mcp.NewTool("webservices_response_example",
		mcp.WithDescription("Display web service response example"),
		mcp.WithInputSchema[client.ApiWebservicesResponseExampleParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(webservicesResponseExampleHandler))
}

func webservicesResponseExampleHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiWebservicesResponseExampleParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiWebservicesResponseExample(ctx, &params, authorizationHeader))
}
