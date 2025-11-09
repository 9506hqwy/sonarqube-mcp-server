package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerDuplicationsShow(s *server.MCPServer) {
	tool := mcp.NewTool("duplications_show",
		mcp.WithDescription("Get duplications. Require Browse permission on file's project"),
		mcp.WithInputSchema[client.ApiDuplicationsShowParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(duplicationsShowHandler))
}

func duplicationsShowHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiDuplicationsShowParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiDuplicationsShow(ctx, &params, authorizationHeader))
}
