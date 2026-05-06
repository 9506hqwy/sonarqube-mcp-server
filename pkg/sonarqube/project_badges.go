package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectBadgesMeasure(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiProjectBadgesMeasureParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("project_badges_measure",
		mcp.WithDescription("Generate badge for project's measure as an SVG.<br/>Requires 'Browse' permission on the specified project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBadgesMeasureHandler))
}

func projectBadgesMeasureHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBadgesMeasureParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBadgesMeasure(ctx, &params, authorizationHeader))
}

func registerProjectBadgesQualityGate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiProjectBadgesQualityGateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("project_badges_quality_gate",
		mcp.WithDescription("Generate badge for project's quality gate as an SVG.<br/>Requires 'Browse' permission on the specified project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBadgesQualityGateHandler))
}

func projectBadgesQualityGateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBadgesQualityGateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBadgesQualityGate(ctx, &params, authorizationHeader))
}

func registerProjectBadgesRenewToken(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiProjectBadgesRenewTokenParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("project_badges_renew_token",
		mcp.WithDescription("Creates new token replacing any existing token for project or application badge access for private projects and applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Administer' permission on the specified project or application."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBadgesRenewTokenHandler))
}

func projectBadgesRenewTokenHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBadgesRenewTokenParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBadgesRenewToken(ctx, &params, authorizationHeader))
}

func registerProjectBadgesToken(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiProjectBadgesTokenParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("project_badges_token",
		mcp.WithDescription("Retrieve a token to use for project or application badge access for private projects or applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Browse' permission on the specified project or application."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectBadgesTokenHandler))
}

func projectBadgesTokenHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectBadgesTokenParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectBadgesToken(ctx, &params, authorizationHeader))
}
