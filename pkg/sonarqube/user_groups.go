package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerUserGroupsAddUser(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_add_user",
		mcp.WithDescription("Add a user to a group.<br />'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiUserGroupsAddUserParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userGroupsAddUserHandler))
}

func userGroupsAddUserHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserGroupsAddUserParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserGroupsAddUser(ctx, &params, authorizationHeader))
}

func registerUserGroupsCreate(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_create",
		mcp.WithDescription("Create a group.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiUserGroupsCreateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userGroupsCreateHandler))
}

func userGroupsCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserGroupsCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserGroupsCreate(ctx, &params, authorizationHeader))
}

func registerUserGroupsDelete(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_delete",
		mcp.WithDescription("Delete a group. The default groups cannot be deleted.<br/>'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiUserGroupsDeleteParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userGroupsDeleteHandler))
}

func userGroupsDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserGroupsDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserGroupsDelete(ctx, &params, authorizationHeader))
}

func registerUserGroupsRemoveUser(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_remove_user",
		mcp.WithDescription("Remove a user from a group.<br />'id' or 'name' must be provided.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiUserGroupsRemoveUserParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userGroupsRemoveUserHandler))
}

func userGroupsRemoveUserHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserGroupsRemoveUserParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserGroupsRemoveUser(ctx, &params, authorizationHeader))
}

func registerUserGroupsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_search",
		mcp.WithDescription("Search for user groups.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiUserGroupsSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userGroupsSearchHandler))
}

func userGroupsSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserGroupsSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserGroupsSearch(ctx, &params, authorizationHeader))
}

func registerUserGroupsUpdate(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_update",
		mcp.WithDescription("Update a group.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiUserGroupsUpdateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userGroupsUpdateHandler))
}

func userGroupsUpdateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserGroupsUpdateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserGroupsUpdate(ctx, &params, authorizationHeader))
}

func registerUserGroupsUsers(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_users",
		mcp.WithDescription("Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiUserGroupsUsersParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userGroupsUsersHandler))
}

func userGroupsUsersHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserGroupsUsersParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserGroupsUsers(ctx, &params, authorizationHeader))
}
