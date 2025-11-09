package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerUsersAnonymize(s *server.MCPServer) {
	tool := mcp.NewTool("users_anonymize",
		mcp.WithDescription("Anonymize a deactivated user. Requires Administer System permission"),
		mcp.WithInputSchema[client.ApiUsersAnonymizeParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersAnonymizeHandler))
}

func usersAnonymizeHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersAnonymizeParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersAnonymize(ctx, &params, authorizationHeader))
}

func registerUsersChangePassword(s *server.MCPServer) {
	tool := mcp.NewTool("users_change_password",
		mcp.WithDescription("Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password."),
		mcp.WithInputSchema[client.ApiUsersChangePasswordParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersChangePasswordHandler))
}

func usersChangePasswordHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersChangePasswordParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersChangePassword(ctx, &params, authorizationHeader))
}

func registerUsersCreate(s *server.MCPServer) {
	tool := mcp.NewTool("users_create",
		mcp.WithDescription("Create a user.<br/>If a deactivated user account exists with the given login, it will be reactivated.<br/>Requires Administer System permission"),
		mcp.WithInputSchema[client.ApiUsersCreateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersCreateHandler))
}

func usersCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersCreate(ctx, &params, authorizationHeader))
}

func registerUsersDeactivate(s *server.MCPServer) {
	tool := mcp.NewTool("users_deactivate",
		mcp.WithDescription("Deactivate a user. Requires Administer System permission"),
		mcp.WithInputSchema[client.ApiUsersDeactivateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersDeactivateHandler))
}

func usersDeactivateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersDeactivateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersDeactivate(ctx, &params, authorizationHeader))
}

func registerUsersDismissSonarlintAd(s *server.MCPServer) {
	tool := mcp.NewTool("users_dismiss_sonarlint_ad",
		mcp.WithDescription("Dismiss SonarLint advertisement. Deprecated since 9.6, replaced api/users/dismiss_notice"),
	)

	s.AddTool(tool, usersDismissSonarlintAdHandler)
}

func usersDismissSonarlintAdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersDismissSonarlintAd(ctx, authorizationHeader))
}

func registerUsersGroups(s *server.MCPServer) {
	tool := mcp.NewTool("users_groups",
		mcp.WithDescription("Lists the groups a user belongs to. <br/>Requires Administer System permission."),
		mcp.WithInputSchema[client.ApiUsersGroupsParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersGroupsHandler))
}

func usersGroupsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersGroupsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersGroups(ctx, &params, authorizationHeader))
}

func registerUsersSearch(s *server.MCPServer) {
	tool := mcp.NewTool("users_search",
		mcp.WithDescription("Get a list of users. By default, only active users are returned.<br/>The following fields are only returned when user has Administer System permission or for logged-in in user :<ul> <li>'email'</li> <li>'externalIdentity'</li> <li>'externalProvider'</li> <li>'groups'</li> <li>'lastConnectionDate'</li> <li>'tokensCount'</li></ul>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user authenticates many times in less than one hour."),
		mcp.WithInputSchema[client.ApiUsersSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersSearchHandler))
}

func usersSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersSearch(ctx, &params, authorizationHeader))
}

func registerUsersUpdate(s *server.MCPServer) {
	tool := mcp.NewTool("users_update",
		mcp.WithDescription("Update a user.<br/>Requires Administer System permission"),
		mcp.WithInputSchema[client.ApiUsersUpdateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersUpdateHandler))
}

func usersUpdateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersUpdateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersUpdate(ctx, &params, authorizationHeader))
}

func registerUsersUpdateIdentityProvider(s *server.MCPServer) {
	tool := mcp.NewTool("users_update_identity_provider",
		mcp.WithDescription("Update identity provider information. <br/>It's only possible to migrate to an installed identity provider. Be careful that as soon as this information has been updated for a user, the user will only be able to authenticate on the new identity provider. It is not possible to migrate external user to local one.<br/>Requires Administer System permission."),
		mcp.WithInputSchema[client.ApiUsersUpdateIdentityProviderParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersUpdateIdentityProviderHandler))
}

func usersUpdateIdentityProviderHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersUpdateIdentityProviderParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersUpdateIdentityProvider(ctx, &params, authorizationHeader))
}

func registerUsersUpdateLogin(s *server.MCPServer) {
	tool := mcp.NewTool("users_update_login",
		mcp.WithDescription("Update a user login. A login can be updated many times.<br/>Requires Administer System permission"),
		mcp.WithInputSchema[client.ApiUsersUpdateLoginParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersUpdateLoginHandler))
}

func usersUpdateLoginHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUsersUpdateLoginParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUsersUpdateLogin(ctx, &params, authorizationHeader))
}
