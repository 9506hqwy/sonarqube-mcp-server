package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectBadgesMeasure(s *server.MCPServer) {
	tool := mcp.NewTool("project_badges_measure",
		mcp.WithDescription("Generate badge for project's measure as an SVG.<br/>Requires 'Browse' permission on the specified project."),
		mcp.WithString("branch",
			mcp.Description("Branch key"),
		),
		mcp.WithString("metric",
			mcp.Description("Metric key"),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("Project or application key"),
			mcp.Required(),
		),
		mcp.WithString("token",
			mcp.Description("Project badge token"),
		),
	)

	s.AddTool(tool, projectBadgesMeasureHandler)
}

func projectBadgesMeasureHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBadgesMeasure(request)
	return toResult(c.ApiProjectBadgesMeasure(ctx, &params, authorizationHeader))
}

func parseProjectBadgesMeasure(request mcp.CallToolRequest) client.ApiProjectBadgesMeasureParams {
	params := client.ApiProjectBadgesMeasureParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	metric := request.GetString("metric", "")
	if metric != "" {
		params.Metric = metric
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	token := request.GetString("token", "")
	if token != "" {
		params.Token = &token
	}

	return params
}

func registerProjectBadgesQualityGate(s *server.MCPServer) {
	tool := mcp.NewTool("project_badges_quality_gate",
		mcp.WithDescription("Generate badge for project's quality gate as an SVG.<br/>Requires 'Browse' permission on the specified project."),
		mcp.WithString("branch",
			mcp.Description("Branch key"),
		),
		mcp.WithString("project",
			mcp.Description("Project or application key"),
			mcp.Required(),
		),
		mcp.WithString("token",
			mcp.Description("Project badge token"),
		),
	)

	s.AddTool(tool, projectBadgesQualityGateHandler)
}

func projectBadgesQualityGateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBadgesQualityGate(request)
	return toResult(c.ApiProjectBadgesQualityGate(ctx, &params, authorizationHeader))
}

func parseProjectBadgesQualityGate(request mcp.CallToolRequest) client.ApiProjectBadgesQualityGateParams {
	params := client.ApiProjectBadgesQualityGateParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	token := request.GetString("token", "")
	if token != "" {
		params.Token = &token
	}

	return params
}

func registerProjectBadgesRenewToken(s *server.MCPServer) {
	tool := mcp.NewTool("project_badges_renew_token",
		mcp.WithDescription("Creates new token replacing any existing token for project or application badge access for private projects and applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Administer' permission on the specified project or application."),
		mcp.WithString("project",
			mcp.Description("Project or application key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectBadgesRenewTokenHandler)
}

func projectBadgesRenewTokenHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBadgesRenewToken(request)
	return toResult(c.ApiProjectBadgesRenewToken(ctx, &params, authorizationHeader))
}

func parseProjectBadgesRenewToken(request mcp.CallToolRequest) client.ApiProjectBadgesRenewTokenParams {
	params := client.ApiProjectBadgesRenewTokenParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerProjectBadgesToken(s *server.MCPServer) {
	tool := mcp.NewTool("project_badges_token",
		mcp.WithDescription("Retrieve a token to use for project or application badge access for private projects or applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Browse' permission on the specified project or application."),
		mcp.WithString("project",
			mcp.Description("Project or application key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectBadgesTokenHandler)
}

func projectBadgesTokenHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectBadgesToken(request)
	return toResult(c.ApiProjectBadgesToken(ctx, &params, authorizationHeader))
}

func parseProjectBadgesToken(request mcp.CallToolRequest) client.ApiProjectBadgesTokenParams {
	params := client.ApiProjectBadgesTokenParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}
