package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerUsersAnonymize(s *server.MCPServer) {
	tool := mcp.NewTool("users_anonymize",
		mcp.WithDescription("Anonymize a deactivated user. Requires Administer System permission "),
		mcp.WithString("login",
			mcp.Description("User login "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, usersAnonymizeHandler)
}

func usersAnonymizeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersAnonymize(request)
	return toResult(c.ApiUsersAnonymize(ctx, &params, authorizationHeader))
}

func parseUsersAnonymize(request mcp.CallToolRequest) client.ApiUsersAnonymizeParams {
	params := client.ApiUsersAnonymizeParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	return params
}

func registerUsersChangePassword(s *server.MCPServer) {
	tool := mcp.NewTool("users_change_password",
		mcp.WithDescription("Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password. "),
		mcp.WithString("login",
			mcp.Description("User login "),
			mcp.Required(),
		),
		mcp.WithString("password",
			mcp.Description("New password "),
			mcp.Required(),
		),
		mcp.WithString("previousPassword",
			mcp.Description("Previous password. Required when changing one's own password. "),
		),
	)

	s.AddTool(tool, usersChangePasswordHandler)
}

func usersChangePasswordHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersChangePassword(request)
	return toResult(c.ApiUsersChangePassword(ctx, &params, authorizationHeader))
}

func parseUsersChangePassword(request mcp.CallToolRequest) client.ApiUsersChangePasswordParams {
	params := client.ApiUsersChangePasswordParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	password := request.GetString("password", "")
	if password != "" {
		params.Password = password
	}

	previousPassword := request.GetString("previousPassword", "")
	if previousPassword != "" {
		params.PreviousPassword = &previousPassword
	}

	return params
}

func registerUsersCreate(s *server.MCPServer) {
	tool := mcp.NewTool("users_create",
		mcp.WithDescription("Create a user.<br/>If a deactivated user account exists with the given login, it will be reactivated.<br/>Requires Administer System permission "),
		mcp.WithString("email",
			mcp.Description("User email "),
		),
		mcp.WithString("local",
			mcp.Description("Specify if the user should be authenticated from SonarQube server or from an external authentication system. Password should not be set when local is set to false. "),
		),
		mcp.WithString("login",
			mcp.Description("User login "),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("User name "),
			mcp.Required(),
		),
		mcp.WithString("password",
			mcp.Description("User password. Only mandatory when creating local user, otherwise it should not be set "),
		),
		mcp.WithString("scmAccount",
			mcp.Description("List of SCM accounts. To set several values, the parameter must be called once for each value. "),
		),
	)

	s.AddTool(tool, usersCreateHandler)
}

func usersCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersCreate(request)
	return toResult(c.ApiUsersCreate(ctx, &params, authorizationHeader))
}

func parseUsersCreate(request mcp.CallToolRequest) client.ApiUsersCreateParams {
	params := client.ApiUsersCreateParams{}

	email := request.GetString("email", "")
	if email != "" {
		params.Email = &email
	}

	local := request.GetString("local", "")
	if local != "" {
		params.Local = &local
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	password := request.GetString("password", "")
	if password != "" {
		params.Password = &password
	}

	scmAccount := request.GetString("scmAccount", "")
	if scmAccount != "" {
		params.ScmAccount = &scmAccount
	}

	return params
}

func registerUsersDeactivate(s *server.MCPServer) {
	tool := mcp.NewTool("users_deactivate",
		mcp.WithDescription("Deactivate a user. Requires Administer System permission "),
		mcp.WithString("anonymize",
			mcp.Description("Anonymize user in addition to deactivating it "),
		),
		mcp.WithString("login",
			mcp.Description("User login "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, usersDeactivateHandler)
}

func usersDeactivateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersDeactivate(request)
	return toResult(c.ApiUsersDeactivate(ctx, &params, authorizationHeader))
}

func parseUsersDeactivate(request mcp.CallToolRequest) client.ApiUsersDeactivateParams {
	params := client.ApiUsersDeactivateParams{}

	anonymize := request.GetString("anonymize", "")
	if anonymize != "" {
		params.Anonymize = &anonymize
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	return params
}

func registerUsersDismissSonarlintAd(s *server.MCPServer) {
	tool := mcp.NewTool("users_dismiss_sonarlint_ad",
		mcp.WithDescription("Dismiss SonarLint advertisement. Deprecated since 9.6, replaced api/users/dismiss_notice "),
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
		mcp.WithDescription("Lists the groups a user belongs to. <br/>Requires Administer System permission. "),
		mcp.WithString("login",
			mcp.Description("A user login "),
			mcp.Required(),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number "),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0. "),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to group names that contain the supplied string. "),
		),
		mcp.WithString("selected",
			mcp.Description("Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all). "),
		),
	)

	s.AddTool(tool, usersGroupsHandler)
}

func usersGroupsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersGroups(request)
	return toResult(c.ApiUsersGroups(ctx, &params, authorizationHeader))
}

func parseUsersGroups(request mcp.CallToolRequest) client.ApiUsersGroupsParams {
	params := client.ApiUsersGroupsParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
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

func registerUsersSearch(s *server.MCPServer) {
	tool := mcp.NewTool("users_search",
		mcp.WithDescription("Get a list of users. By default, only active users are returned.<br/>The following fields are only returned when user has Administer System permission or for logged-in in user :<ul> <li>'email'</li> <li>'externalIdentity'</li> <li>'externalProvider'</li> <li>'groups'</li> <li>'lastConnectionDate'</li> <li>'tokensCount'</li></ul>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user authenticates many times in less than one hour. "),
		mcp.WithString("deactivated",
			mcp.Description("Return deactivated users instead of active users "),
		),
		mcp.WithString("externalIdentity",
			mcp.Description("Find a user by its external identity (ie. its login in the Identity Provider). This is case sensitive and only available with Administer System permission. "),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number "),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500 "),
		),
		mcp.WithString("q",
			mcp.Description("Filter on login, name and email.<br />This parameter can either be case sensitive and perform an exact match, or case insensitive and perform a partial match (contains), depending on the scenario:<br /><ul> <li> If the search query is <em>less or equal to 15 characters</em>, then the query is <em>case insensitive</em>, and will match any login, name, or email, that <em>contains</em> the search query. </li> <li> If the search query is <em>greater than 15 characters</em>, then the query becomes <em>case sensitive</em>, and will match any login, name, or email, that <em>exactly matches</em> the search query. </li></ul> "),
		),
	)

	s.AddTool(tool, usersSearchHandler)
}

func usersSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersSearch(request)
	return toResult(c.ApiUsersSearch(ctx, &params, authorizationHeader))
}

func parseUsersSearch(request mcp.CallToolRequest) client.ApiUsersSearchParams {
	params := client.ApiUsersSearchParams{}

	deactivated := request.GetString("deactivated", "")
	if deactivated != "" {
		params.Deactivated = &deactivated
	}

	externalIdentity := request.GetString("externalIdentity", "")
	if externalIdentity != "" {
		params.ExternalIdentity = &externalIdentity
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

func registerUsersUpdate(s *server.MCPServer) {
	tool := mcp.NewTool("users_update",
		mcp.WithDescription("Update a user.<br/>Requires Administer System permission "),
		mcp.WithString("email",
			mcp.Description("User email "),
		),
		mcp.WithString("login",
			mcp.Description("User login "),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("User name "),
		),
		mcp.WithString("scmAccount",
			mcp.Description("SCM accounts. To set several values, the parameter must be called once for each value. "),
		),
	)

	s.AddTool(tool, usersUpdateHandler)
}

func usersUpdateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersUpdate(request)
	return toResult(c.ApiUsersUpdate(ctx, &params, authorizationHeader))
}

func parseUsersUpdate(request mcp.CallToolRequest) client.ApiUsersUpdateParams {
	params := client.ApiUsersUpdateParams{}

	email := request.GetString("email", "")
	if email != "" {
		params.Email = &email
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	scmAccount := request.GetString("scmAccount", "")
	if scmAccount != "" {
		params.ScmAccount = &scmAccount
	}

	return params
}

func registerUsersUpdateIdentityProvider(s *server.MCPServer) {
	tool := mcp.NewTool("users_update_identity_provider",
		mcp.WithDescription("Update identity provider information. <br/>It's only possible to migrate to an installed identity provider. Be careful that as soon as this information has been updated for a user, the user will only be able to authenticate on the new identity provider. It is not possible to migrate external user to local one.<br/>Requires Administer System permission. "),
		mcp.WithString("login",
			mcp.Description("User login "),
			mcp.Required(),
		),
		mcp.WithString("newExternalIdentity",
			mcp.Description("New external identity, usually the login used in the authentication system. If not provided previous identity will be used. "),
		),
		mcp.WithString("newExternalProvider",
			mcp.Description("New external provider. Only authentication system installed are available. Use 'LDAP' identity provider for single server LDAP setup.User 'LDAP_{serverKey}' identity provider for multiple LDAP server setup. "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, usersUpdateIdentityProviderHandler)
}

func usersUpdateIdentityProviderHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersUpdateIdentityProvider(request)
	return toResult(c.ApiUsersUpdateIdentityProvider(ctx, &params, authorizationHeader))
}

func parseUsersUpdateIdentityProvider(request mcp.CallToolRequest) client.ApiUsersUpdateIdentityProviderParams {
	params := client.ApiUsersUpdateIdentityProviderParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	newExternalIdentity := request.GetString("newExternalIdentity", "")
	if newExternalIdentity != "" {
		params.NewExternalIdentity = &newExternalIdentity
	}

	newExternalProvider := request.GetString("newExternalProvider", "")
	if newExternalProvider != "" {
		params.NewExternalProvider = newExternalProvider
	}

	return params
}

func registerUsersUpdateLogin(s *server.MCPServer) {
	tool := mcp.NewTool("users_update_login",
		mcp.WithDescription("Update a user login. A login can be updated many times.<br/>Requires Administer System permission "),
		mcp.WithString("login",
			mcp.Description("The current login (case-sensitive) "),
			mcp.Required(),
		),
		mcp.WithString("newLogin",
			mcp.Description("The new login. It must not already exist. "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, usersUpdateLoginHandler)
}

func usersUpdateLoginHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersUpdateLogin(request)
	return toResult(c.ApiUsersUpdateLogin(ctx, &params, authorizationHeader))
}

func parseUsersUpdateLogin(request mcp.CallToolRequest) client.ApiUsersUpdateLoginParams {
	params := client.ApiUsersUpdateLoginParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	newLogin := request.GetString("newLogin", "")
	if newLogin != "" {
		params.NewLogin = newLogin
	}

	return params
}
