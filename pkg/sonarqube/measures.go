package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerMeasuresComponent(s *server.MCPServer) {
	tool := mcp.NewTool("measures_component",
		mcp.WithDescription("Return component with specified measures.<br>Requires the following permission: 'Browse' on the project of specified component."),
		mcp.WithString("additionalFields",
			mcp.Description("Comma-separated list of additional fields that can be returned in the response."),
		),
		mcp.WithString("branch",
			mcp.Description("Branch key. Not available in the community edition."),
		),
		mcp.WithString("component",
			mcp.Description("Component key"),
			mcp.Required(),
		),
		mcp.WithString("metricKeys",
			mcp.Description("Comma-separated list of metric keys"),
			mcp.Required(),
		),
		mcp.WithString("pullRequest",
			mcp.Description("Pull request id. Not available in the community edition."),
		),
	)

	s.AddTool(tool, measuresComponentHandler)
}

func measuresComponentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseMeasuresComponent(request)
	return toResult(c.ApiMeasuresComponent(ctx, &params, authorizationHeader))
}

func parseMeasuresComponent(request mcp.CallToolRequest) client.ApiMeasuresComponentParams {
	params := client.ApiMeasuresComponentParams{}

	additionalFields := request.GetString("additionalFields", "")
	if additionalFields != "" {
		params.AdditionalFields = &additionalFields
	}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = component
	}

	metricKeys := request.GetString("metricKeys", "")
	if metricKeys != "" {
		params.MetricKeys = metricKeys
	}

	pullRequest := request.GetString("pullRequest", "")
	if pullRequest != "" {
		params.PullRequest = &pullRequest
	}

	return params
}

func registerMeasuresComponentTree(s *server.MCPServer) {
	tool := mcp.NewTool("measures_component_tree",
		mcp.WithDescription("Navigate through components based on the chosen strategy with specified measures.<br>Requires the following permission: 'Browse' on the specified project.<br>For applications, it also requires 'Browse' permission on its child projects. <br>When limiting search with the q parameter, directories are not returned."),
		mcp.WithString("additionalFields",
			mcp.Description("Comma-separated list of additional fields that can be returned in the response."),
		),
		mcp.WithString("asc",
			mcp.Description("Ascending sort"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch key. Not available in the community edition."),
		),
		mcp.WithString("component",
			mcp.Description("Component key. The search is based on this component."),
			mcp.Required(),
		),
		mcp.WithString("metricKeys",
			mcp.Description("Comma-separated list of metric keys. Types DATA, DISTRIB are not allowed."),
			mcp.Required(),
		),
		mcp.WithString("metricPeriodSort",
			mcp.Description("Sort measures by leak period or not ?. The 's' parameter must contain the 'metricPeriod' value."),
		),
		mcp.WithString("metricSort",
			mcp.Description("Metric key to sort by. The 's' parameter must contain the 'metric' or 'metricPeriod' value. It must be part of the 'metricKeys' parameter"),
		),
		mcp.WithString("metricSortFilter",
			mcp.Description("Filter components. Sort must be on a metric. Possible values are: <ul><li>all: return all components</li><li>withMeasuresOnly: filter out components that do not have a measure on the sorted metric</li></ul>"),
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

	s.AddTool(tool, measuresComponentTreeHandler)
}

func measuresComponentTreeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseMeasuresComponentTree(request)
	return toResult(c.ApiMeasuresComponentTree(ctx, &params, authorizationHeader))
}

func parseMeasuresComponentTree(request mcp.CallToolRequest) client.ApiMeasuresComponentTreeParams {
	params := client.ApiMeasuresComponentTreeParams{}

	additionalFields := request.GetString("additionalFields", "")
	if additionalFields != "" {
		params.AdditionalFields = &additionalFields
	}

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

	metricKeys := request.GetString("metricKeys", "")
	if metricKeys != "" {
		params.MetricKeys = metricKeys
	}

	metricPeriodSort := request.GetString("metricPeriodSort", "")
	if metricPeriodSort != "" {
		params.MetricPeriodSort = &metricPeriodSort
	}

	metricSort := request.GetString("metricSort", "")
	if metricSort != "" {
		params.MetricSort = &metricSort
	}

	metricSortFilter := request.GetString("metricSortFilter", "")
	if metricSortFilter != "" {
		params.MetricSortFilter = &metricSortFilter
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

func registerMeasuresSearchHistory(s *server.MCPServer) {
	tool := mcp.NewTool("measures_search_history",
		mcp.WithDescription("Search measures history of a component.<br>Measures are ordered chronologically.<br>Pagination applies to the number of measures for each metric.<br>Requires the following permission: 'Browse' on the specified component. <br>For applications, it also requires 'Browse' permission on its child projects."),
		mcp.WithString("branch",
			mcp.Description("Branch key. Not available in the community edition."),
		),
		mcp.WithString("component",
			mcp.Description("Component key"),
			mcp.Required(),
		),
		mcp.WithString("from",
			mcp.Description("Filter measures created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided"),
		),
		mcp.WithString("metrics",
			mcp.Description("Comma-separated list of metric keys"),
			mcp.Required(),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 1000"),
		),
		mcp.WithString("pullRequest",
			mcp.Description("Pull request id. Not available in the community edition."),
		),
		mcp.WithString("to",
			mcp.Description("Filter measures created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided"),
		),
	)

	s.AddTool(tool, measuresSearchHistoryHandler)
}

func measuresSearchHistoryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseMeasuresSearchHistory(request)
	return toResult(c.ApiMeasuresSearchHistory(ctx, &params, authorizationHeader))
}

func parseMeasuresSearchHistory(request mcp.CallToolRequest) client.ApiMeasuresSearchHistoryParams {
	params := client.ApiMeasuresSearchHistoryParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	component := request.GetString("component", "")
	if component != "" {
		params.Component = component
	}

	from := request.GetString("from", "")
	if from != "" {
		params.From = &from
	}

	metrics := request.GetString("metrics", "")
	if metrics != "" {
		params.Metrics = metrics
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

	to := request.GetString("to", "")
	if to != "" {
		params.To = &to
	}

	return params
}
