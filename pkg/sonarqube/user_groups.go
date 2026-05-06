package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerUserGroupsAddUser(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiUserGroupsAddUserParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("user_groups_add_user",
		mcp.WithDescription("Add a user to a group.<br />'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiUserGroupsCreateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("user_groups_create",
		mcp.WithDescription("Create a group.<br>Requires the following permission: 'Administer System'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiUserGroupsDeleteParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("user_groups_delete",
		mcp.WithDescription("Delete a group. The default groups cannot be deleted.<br/>'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiUserGroupsRemoveUserParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("user_groups_remove_user",
		mcp.WithDescription("Remove a user from a group.<br />'id' or 'name' must be provided.<br>Requires the following permission: 'Administer System'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiUserGroupsSearchParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("user_groups_search",
		mcp.WithDescription("Search for user groups.<br>Requires the following permission: 'Administer System'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiUserGroupsUpdateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("user_groups_update",
		mcp.WithDescription("Update a group.<br>Requires the following permission: 'Administer System'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiUserGroupsUsersParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("user_groups_users",
		mcp.WithDescription("Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
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
