package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerComponentsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("components_search",
		mcp.WithDescription("Search for components"),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>"),
		),
		mcp.WithString("qualifiers",
			mcp.Description("Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>TRK - Projects</li></ul>"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, componentsSearchHandler)
}

func componentsSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseComponentsSearch(request)
	return toResult(c.ApiComponentsSearch(ctx, &params, authorizationHeader))
}

func parseComponentsSearch(request mcp.CallToolRequest) client.ApiComponentsSearchParams {
	params := client.ApiComponentsSearchParams{}

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

	qualifiers := request.GetString("qualifiers", "")
	if qualifiers != "" {
		params.Qualifiers = qualifiers
	}

	return params
}

func registerComponentsShow(s *server.MCPServer) {
	tool := mcp.NewTool("components_show",
		mcp.WithDescription("Returns a component (file, directory, project, portfolio…) and its ancestors. The ancestors are ordered from the parent to the root project. Requires the following permission: 'Browse' on the project of the specified component."),
		mcp.WithString("branch",
			mcp.Description("Branch key. Not available in the community edition."),
		),
		mcp.WithString("component",
			mcp.Description("Component key"),
			mcp.Required(),
		),
		mcp.WithString("pullRequest",
			mcp.Description("Pull request id. Not available in the community edition."),
		),
	)

	s.AddTool(tool, componentsShowHandler)
}

func componentsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseComponentsShow(request)
	return toResult(c.ApiComponentsShow(ctx, &params, authorizationHeader))
}

func parseComponentsShow(request mcp.CallToolRequest) client.ApiComponentsShowParams {
	params := client.ApiComponentsShowParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = component
	}

	pullRequest := request.GetString("pullRequest", "")
	if pullRequest != "" {
		params.PullRequest = &pullRequest
	}

	return params
}

func registerComponentsTree(s *server.MCPServer) {
	tool := mcp.NewTool("components_tree",
		mcp.WithDescription("Navigate through components based on the chosen strategy.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned."),
		mcp.WithString("asc",
			mcp.Description("Ascending sort"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch key. Not available in the community edition."),
		),
		mcp.WithString("component",
			mcp.Description("Base component key. The search is based on this component."),
			mcp.Required(),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("pullRequest",
			mcp.Description("Pull request id. Not available in the community edition."),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>"),
		),
		mcp.WithString("qualifiers",
			mcp.Description("Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>BRC - no description available</li><li>UTS - Test Files</li><li>FIL - Files</li><li>DIR - Directories</li><li>TRK - Projects</li></ul>"),
		),
		mcp.WithString("s",
			mcp.Description("Comma-separated list of sort fields"),
		),
		mcp.WithString("strategy",
			mcp.Description("Strategy to search for base component descendants:<ul><li>children: return the children components of the base component. Grandchildren components are not returned</li><li>all: return all the descendants components of the base component. Grandchildren are returned.</li><li>leaves: return all the descendant components (files, in general) which don't have other children. They are the leaves of the component tree.</li></ul>"),
		),
	)

	s.AddTool(tool, componentsTreeHandler)
}

func componentsTreeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseComponentsTree(request)
	return toResult(c.ApiComponentsTree(ctx, &params, authorizationHeader))
}

func parseComponentsTree(request mcp.CallToolRequest) client.ApiComponentsTreeParams {
	params := client.ApiComponentsTreeParams{}

	asc := request.GetString("asc", "")
	if asc != "" {
		params.Asc = &asc
	}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = component
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	pullRequest := request.GetString("pullRequest", "")
	if pullRequest != "" {
		params.PullRequest = &pullRequest
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	qualifiers := request.GetString("qualifiers", "")
	if qualifiers != "" {
		params.Qualifiers = &qualifiers
	}

	s := request.GetString("s", "")
	if s != "" {
		params.S = &s
	}

	strategy := request.GetString("strategy", "")
	if strategy != "" {
		params.Strategy = &strategy
	}

	return params
}
