package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerAuthenticationLogin(s *server.MCPServer) {
	tool := mcp.NewTool("authentication_login",
		mcp.WithDescription("Authenticate a user. "),
		mcp.WithString("login",
			mcp.Description("Login of the user "),
			mcp.Required(),
		),
		mcp.WithString("password",
			mcp.Description("Password of the user "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, authenticationLoginHandler)
}

func authenticationLoginHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAuthenticationLogin(request)
	return toResult(c.ApiAuthenticationLogin(ctx, &params, authorizationHeader))
}

func parseAuthenticationLogin(request mcp.CallToolRequest) client.ApiAuthenticationLoginParams {
	params := client.ApiAuthenticationLoginParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	password := request.GetString("password", "")
	if password != "" {
		params.Password = password
	}

	return params
}

func registerAuthenticationLogout(s *server.MCPServer) {
	tool := mcp.NewTool("authentication_logout",
		mcp.WithDescription("Logout a user. "),
	)

	s.AddTool(tool, authenticationLogoutHandler)
}

func authenticationLogoutHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAuthenticationLogout(ctx, authorizationHeader))
}

func registerAuthenticationValidate(s *server.MCPServer) {
	tool := mcp.NewTool("authentication_validate",
		mcp.WithDescription("Check credentials. "),
	)

	s.AddTool(tool, authenticationValidateHandler)
}

func authenticationValidateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAuthenticationValidate(ctx, authorizationHeader))
}
