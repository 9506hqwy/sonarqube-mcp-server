package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerUserGroupsAddUser(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_add_user",
		mcp.WithDescription("Add a user to a group.<br />'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'. "),
		mcp.WithString("id",
			mcp.Description("Group id, use 'name' instead "),
		),
		mcp.WithString("login",
			mcp.Description("User login "),
		),
		mcp.WithString("name",
			mcp.Description("Group name "),
		),
	)

	s.AddTool(tool, userGroupsAddUserHandler)
}

func userGroupsAddUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserGroupsAddUser(request)
	return toResult(c.ApiUserGroupsAddUser(ctx, &params, authorizationHeader))
}

func parseUserGroupsAddUser(request mcp.CallToolRequest) client.ApiUserGroupsAddUserParams {
	params := client.ApiUserGroupsAddUserParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	return params
}

func registerUserGroupsCreate(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_create",
		mcp.WithDescription("Create a group.<br>Requires the following permission: 'Administer System'. "),
		mcp.WithString("description",
			mcp.Description("Description for the new group. A group description cannot be larger than 200 characters. "),
		),
		mcp.WithString("name",
			mcp.Description("Name for the new group. A group name cannot be larger than 255 characters and must be unique. The value 'anyone' (whatever the case) is reserved and cannot be used. "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, userGroupsCreateHandler)
}

func userGroupsCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserGroupsCreate(request)
	return toResult(c.ApiUserGroupsCreate(ctx, &params, authorizationHeader))
}

func parseUserGroupsCreate(request mcp.CallToolRequest) client.ApiUserGroupsCreateParams {
	params := client.ApiUserGroupsCreateParams{}

	description := request.GetString("description", "")
	if description != "" {
		params.Description = &description
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}

func registerUserGroupsDelete(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_delete",
		mcp.WithDescription("Delete a group. The default groups cannot be deleted.<br/>'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'. "),
		mcp.WithString("id",
			mcp.Description("Group id, use 'name' instead "),
		),
		mcp.WithString("name",
			mcp.Description("Group name "),
		),
	)

	s.AddTool(tool, userGroupsDeleteHandler)
}

func userGroupsDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserGroupsDelete(request)
	return toResult(c.ApiUserGroupsDelete(ctx, &params, authorizationHeader))
}

func parseUserGroupsDelete(request mcp.CallToolRequest) client.ApiUserGroupsDeleteParams {
	params := client.ApiUserGroupsDeleteParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	return params
}

func registerUserGroupsRemoveUser(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_remove_user",
		mcp.WithDescription("Remove a user from a group.<br />'id' or 'name' must be provided.<br>Requires the following permission: 'Administer System'. "),
		mcp.WithString("id",
			mcp.Description("Group id, use 'name' instead "),
		),
		mcp.WithString("login",
			mcp.Description("User login "),
		),
		mcp.WithString("name",
			mcp.Description("Group name "),
		),
	)

	s.AddTool(tool, userGroupsRemoveUserHandler)
}

func userGroupsRemoveUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserGroupsRemoveUser(request)
	return toResult(c.ApiUserGroupsRemoveUser(ctx, &params, authorizationHeader))
}

func parseUserGroupsRemoveUser(request mcp.CallToolRequest) client.ApiUserGroupsRemoveUserParams {
	params := client.ApiUserGroupsRemoveUserParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	return params
}

func registerUserGroupsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_search",
		mcp.WithDescription("Search for user groups.<br>Requires the following permission: 'Administer System'. "),
		mcp.WithString("f",
			mcp.Description("Comma-separated list of the fields to be returned in response. All the fields are returned by default. "),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number "),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500 "),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to names that contain the supplied string. "),
		),
	)

	s.AddTool(tool, userGroupsSearchHandler)
}

func userGroupsSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserGroupsSearch(request)
	return toResult(c.ApiUserGroupsSearch(ctx, &params, authorizationHeader))
}

func parseUserGroupsSearch(request mcp.CallToolRequest) client.ApiUserGroupsSearchParams {
	params := client.ApiUserGroupsSearchParams{}

	f := request.GetString("f", "")
	if f != "" {
		params.F = &f
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

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

func registerUserGroupsUpdate(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_update",
		mcp.WithDescription("Update a group.<br>Requires the following permission: 'Administer System'. "),
		mcp.WithString("currentName",
			mcp.Description("Name of the group to be updated. Mandatory unless 'id' is used. "),
		),
		mcp.WithString("description",
			mcp.Description("New optional description for the group. A group description cannot be larger than 200 characters. If value is not defined, then description is not changed. "),
		),
		mcp.WithString("id",
			mcp.Description("Identifier of the group. Use 'currentName' instead. "),
		),
		mcp.WithString("name",
			mcp.Description("New optional name for the group. A group name cannot be larger than 255 characters and must be unique. Value 'anyone' (whatever the case) is reserved and cannot be used. If value is empty or not defined, then name is not changed. "),
		),
	)

	s.AddTool(tool, userGroupsUpdateHandler)
}

func userGroupsUpdateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserGroupsUpdate(request)
	return toResult(c.ApiUserGroupsUpdate(ctx, &params, authorizationHeader))
}

func parseUserGroupsUpdate(request mcp.CallToolRequest) client.ApiUserGroupsUpdateParams {
	params := client.ApiUserGroupsUpdateParams{}

	currentName := request.GetString("currentName", "")
	if currentName != "" {
		params.CurrentName = &currentName
	}

	description := request.GetString("description", "")
	if description != "" {
		params.Description = &description
	}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	return params
}

func registerUserGroupsUsers(s *server.MCPServer) {
	tool := mcp.NewTool("user_groups_users",
		mcp.WithDescription("Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'. "),
		mcp.WithString("id",
			mcp.Description("Group id, use 'name' instead "),
		),
		mcp.WithString("name",
			mcp.Description("Group name "),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number "),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0. "),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to names or logins that contain the supplied string. "),
		),
		mcp.WithString("selected",
			mcp.Description("Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all). "),
		),
	)

	s.AddTool(tool, userGroupsUsersHandler)
}

func userGroupsUsersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserGroupsUsers(request)
	return toResult(c.ApiUserGroupsUsers(ctx, &params, authorizationHeader))
}

func parseUserGroupsUsers(request mcp.CallToolRequest) client.ApiUserGroupsUsersParams {
	params := client.ApiUserGroupsUsersParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = &id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	selected := request.GetString("selected", "")
	if selected != "" {
		params.Selected = &selected
	}

	return params
}
