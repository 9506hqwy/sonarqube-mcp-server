package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerFavoritesAdd(s *server.MCPServer) {
	tool := mcp.NewTool("favorites_add",
		mcp.WithDescription("Add a component (project, file etc.) as favorite for the authenticated user.<br>Only 100 components by qualifier can be added as favorite.<br>Requires authentication and the following permission: 'Browse' on the project of the specified component. "),
		mcp.WithString("component",
			mcp.Description("Component key. Only components with qualifiers TRK, VW, SVW, APP are supported "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, favoritesAddHandler)
}

func favoritesAddHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseFavoritesAdd(request)
	return toResult(c.ApiFavoritesAdd(ctx, &params, authorizationHeader))
}

func parseFavoritesAdd(request mcp.CallToolRequest) client.ApiFavoritesAddParams {
	params := client.ApiFavoritesAddParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = component
	}

	return params
}

func registerFavoritesRemove(s *server.MCPServer) {
	tool := mcp.NewTool("favorites_remove",
		mcp.WithDescription("Remove a component (project, portfolio, application etc.) as favorite for the authenticated user.<br>Requires authentication. "),
		mcp.WithString("component",
			mcp.Description("Component key "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, favoritesRemoveHandler)
}

func favoritesRemoveHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseFavoritesRemove(request)
	return toResult(c.ApiFavoritesRemove(ctx, &params, authorizationHeader))
}

func parseFavoritesRemove(request mcp.CallToolRequest) client.ApiFavoritesRemoveParams {
	params := client.ApiFavoritesRemoveParams{}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = component
	}

	return params
}

func registerFavoritesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("favorites_search",
		mcp.WithDescription("Search for the authenticated user favorites.<br>Requires authentication. "),
		mcp.WithString("p",
			mcp.Description("1-based page number "),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500 "),
		),
	)

	s.AddTool(tool, favoritesSearchHandler)
}

func favoritesSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseFavoritesSearch(request)
	return toResult(c.ApiFavoritesSearch(ctx, &params, authorizationHeader))
}

func parseFavoritesSearch(request mcp.CallToolRequest) client.ApiFavoritesSearchParams {
	params := client.ApiFavoritesSearchParams{}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	return params
}
