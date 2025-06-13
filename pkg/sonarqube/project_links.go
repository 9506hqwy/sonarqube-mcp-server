package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerProjectLinksCreate(s *server.MCPServer) {
	tool := mcp.NewTool("project_links_create",
		mcp.WithDescription("Create a new project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission. "),
		mcp.WithString("name",
			mcp.Description("Link name "),
			mcp.Required(),
		),
		mcp.WithString("projectId",
			mcp.Description("Project id "),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key "),
		),
		mcp.WithString("url",
			mcp.Description("Link url "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectLinksCreateHandler)
}

func projectLinksCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectLinksCreate(request)
	return toResult(c.ApiProjectLinksCreate(ctx, &params, authorizationHeader))
}

func parseProjectLinksCreate(request mcp.CallToolRequest) client.ApiProjectLinksCreateParams {
	params := client.ApiProjectLinksCreateParams{}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	projectId := request.GetString("projectId", "")
	if projectId != "" {
		params.ProjectId = &projectId
	}

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = &projectKey
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerProjectLinksDelete(s *server.MCPServer) {
	tool := mcp.NewTool("project_links_delete",
		mcp.WithDescription("Delete existing project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission. "),
		mcp.WithString("id",
			mcp.Description("Link id "),
			mcp.Required(),
		),
	)

	s.AddTool(tool, projectLinksDeleteHandler)
}

func projectLinksDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectLinksDelete(request)
	return toResult(c.ApiProjectLinksDelete(ctx, &params, authorizationHeader))
}

func parseProjectLinksDelete(request mcp.CallToolRequest) client.ApiProjectLinksDeleteParams {
	params := client.ApiProjectLinksDeleteParams{}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = id
	}

	return params
}

func registerProjectLinksSearch(s *server.MCPServer) {
	tool := mcp.NewTool("project_links_search",
		mcp.WithDescription("List links of a project.<br>The 'projectId' or 'projectKey' must be provided.<br>Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul> "),
		mcp.WithString("projectId",
			mcp.Description("Project Id "),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project Key "),
		),
	)

	s.AddTool(tool, projectLinksSearchHandler)
}

func projectLinksSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectLinksSearch(request)
	return toResult(c.ApiProjectLinksSearch(ctx, &params, authorizationHeader))
}

func parseProjectLinksSearch(request mcp.CallToolRequest) client.ApiProjectLinksSearchParams {
	params := client.ApiProjectLinksSearchParams{}

	projectId := request.GetString("projectId", "")
	if projectId != "" {
		params.ProjectId = &projectId
	}

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = &projectKey
	}

	return params
}
