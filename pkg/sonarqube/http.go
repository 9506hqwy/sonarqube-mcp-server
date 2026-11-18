package sonarqube

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
	"github.com/mark3labs/mcp-go/mcp"
)

type URLKey struct{}
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
		return errors.New("missing user")
	}

	password, ok := ctx.Value(PasswordKey{}).(string)
	if !ok {
		return errors.New("missing password")
	}

	req.SetBasicAuth(user, password)
	return nil
}

func bearerAuth(ctx context.Context, req *http.Request) error {
	token, ok := ctx.Value(UserTokenKey{}).(string)
	if !ok || token == "" {
		return errors.New("missing token")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

func newClient(ctx context.Context) (*client.ClientWithResponses, error) {
	hc := http.Client{}

	url, ok := ctx.Value(URLKey{}).(string)
	if !ok || url == "" {
		return nil, errors.New("missing url")
	}

	return client.NewClientWithResponses(url, client.WithHTTPClient(&hc))
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
