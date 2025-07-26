package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerIssuesAddComment(s *server.MCPServer) {
	tool := mcp.NewTool("issues_add_comment",
		mcp.WithDescription("Add a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue."),
		mcp.WithString("issue",
			mcp.Description("Issue key"),
			mcp.Required(),
		),
		mcp.WithString("text",
			mcp.Description("Comment text"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesAddCommentHandler)
}

func issuesAddCommentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesAddComment(request)
	return toResult(c.ApiIssuesAddComment(ctx, &params, authorizationHeader))
}

func parseIssuesAddComment(request mcp.CallToolRequest) client.ApiIssuesAddCommentParams {
	params := client.ApiIssuesAddCommentParams{}

	issue := request.GetString("issue", "")
	if issue != "" {
		params.Issue = issue
	}

	text := request.GetString("text", "")
	if text != "" {
		params.Text = text
	}

	return params
}

func registerIssuesAssign(s *server.MCPServer) {
	tool := mcp.NewTool("issues_assign",
		mcp.WithDescription("Assign/Unassign an issue. Requires authentication and Browse permission on project"),
		mcp.WithString("assignee",
			mcp.Description("Login of the assignee. When not set, it will unassign the issue. Use '_me' to assign to current user"),
		),
		mcp.WithString("issue",
			mcp.Description("Issue key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesAssignHandler)
}

func issuesAssignHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesAssign(request)
	return toResult(c.ApiIssuesAssign(ctx, &params, authorizationHeader))
}

func parseIssuesAssign(request mcp.CallToolRequest) client.ApiIssuesAssignParams {
	params := client.ApiIssuesAssignParams{}

	assignee := request.GetString("assignee", "")
	if assignee != "" {
		params.Assignee = &assignee
	}

	issue := request.GetString("issue", "")
	if issue != "" {
		params.Issue = issue
	}

	return params
}

func registerIssuesAuthors(s *server.MCPServer) {
	tool := mcp.NewTool("issues_authors",
		mcp.WithDescription("Search SCM accounts which match a given query.<br/>Requires authentication.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code."),
		mcp.WithString("project",
			mcp.Description("Project key"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 100"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to authors that contain the supplied string."),
		),
	)

	s.AddTool(tool, issuesAuthorsHandler)
}

func issuesAuthorsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesAuthors(request)
	return toResult(c.ApiIssuesAuthors(ctx, &params, authorizationHeader))
}

func parseIssuesAuthors(request mcp.CallToolRequest) client.ApiIssuesAuthorsParams {
	params := client.ApiIssuesAuthorsParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
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

func registerIssuesBulkChange(s *server.MCPServer) {
	tool := mcp.NewTool("issues_bulk_change",
		mcp.WithDescription("Bulk change on issues. Up to 500 issues can be updated. <br/>Requires authentication."),
		mcp.WithString("add_tags",
			mcp.Description("Add tags"),
		),
		mcp.WithString("assign",
			mcp.Description("To assign the list of issues to a specific user (login), or un-assign all the issues"),
		),
		mcp.WithString("comment",
			mcp.Description("Add a comment. The comment will only be added to issues that are affected either by a change of type or a change of severity as a result of the same WS call."),
		),
		mcp.WithString("do_transition",
			mcp.Description("Transition"),
		),
		mcp.WithString("issues",
			mcp.Description("Comma-separated list of issue keys"),
			mcp.Required(),
		),
		mcp.WithString("remove_tags",
			mcp.Description("Remove tags"),
		),
		mcp.WithString("sendNotifications",
			mcp.Description("null"),
		),
		mcp.WithString("set_severity",
			mcp.Description("To change the severity of the list of issues"),
		),
		mcp.WithString("set_type",
			mcp.Description("To change the type of the list of issues"),
		),
	)

	s.AddTool(tool, issuesBulkChangeHandler)
}

func issuesBulkChangeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesBulkChange(request)
	return toResult(c.ApiIssuesBulkChange(ctx, &params, authorizationHeader))
}

func parseIssuesBulkChange(request mcp.CallToolRequest) client.ApiIssuesBulkChangeParams {
	params := client.ApiIssuesBulkChangeParams{}

	add_tags := request.GetString("add_tags", "")
	if add_tags != "" {
		params.AddTags = &add_tags
	}

	assign := request.GetString("assign", "")
	if assign != "" {
		params.Assign = &assign
	}

	comment := request.GetString("comment", "")
	if comment != "" {
		params.Comment = &comment
	}

	do_transition := request.GetString("do_transition", "")
	if do_transition != "" {
		params.DoTransition = &do_transition
	}

	issues := request.GetString("issues", "")
	if issues != "" {
		params.Issues = issues
	}

	remove_tags := request.GetString("remove_tags", "")
	if remove_tags != "" {
		params.RemoveTags = &remove_tags
	}

	sendNotifications := request.GetString("sendNotifications", "")
	if sendNotifications != "" {
		params.SendNotifications = &sendNotifications
	}

	set_severity := request.GetString("set_severity", "")
	if set_severity != "" {
		params.SetSeverity = &set_severity
	}

	set_type := request.GetString("set_type", "")
	if set_type != "" {
		params.SetType = &set_type
	}

	return params
}

func registerIssuesChangelog(s *server.MCPServer) {
	tool := mcp.NewTool("issues_changelog",
		mcp.WithDescription("Display changelog of an issue.<br/>Requires the 'Browse' permission on the project of the specified issue."),
		mcp.WithString("issue",
			mcp.Description("Issue key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesChangelogHandler)
}

func issuesChangelogHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesChangelog(request)
	return toResult(c.ApiIssuesChangelog(ctx, &params, authorizationHeader))
}

func parseIssuesChangelog(request mcp.CallToolRequest) client.ApiIssuesChangelogParams {
	params := client.ApiIssuesChangelogParams{}

	issue := request.GetString("issue", "")
	if issue != "" {
		params.Issue = issue
	}

	return params
}

func registerIssuesDeleteComment(s *server.MCPServer) {
	tool := mcp.NewTool("issues_delete_comment",
		mcp.WithDescription("Delete a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue."),
		mcp.WithString("comment",
			mcp.Description("Comment key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesDeleteCommentHandler)
}

func issuesDeleteCommentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesDeleteComment(request)
	return toResult(c.ApiIssuesDeleteComment(ctx, &params, authorizationHeader))
}

func parseIssuesDeleteComment(request mcp.CallToolRequest) client.ApiIssuesDeleteCommentParams {
	params := client.ApiIssuesDeleteCommentParams{}

	comment := request.GetString("comment", "")
	if comment != "" {
		params.Comment = comment
	}

	return params
}

func registerIssuesDoTransition(s *server.MCPServer) {
	tool := mcp.NewTool("issues_do_transition",
		mcp.WithDescription("Do workflow transition on an issue. Requires authentication and Browse permission on project.<br/>The transitions 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.<br/>The transitions involving security hotspots require the permission 'Administer Security Hotspot'."),
		mcp.WithString("issue",
			mcp.Description("Issue key"),
			mcp.Required(),
		),
		mcp.WithString("transition",
			mcp.Description("Transition"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesDoTransitionHandler)
}

func issuesDoTransitionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesDoTransition(request)
	return toResult(c.ApiIssuesDoTransition(ctx, &params, authorizationHeader))
}

func parseIssuesDoTransition(request mcp.CallToolRequest) client.ApiIssuesDoTransitionParams {
	params := client.ApiIssuesDoTransitionParams{}

	issue := request.GetString("issue", "")
	if issue != "" {
		params.Issue = issue
	}

	transition := request.GetString("transition", "")
	if transition != "" {
		params.Transition = transition
	}

	return params
}

func registerIssuesEditComment(s *server.MCPServer) {
	tool := mcp.NewTool("issues_edit_comment",
		mcp.WithDescription("Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue."),
		mcp.WithString("comment",
			mcp.Description("Comment key"),
			mcp.Required(),
		),
		mcp.WithString("text",
			mcp.Description("Comment text"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesEditCommentHandler)
}

func issuesEditCommentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesEditComment(request)
	return toResult(c.ApiIssuesEditComment(ctx, &params, authorizationHeader))
}

func parseIssuesEditComment(request mcp.CallToolRequest) client.ApiIssuesEditCommentParams {
	params := client.ApiIssuesEditCommentParams{}

	comment := request.GetString("comment", "")
	if comment != "" {
		params.Comment = comment
	}

	text := request.GetString("text", "")
	if text != "" {
		params.Text = text
	}

	return params
}

func registerIssuesReindex(s *server.MCPServer) {
	tool := mcp.NewTool("issues_reindex",
		mcp.WithDescription("Reindex issues for a project.<br> Require 'Administer System' permission."),
		mcp.WithString("project",
			mcp.Description("Project key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesReindexHandler)
}

func issuesReindexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesReindex(request)
	return toResult(c.ApiIssuesReindex(ctx, &params, authorizationHeader))
}

func parseIssuesReindex(request mcp.CallToolRequest) client.ApiIssuesReindexParams {
	params := client.ApiIssuesReindexParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerIssuesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("issues_search",
		mcp.WithDescription("Search for issues.<br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code."),
		mcp.WithString("additionalFields",
			mcp.Description("Comma-separated list of the optional fields to be returned in response. Action plans are dropped in 5.5, it is not returned in the response."),
		),
		mcp.WithString("asc",
			mcp.Description("Ascending sort"),
		),
		mcp.WithString("assigned",
			mcp.Description("To retrieve assigned or unassigned issues"),
		),
		mcp.WithString("assignees",
			mcp.Description("Comma-separated list of assignee logins. The value '__me__' can be used as a placeholder for user who performs the request"),
		),
		mcp.WithString("author",
			mcp.Description("SCM accounts. To set several values, the parameter must be called once for each value."),
		),
		mcp.WithString("branch",
			mcp.Description("Branch key. Not available in the community edition."),
		),
		mcp.WithString("componentKeys",
			mcp.Description("Comma-separated list of component keys. Retrieve issues associated to a specific list of components (and all its descendants). A component can be a portfolio, project, module, directory or file."),
		),
		mcp.WithString("createdAfter",
			mcp.Description("To retrieve issues created after the given date (inclusive). <br>Either a date (use 'timeZone' attribute or it will default to server timezone) or datetime can be provided. <br>If this parameter is set, createdInLast must not be set"),
		),
		mcp.WithString("createdAt",
			mcp.Description("Datetime to retrieve issues created during a specific analysis"),
		),
		mcp.WithString("createdBefore",
			mcp.Description("To retrieve issues created before the given date (exclusive). <br>Either a date (use 'timeZone' attribute or it will default to server timezone) or datetime can be provided."),
		),
		mcp.WithString("createdInLast",
			mcp.Description("To retrieve issues created during a time span before the current time (exclusive). Accepted units are 'y' for year, 'm' for month, 'w' for week and 'd' for day. If this parameter is set, createdAfter must not be set"),
		),
		mcp.WithString("cwe",
			mcp.Description("Comma-separated list of CWE identifiers. Use 'unknown' to select issues not associated to any CWE."),
		),
		mcp.WithString("facets",
			mcp.Description("Comma-separated list of the facets to be computed. No facet is computed by default."),
		),
		mcp.WithString("inNewCodePeriod",
			mcp.Description("To retrieve issues created in the new code period.<br>If this parameter is set to a truthy value, createdAfter must not be set and one component uuid or key must be provided."),
		),
		mcp.WithString("issues",
			mcp.Description("Comma-separated list of issue keys"),
		),
		mcp.WithString("languages",
			mcp.Description("Comma-separated list of languages. Available since 4.4"),
		),
		mcp.WithString("onComponentOnly",
			mcp.Description("Return only issues at a component's level, not on its descendants (modules, directories, files, etc). This parameter is only considered when componentKeys is set."),
		),
		mcp.WithString("owaspAsvs-4.0",
			mcp.Description("Comma-separated list of OWASP ASVS v4.0 categories."),
		),
		mcp.WithString("owaspAsvsLevel",
			mcp.Description("Level of OWASP ASVS categories."),
		),
		mcp.WithString("owaspTop10",
			mcp.Description("Comma-separated list of OWASP Top 10 2017 lowercase categories."),
		),
		mcp.WithString("owaspTop10-2021",
			mcp.Description("Comma-separated list of OWASP Top 10 2021 lowercase categories."),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("pciDss-3.2",
			mcp.Description("Comma-separated list of PCI DSS v3.2 categories."),
		),
		mcp.WithString("pciDss-4.0",
			mcp.Description("Comma-separated list of PCI DSS v4.0 categories."),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("pullRequest",
			mcp.Description("Pull request id. Not available in the community edition."),
		),
		mcp.WithString("resolutions",
			mcp.Description("Comma-separated list of resolutions"),
		),
		mcp.WithString("resolved",
			mcp.Description("To match resolved or unresolved issues"),
		),
		mcp.WithString("rules",
			mcp.Description("Comma-separated list of coding rule keys. Format is &lt;repository&gt;:&lt;rule&gt;"),
		),
		mcp.WithString("s",
			mcp.Description("Sort field"),
		),
		mcp.WithString("sansTop25",
			mcp.Description("Comma-separated list of SANS Top 25 categories."),
		),
		mcp.WithString("scopes",
			mcp.Description("Comma-separated list of scopes. Available since 8.5"),
		),
		mcp.WithString("severities",
			mcp.Description("Comma-separated list of severities"),
		),
		mcp.WithString("sinceLeakPeriod",
			mcp.Description("To retrieve issues created since the leak period.<br>If this parameter is set to a truthy value, createdAfter must not be set and one component uuid or key must be provided."),
		),
		mcp.WithString("sonarsourceSecurity",
			mcp.Description("Comma-separated list of SonarSource security categories. Use 'others' to select issues not associated with any category"),
		),
		mcp.WithString("statuses",
			mcp.Description("Comma-separated list of statuses"),
		),
		mcp.WithString("tags",
			mcp.Description("Comma-separated list of tags."),
		),
		mcp.WithString("timeZone",
			mcp.Description("To resolve dates passed to 'createdAfter' or 'createdBefore' (does not apply to datetime) and to compute creation date histogram"),
		),
		mcp.WithString("types",
			mcp.Description("Comma-separated list of types."),
		),
	)

	s.AddTool(tool, issuesSearchHandler)
}

func issuesSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesSearch(request)
	return toResult(c.ApiIssuesSearch(ctx, &params, authorizationHeader))
}

func parseIssuesSearch(request mcp.CallToolRequest) client.ApiIssuesSearchParams {
	params := client.ApiIssuesSearchParams{}

	additionalFields := request.GetString("additionalFields", "")
	if additionalFields != "" {
		params.AdditionalFields = &additionalFields
	}

	asc := request.GetString("asc", "")
	if asc != "" {
		params.Asc = &asc
	}

	assigned := request.GetString("assigned", "")
	if assigned != "" {
		params.Assigned = &assigned
	}

	assignees := request.GetString("assignees", "")
	if assignees != "" {
		params.Assignees = &assignees
	}

	author := request.GetString("author", "")
	if author != "" {
		params.Author = &author
	}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	componentKeys := request.GetString("componentKeys", "")
	if componentKeys != "" {
		params.ComponentKeys = &componentKeys
	}

	createdAfter := request.GetString("createdAfter", "")
	if createdAfter != "" {
		params.CreatedAfter = &createdAfter
	}

	createdAt := request.GetString("createdAt", "")
	if createdAt != "" {
		params.CreatedAt = &createdAt
	}

	createdBefore := request.GetString("createdBefore", "")
	if createdBefore != "" {
		params.CreatedBefore = &createdBefore
	}

	createdInLast := request.GetString("createdInLast", "")
	if createdInLast != "" {
		params.CreatedInLast = &createdInLast
	}

	cwe := request.GetString("cwe", "")
	if cwe != "" {
		params.Cwe = &cwe
	}

	facets := request.GetString("facets", "")
	if facets != "" {
		params.Facets = &facets
	}

	inNewCodePeriod := request.GetString("inNewCodePeriod", "")
	if inNewCodePeriod != "" {
		params.InNewCodePeriod = &inNewCodePeriod
	}

	issues := request.GetString("issues", "")
	if issues != "" {
		params.Issues = &issues
	}

	languages := request.GetString("languages", "")
	if languages != "" {
		params.Languages = &languages
	}

	onComponentOnly := request.GetString("onComponentOnly", "")
	if onComponentOnly != "" {
		params.OnComponentOnly = &onComponentOnly
	}

	owaspAsvs40 := request.GetString("owaspAsvs-4.0", "")
	if owaspAsvs40 != "" {
		params.OwaspAsvs40 = &owaspAsvs40
	}

	owaspAsvsLevel := request.GetString("owaspAsvsLevel", "")
	if owaspAsvsLevel != "" {
		params.OwaspAsvsLevel = &owaspAsvsLevel
	}

	owaspTop10 := request.GetString("owaspTop10", "")
	if owaspTop10 != "" {
		params.OwaspTop10 = &owaspTop10
	}

	owaspTop102021 := request.GetString("owaspTop10-2021", "")
	if owaspTop102021 != "" {
		params.OwaspTop102021 = &owaspTop102021
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	pciDss32 := request.GetString("pciDss-3.2", "")
	if pciDss32 != "" {
		params.PciDss32 = &pciDss32
	}

	pciDss40 := request.GetString("pciDss-4.0", "")
	if pciDss40 != "" {
		params.PciDss40 = &pciDss40
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	pullRequest := request.GetString("pullRequest", "")
	if pullRequest != "" {
		params.PullRequest = &pullRequest
	}

	resolutions := request.GetString("resolutions", "")
	if resolutions != "" {
		params.Resolutions = &resolutions
	}

	resolved := request.GetString("resolved", "")
	if resolved != "" {
		params.Resolved = &resolved
	}

	rules := request.GetString("rules", "")
	if rules != "" {
		params.Rules = &rules
	}

	s := request.GetString("s", "")
	if s != "" {
		params.S = &s
	}

	sansTop25 := request.GetString("sansTop25", "")
	if sansTop25 != "" {
		params.SansTop25 = &sansTop25
	}

	scopes := request.GetString("scopes", "")
	if scopes != "" {
		params.Scopes = &scopes
	}

	severities := request.GetString("severities", "")
	if severities != "" {
		params.Severities = &severities
	}

	sinceLeakPeriod := request.GetString("sinceLeakPeriod", "")
	if sinceLeakPeriod != "" {
		params.SinceLeakPeriod = &sinceLeakPeriod
	}

	sonarsourceSecurity := request.GetString("sonarsourceSecurity", "")
	if sonarsourceSecurity != "" {
		params.SonarsourceSecurity = &sonarsourceSecurity
	}

	statuses := request.GetString("statuses", "")
	if statuses != "" {
		params.Statuses = &statuses
	}

	tags := request.GetString("tags", "")
	if tags != "" {
		params.Tags = &tags
	}

	timeZone := request.GetString("timeZone", "")
	if timeZone != "" {
		params.TimeZone = &timeZone
	}

	types := request.GetString("types", "")
	if types != "" {
		params.Types = &types
	}

	return params
}

func registerIssuesSetSeverity(s *server.MCPServer) {
	tool := mcp.NewTool("issues_set_severity",
		mcp.WithDescription("Change severity.<br/>Requires the following permissions:<ul> <li>'Authentication'</li> <li>'Browse' rights on project of the specified issue</li> <li>'Administer Issues' rights on project of the specified issue</li></ul>"),
		mcp.WithString("issue",
			mcp.Description("Issue key"),
			mcp.Required(),
		),
		mcp.WithString("severity",
			mcp.Description("New severity"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesSetSeverityHandler)
}

func issuesSetSeverityHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesSetSeverity(request)
	return toResult(c.ApiIssuesSetSeverity(ctx, &params, authorizationHeader))
}

func parseIssuesSetSeverity(request mcp.CallToolRequest) client.ApiIssuesSetSeverityParams {
	params := client.ApiIssuesSetSeverityParams{}

	issue := request.GetString("issue", "")
	if issue != "" {
		params.Issue = issue
	}

	severity := request.GetString("severity", "")
	if severity != "" {
		params.Severity = severity
	}

	return params
}

func registerIssuesSetTags(s *server.MCPServer) {
	tool := mcp.NewTool("issues_set_tags",
		mcp.WithDescription("Set tags on an issue. <br/>Requires authentication and Browse permission on project"),
		mcp.WithString("issue",
			mcp.Description("Issue key"),
			mcp.Required(),
		),
		mcp.WithString("tags",
			mcp.Description("Comma-separated list of tags. All tags are removed if parameter is empty or not set."),
		),
	)

	s.AddTool(tool, issuesSetTagsHandler)
}

func issuesSetTagsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesSetTags(request)
	return toResult(c.ApiIssuesSetTags(ctx, &params, authorizationHeader))
}

func parseIssuesSetTags(request mcp.CallToolRequest) client.ApiIssuesSetTagsParams {
	params := client.ApiIssuesSetTagsParams{}

	issue := request.GetString("issue", "")
	if issue != "" {
		params.Issue = issue
	}

	tags := request.GetString("tags", "")
	if tags != "" {
		params.Tags = &tags
	}

	return params
}

func registerIssuesSetType(s *server.MCPServer) {
	tool := mcp.NewTool("issues_set_type",
		mcp.WithDescription("Change type of issue, for instance from 'code smell' to 'bug'.<br/>Requires the following permissions:<ul> <li>'Authentication'</li> <li>'Browse' rights on project of the specified issue</li> <li>'Administer Issues' rights on project of the specified issue</li></ul>"),
		mcp.WithString("issue",
			mcp.Description("Issue key"),
			mcp.Required(),
		),
		mcp.WithString("type",
			mcp.Description("New type"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, issuesSetTypeHandler)
}

func issuesSetTypeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesSetType(request)
	return toResult(c.ApiIssuesSetType(ctx, &params, authorizationHeader))
}

func parseIssuesSetType(request mcp.CallToolRequest) client.ApiIssuesSetTypeParams {
	params := client.ApiIssuesSetTypeParams{}

	issue := request.GetString("issue", "")
	if issue != "" {
		params.Issue = issue
	}

	_type := request.GetString("type", "")
	if _type != "" {
		params.Type = _type
	}

	return params
}

func registerIssuesTags(s *server.MCPServer) {
	tool := mcp.NewTool("issues_tags",
		mcp.WithDescription("List tags matching a given query"),
		mcp.WithString("all",
			mcp.Description("Indicator to search for all tags or only for tags in the main branch of a project"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch key"),
		),
		mcp.WithString("project",
			mcp.Description("Project key"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to tags that contain the supplied string."),
		),
	)

	s.AddTool(tool, issuesTagsHandler)
}

func issuesTagsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesTags(request)
	return toResult(c.ApiIssuesTags(ctx, &params, authorizationHeader))
}

func parseIssuesTags(request mcp.CallToolRequest) client.ApiIssuesTagsParams {
	params := client.ApiIssuesTagsParams{}

	all := request.GetString("all", "")
	if all != "" {
		params.All = &all
	}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
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
