package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerComponentsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("components_search",
		mcp.WithDescription("Search for components"),
		mcp.WithInputSchema[client.ApiComponentsSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(componentsSearchHandler))
}

func componentsSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiComponentsSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiComponentsSearch(ctx, &params, authorizationHeader))
}

func registerComponentsShow(s *server.MCPServer) {
	tool := mcp.NewTool("components_show",
		mcp.WithDescription("Returns a component (file, directory, project, portfolio…) and its ancestors. The ancestors are ordered from the parent to the root project. Requires the following permission: 'Browse' on the project of the specified component."),
		mcp.WithInputSchema[client.ApiComponentsShowParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(componentsShowHandler))
}

func componentsShowHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiComponentsShowParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiComponentsShow(ctx, &params, authorizationHeader))
}

func registerComponentsTree(s *server.MCPServer) {
	tool := mcp.NewTool("components_tree",
		mcp.WithDescription("Navigate through components based on the chosen strategy.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned."),
		mcp.WithInputSchema[client.ApiComponentsTreeParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(componentsTreeHandler))
}

func componentsTreeHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiComponentsTreeParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiComponentsTree(ctx, &params, authorizationHeader))
}
