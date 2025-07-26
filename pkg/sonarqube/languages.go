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
		mcp.WithString("ps",
			mcp.Description("The size of the list to return, 0 for all languages"),
		),
		mcp.WithString("q",
			mcp.Description("A pattern to match language keys/names against"),
		),
	)

	s.AddTool(tool, languagesListHandler)
}

func languagesListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseLanguagesList(request)
	return toResult(c.ApiLanguagesList(ctx, &params, authorizationHeader))
}

func parseLanguagesList(request mcp.CallToolRequest) client.ApiLanguagesListParams {
	params := client.ApiLanguagesListParams{}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	return params
}
