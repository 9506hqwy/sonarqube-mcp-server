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
		mcp.WithInputSchema[client.ApiUserTokensGenerateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userTokensGenerateHandler))
}

func userTokensGenerateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserTokensGenerateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserTokensGenerate(ctx, &params, authorizationHeader))
}

func registerUserTokensRevoke(s *server.MCPServer) {
	tool := mcp.NewTool("user_tokens_revoke",
		mcp.WithDescription("Revoke a user access token. <br/>It requires administration permissions to specify a 'login' and revoke a token for another user. Otherwise, the token for the current user is revoked."),
		mcp.WithInputSchema[client.ApiUserTokensRevokeParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userTokensRevokeHandler))
}

func userTokensRevokeHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserTokensRevokeParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserTokensRevoke(ctx, &params, authorizationHeader))
}

func registerUserTokensSearch(s *server.MCPServer) {
	tool := mcp.NewTool("user_tokens_search",
		mcp.WithDescription("List the access tokens of a user.<br>The login must exist and active.<br>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user is using a token many times in less than one hour.<br> It requires administration permissions to specify a 'login' and list the tokens of another user. Otherwise, tokens for the current user are listed. <br> Authentication is required for this API endpoint"),
		mcp.WithInputSchema[client.ApiUserTokensSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(userTokensSearchHandler))
}

func userTokensSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiUserTokensSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiUserTokensSearch(ctx, &params, authorizationHeader))
}
