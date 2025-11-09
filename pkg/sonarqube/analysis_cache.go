package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerAnalysisCacheGet(s *server.MCPServer) {
	tool := mcp.NewTool("analysis_cache_get",
		mcp.WithDescription("Get the scanner's cached data for a branch. Requires scan permission on the project. Data is returned gzipped if the corresponding 'Accept-Encoding' header is set in the request."),
		mcp.WithInputSchema[client.ApiAnalysisCacheGetParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(analysisCacheGetHandler))
}

func analysisCacheGetHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAnalysisCacheGetParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAnalysisCacheGet(ctx, &params, authorizationHeader))
}
