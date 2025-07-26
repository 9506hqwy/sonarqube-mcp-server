package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerSourcesRaw(s *server.MCPServer) {
	tool := mcp.NewTool("sources_raw",
		mcp.WithDescription("Get source code as raw text. Require 'See Source Code' permission on file"),
		mcp.WithString("key",
			mcp.Description("File key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, sourcesRawHandler)
}

func sourcesRawHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSourcesRaw(request)
	return toResult(c.ApiSourcesRaw(ctx, &params, authorizationHeader))
}

func parseSourcesRaw(request mcp.CallToolRequest) client.ApiSourcesRawParams {
	params := client.ApiSourcesRawParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}

func registerSourcesScm(s *server.MCPServer) {
	tool := mcp.NewTool("sources_scm",
		mcp.WithDescription("Get SCM information of source files. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Author of the commit</li><li>Datetime of the commit (before 5.2 it was only the Date)</li><li>Revision of the commit (added in 5.2)</li></ol>"),
		mcp.WithString("commits_by_line",
			mcp.Description("Group lines by SCM commit if value is false, else display commits for each line, even if two consecutive lines relate to the same commit."),
		),
		mcp.WithString("from",
			mcp.Description("First line to return. Starts at 1"),
		),
		mcp.WithString("key",
			mcp.Description("File key"),
			mcp.Required(),
		),
		mcp.WithString("to",
			mcp.Description("Last line to return (inclusive)"),
		),
	)

	s.AddTool(tool, sourcesScmHandler)
}

func sourcesScmHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSourcesScm(request)
	return toResult(c.ApiSourcesScm(ctx, &params, authorizationHeader))
}

func parseSourcesScm(request mcp.CallToolRequest) client.ApiSourcesScmParams {
	params := client.ApiSourcesScmParams{}

	commits_by_line := request.GetString("commits_by_line", "")
	if commits_by_line != "" {
		params.CommitsByLine = &commits_by_line
	}

	from := request.GetString("from", "")
	if from != "" {
		params.From = &from
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	to := request.GetString("to", "")
	if to != "" {
		params.To = &to
	}

	return params
}

func registerSourcesShow(s *server.MCPServer) {
	tool := mcp.NewTool("sources_show",
		mcp.WithDescription("Get source code. Requires See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Content of the line</li></ol>"),
		mcp.WithString("from",
			mcp.Description("First line to return. Starts at 1"),
		),
		mcp.WithString("key",
			mcp.Description("File key"),
			mcp.Required(),
		),
		mcp.WithString("to",
			mcp.Description("Last line to return (inclusive)"),
		),
	)

	s.AddTool(tool, sourcesShowHandler)
}

func sourcesShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSourcesShow(request)
	return toResult(c.ApiSourcesShow(ctx, &params, authorizationHeader))
}

func parseSourcesShow(request mcp.CallToolRequest) client.ApiSourcesShowParams {
	params := client.ApiSourcesShowParams{}

	from := request.GetString("from", "")
	if from != "" {
		params.From = &from
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	to := request.GetString("to", "")
	if to != "" {
		params.To = &to
	}

	return params
}
