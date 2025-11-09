package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerLanguagesList(s *server.MCPServer) {
	tool := mcp.NewTool("languages_list",
		mcp.WithDescription("List supported programming languages"),
		mcp.WithInputSchema[client.ApiLanguagesListParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(languagesListHandler))
}

func languagesListHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiLanguagesListParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiLanguagesList(ctx, &params, authorizationHeader))
}
