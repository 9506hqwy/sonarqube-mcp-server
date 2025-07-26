package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectDumpExport(s *server.MCPServer) {
	tool := mcp.NewTool("project_dump_export",
		mcp.WithDescription("Triggers project dump so that the project can be imported to another SonarQube server (see api/project_dump/import, available in Enterprise Edition). Requires the 'Administer' permission."),
		mcp.WithString("key",
			mcp.Description("null"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectDumpExportHandler)
}

func projectDumpExportHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectDumpExport(request)
	return toResult(c.ApiProjectDumpExport(ctx, &params, authorizationHeader))
}

func parseProjectDumpExport(request mcp.CallToolRequest) client.ApiProjectDumpExportParams {
	params := client.ApiProjectDumpExportParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}
