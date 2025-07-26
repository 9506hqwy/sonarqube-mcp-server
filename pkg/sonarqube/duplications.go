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
		mcp.WithString("key",
			mcp.Description("File key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, duplicationsShowHandler)
}

func duplicationsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseDuplicationsShow(request)
	return toResult(c.ApiDuplicationsShow(ctx, &params, authorizationHeader))
}

func parseDuplicationsShow(request mcp.CallToolRequest) client.ApiDuplicationsShowParams {
	params := client.ApiDuplicationsShowParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}
