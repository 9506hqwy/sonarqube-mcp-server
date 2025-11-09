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
		mcp.WithInputSchema[client.ApiSourcesRawParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(sourcesRawHandler))
}

func sourcesRawHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiSourcesRawParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSourcesRaw(ctx, &params, authorizationHeader))
}

func registerSourcesScm(s *server.MCPServer) {
	tool := mcp.NewTool("sources_scm",
		mcp.WithDescription("Get SCM information of source files. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Author of the commit</li><li>Datetime of the commit (before 5.2 it was only the Date)</li><li>Revision of the commit (added in 5.2)</li></ol>"),
		mcp.WithInputSchema[client.ApiSourcesScmParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(sourcesScmHandler))
}

func sourcesScmHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiSourcesScmParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSourcesScm(ctx, &params, authorizationHeader))
}

func registerSourcesShow(s *server.MCPServer) {
	tool := mcp.NewTool("sources_show",
		mcp.WithDescription("Get source code. Requires See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Content of the line</li></ol>"),
		mcp.WithInputSchema[client.ApiSourcesShowParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(sourcesShowHandler))
}

func sourcesShowHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiSourcesShowParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiSourcesShow(ctx, &params, authorizationHeader))
}
