package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerFavoritesAdd(s *server.MCPServer) {
	tool := mcp.NewTool("favorites_add",
		mcp.WithDescription("Add a component (project, file etc.) as favorite for the authenticated user.<br>Only 100 components by qualifier can be added as favorite.<br>Requires authentication and the following permission: 'Browse' on the project of the specified component."),
		mcp.WithInputSchema[client.ApiFavoritesAddParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(favoritesAddHandler))
}

func favoritesAddHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiFavoritesAddParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiFavoritesAdd(ctx, &params, authorizationHeader))
}

func registerFavoritesRemove(s *server.MCPServer) {
	tool := mcp.NewTool("favorites_remove",
		mcp.WithDescription("Remove a component (project, portfolio, application etc.) as favorite for the authenticated user.<br>Requires authentication."),
		mcp.WithInputSchema[client.ApiFavoritesRemoveParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(favoritesRemoveHandler))
}

func favoritesRemoveHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiFavoritesRemoveParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiFavoritesRemove(ctx, &params, authorizationHeader))
}

func registerFavoritesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("favorites_search",
		mcp.WithDescription("Search for the authenticated user favorites.<br>Requires authentication."),
		mcp.WithInputSchema[client.ApiFavoritesSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(favoritesSearchHandler))
}

func favoritesSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiFavoritesSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiFavoritesSearch(ctx, &params, authorizationHeader))
}
