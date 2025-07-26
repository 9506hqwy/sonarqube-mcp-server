package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerUserTokensGenerate(s *server.MCPServer) {
	tool := mcp.NewTool("user_tokens_generate",
		mcp.WithDescription("Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />It requires administration permissions to specify a 'login' and generate a token for another user. Otherwise, a token is generated for the current user."),
		mcp.WithString("expirationDate",
			mcp.Description("The expiration date of the token being generated, in ISO 8601 format (YYYY-MM-DD). If not set, default to no expiration."),
		),
		mcp.WithString("login",
			mcp.Description("User login. If not set, the token is generated for the authenticated user."),
		),
		mcp.WithString("name",
			mcp.Description("Token name"),
			mcp.Required(),
		),
		mcp.WithString("projectKey",
			mcp.Description("The key of the only project that can be analyzed by the PROJECT_ANALYSIS_TOKEN being generated."),
		),
		mcp.WithString("type",
			mcp.Description("Token Type. If this parameters is set to PROJECT_ANALYSIS_TOKEN, it is necessary to provide the projectKey parameter too."),
		),
	)

	s.AddTool(tool, userTokensGenerateHandler)
}

func userTokensGenerateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserTokensGenerate(request)
	return toResult(c.ApiUserTokensGenerate(ctx, &params, authorizationHeader))
}

func parseUserTokensGenerate(request mcp.CallToolRequest) client.ApiUserTokensGenerateParams {
	params := client.ApiUserTokensGenerateParams{}

	expirationDate := request.GetString("expirationDate", "")
	if expirationDate != "" {
		params.ExpirationDate = &expirationDate
	}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = &projectKey
	}

	_type := request.GetString("type", "")
	if _type != "" {
		params.Type = &_type
	}

	return params
}

func registerUserTokensRevoke(s *server.MCPServer) {
	tool := mcp.NewTool("user_tokens_revoke",
		mcp.WithDescription("Revoke a user access token. <br/>It requires administration permissions to specify a 'login' and revoke a token for another user. Otherwise, the token for the current user is revoked."),
		mcp.WithString("login",
			mcp.Description("User login"),
		),
		mcp.WithString("name",
			mcp.Description("Token name"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, userTokensRevokeHandler)
}

func userTokensRevokeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserTokensRevoke(request)
	return toResult(c.ApiUserTokensRevoke(ctx, &params, authorizationHeader))
}

func parseUserTokensRevoke(request mcp.CallToolRequest) client.ApiUserTokensRevokeParams {
	params := client.ApiUserTokensRevokeParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}

func registerUserTokensSearch(s *server.MCPServer) {
	tool := mcp.NewTool("user_tokens_search",
		mcp.WithDescription("List the access tokens of a user.<br>The login must exist and active.<br>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user is using a token many times in less than one hour.<br> It requires administration permissions to specify a 'login' and list the tokens of another user. Otherwise, tokens for the current user are listed. <br> Authentication is required for this API endpoint"),
		mcp.WithString("login",
			mcp.Description("User login"),
		),
	)

	s.AddTool(tool, userTokensSearchHandler)
}

func userTokensSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUserTokensSearch(request)
	return toResult(c.ApiUserTokensSearch(ctx, &params, authorizationHeader))
}

func parseUserTokensSearch(request mcp.CallToolRequest) client.ApiUserTokensSearchParams {
	params := client.ApiUserTokensSearchParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = &login
	}

	return params
}
