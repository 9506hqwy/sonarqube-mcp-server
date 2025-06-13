package sonarqube

import (
	"context"
	"fmt"
	"io"
	"net/http"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
	"github.com/mark3labs/mcp-go/mcp"
)

type ServerKey struct{}
type PortKey struct{}
type UserKey struct{}
type UserTokenKey struct{}
type PasswordKey struct{}

func authorizationHeader(ctx context.Context, req *http.Request) error {
	if bearerAuth(ctx, req) == nil {
		return nil
	}

	return basicAuth(ctx, req)
}

func basicAuth(ctx context.Context, req *http.Request) error {
	user, ok := ctx.Value(UserKey{}).(string)
	if !ok || user == "" {
		return fmt.Errorf("missing user")
	}

	password, ok := ctx.Value(PasswordKey{}).(string)
	if !ok {
		return fmt.Errorf("missing password")
	}

	req.SetBasicAuth(user, password)
	return nil
}

func bearerAuth(ctx context.Context, req *http.Request) error {
	token, ok := ctx.Value(UserTokenKey{}).(string)
	if !ok || token == "" {
		return fmt.Errorf("missing token")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

func newClient(ctx context.Context) (*client.ClientWithResponses, error) {
	hc := http.Client{}

	server, ok := ctx.Value(ServerKey{}).(string)
	if !ok {
		return nil, fmt.Errorf("missing server")
	}

	port, ok := ctx.Value(PortKey{}).(uint16)
	if !ok {
		return nil, fmt.Errorf("missing port")
	}

	return client.NewClientWithResponses(fmt.Sprintf("http://%s:%d", server, port), client.WithHTTPClient(&hc))
}

func toResult(response *http.Response, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if response.StatusCode < http.StatusOK || http.StatusMultipleChoices <= response.StatusCode {
		return mcp.NewToolResultError(fmt.Sprintf("%s: %s", response.Status, string(body))), nil
	}

	return mcp.NewToolResultText(string(body)), nil
}
