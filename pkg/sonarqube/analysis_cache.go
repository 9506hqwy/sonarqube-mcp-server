package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerAnalysisCacheGet(s *server.MCPServer) {
	tool := mcp.NewTool("analysis_cache_get",
		mcp.WithDescription("Get the scanner's cached data for a branch. Requires scan permission on the project. Data is returned gzipped if the corresponding 'Accept-Encoding' header is set in the request. "),
		mcp.WithString("branch",
			mcp.Description("Branch key. If not provided, main branch will be used. "),
		),
		mcp.WithString("project",
			mcp.Description("Project key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, analysisCacheGetHandler)
}

func analysisCacheGetHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAnalysisCacheGet(request)
	return toResult(c.ApiAnalysisCacheGet(ctx, &params, authorizationHeader))
}

func parseAnalysisCacheGet(request mcp.CallToolRequest) client.ApiAnalysisCacheGetParams {
	params := client.ApiAnalysisCacheGetParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}
