package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerWebservicesList(s *server.MCPServer) {
	schemaObj := jsonschema.Reflect(&client.ApiWebservicesListParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webservices_list",
		mcp.WithDescription("List web services"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
	schemaObj := jsonschema.Reflect(&client.ApiWebservicesResponseExampleParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("webservices_response_example",
		mcp.WithDescription("Display web service response example"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
