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
		mcp.WithInputSchema[client.ApiProjectDumpExportParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectDumpExportHandler))
}

func projectDumpExportHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiProjectDumpExportParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiProjectDumpExport(ctx, &params, authorizationHeader))
}
