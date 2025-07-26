package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerRulesCreate(s *server.MCPServer) {
	tool := mcp.NewTool("rules_create",
		mcp.WithDescription("Create a custom rule.<br>Requires the 'Administer Quality Profiles' permission"),
		mcp.WithString("customKey",
			mcp.Description("Key of the custom rule"),
			mcp.Required(),
		),
		mcp.WithString("markdownDescription",
			mcp.Description("Rule description in <a href='/formatting/help'>markdown format</a>"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("Rule name"),
			mcp.Required(),
		),
		mcp.WithString("params",
			mcp.Description("Parameters as semi-colon list of <key>=<value>, for example 'params=key1=v1;key2=v2' (Only for custom rule)"),
		),
		mcp.WithString("preventReactivation",
			mcp.Description("If set to true and if the rule has been deactivated (status 'REMOVED'), a status 409 will be returned"),
		),
		mcp.WithString("severity",
			mcp.Description("Rule severity"),
		),
		mcp.WithString("status",
			mcp.Description("Rule status"),
		),
		mcp.WithString("templateKey",
			mcp.Description("Key of the template rule in order to create a custom rule (mandatory for custom rule)"),
			mcp.Required(),
		),
		mcp.WithString("type",
			mcp.Description("Rule type"),
		),
	)

	s.AddTool(tool, rulesCreateHandler)
}

func rulesCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRulesCreate(request)
	return toResult(c.ApiRulesCreate(ctx, &params, authorizationHeader))
}

func parseRulesCreate(request mcp.CallToolRequest) client.ApiRulesCreateParams {
	params := client.ApiRulesCreateParams{}

	customKey := request.GetString("customKey", "")
	if customKey != "" {
		params.CustomKey = customKey
	}

	markdownDescription := request.GetString("markdownDescription", "")
	if markdownDescription != "" {
		params.MarkdownDescription = markdownDescription
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	_params := request.GetString("params", "")
	if _params != "" {
		params.Params = &_params
	}

	preventReactivation := request.GetString("preventReactivation", "")
	if preventReactivation != "" {
		params.PreventReactivation = &preventReactivation
	}

	severity := request.GetString("severity", "")
	if severity != "" {
		params.Severity = &severity
	}

	status := request.GetString("status", "")
	if status != "" {
		params.Status = &status
	}

	templateKey := request.GetString("templateKey", "")
	if templateKey != "" {
		params.TemplateKey = templateKey
	}

	_type := request.GetString("type", "")
	if _type != "" {
		params.Type = &_type
	}

	return params
}

func registerRulesDelete(s *server.MCPServer) {
	tool := mcp.NewTool("rules_delete",
		mcp.WithDescription("Delete custom rule.<br/>Requires the 'Administer Quality Profiles' permission"),
		mcp.WithString("key",
			mcp.Description("Rule key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, rulesDeleteHandler)
}

func rulesDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRulesDelete(request)
	return toResult(c.ApiRulesDelete(ctx, &params, authorizationHeader))
}

func parseRulesDelete(request mcp.CallToolRequest) client.ApiRulesDeleteParams {
	params := client.ApiRulesDeleteParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}

func registerRulesRepositories(s *server.MCPServer) {
	tool := mcp.NewTool("rules_repositories",
		mcp.WithDescription("List available rule repositories"),
		mcp.WithString("language",
			mcp.Description("A language key; if provided, only repositories for the given language will be returned"),
		),
		mcp.WithString("q",
			mcp.Description("A pattern to match repository keys/names against"),
		),
	)

	s.AddTool(tool, rulesRepositoriesHandler)
}

func rulesRepositoriesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRulesRepositories(request)
	return toResult(c.ApiRulesRepositories(ctx, &params, authorizationHeader))
}

func parseRulesRepositories(request mcp.CallToolRequest) client.ApiRulesRepositoriesParams {
	params := client.ApiRulesRepositoriesParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = &language
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	return params
}

func registerRulesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("rules_search",
		mcp.WithDescription("Search for a collection of relevant rules matching a specified query.<br/>"),
		mcp.WithString("activation",
			mcp.Description("Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set."),
		),
		mcp.WithString("active_severities",
			mcp.Description("Comma-separated list of activation severities, i.e the severity of rules in Quality profiles."),
		),
		mcp.WithString("asc",
			mcp.Description("Ascending sort"),
		),
		mcp.WithString("available_since",
			mcp.Description("Filters rules added since date. Format is yyyy-MM-dd"),
		),
		mcp.WithString("cwe",
			mcp.Description("Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE."),
		),
		mcp.WithString("f",
			mcp.Description("Comma-separated list of additional fields to be returned in the response. All the fields are returned by default, except actives."),
		),
		mcp.WithString("facets",
			mcp.Description("Comma-separated list of the facets to be computed. No facet is computed by default."),
		),
		mcp.WithString("include_external",
			mcp.Description("Include external engine rules in the results"),
		),
		mcp.WithString("inheritance",
			mcp.Description("Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set."),
		),
		mcp.WithString("is_template",
			mcp.Description("Filter template rules"),
		),
		mcp.WithString("languages",
			mcp.Description("Comma-separated list of languages"),
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
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("q",
			mcp.Description("UTF-8 search query"),
		),
		mcp.WithString("qprofile",
			mcp.Description("Quality profile key to filter on. Used only if the parameter 'activation' is set."),
		),
		mcp.WithString("repositories",
			mcp.Description("Comma-separated list of repositories"),
		),
		mcp.WithString("rule_key",
			mcp.Description("Key of rule to search for"),
		),
		mcp.WithString("s",
			mcp.Description("Sort field"),
		),
		mcp.WithString("sansTop25",
			mcp.Description("Comma-separated list of SANS Top 25 categories."),
		),
		mcp.WithString("severities",
			mcp.Description("Comma-separated list of default severities. Not the same than severity of rules in Quality profiles."),
		),
		mcp.WithString("sonarsourceSecurity",
			mcp.Description("Comma-separated list of SonarSource security categories. Use 'others' to select rules not associated with any category"),
		),
		mcp.WithString("statuses",
			mcp.Description("Comma-separated list of status codes"),
		),
		mcp.WithString("tags",
			mcp.Description("Comma-separated list of tags. Returned rules match any of the tags (OR operator)"),
		),
		mcp.WithString("template_key",
			mcp.Description("Key of the template rule to filter on. Used to search for the custom rules based on this template."),
		),
		mcp.WithString("types",
			mcp.Description("Comma-separated list of types. Returned rules match any of the tags (OR operator)"),
		),
	)

	s.AddTool(tool, rulesSearchHandler)
}

func rulesSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRulesSearch(request)
	return toResult(c.ApiRulesSearch(ctx, &params, authorizationHeader))
}

func parseRulesSearch(request mcp.CallToolRequest) client.ApiRulesSearchParams {
	params := client.ApiRulesSearchParams{}

	activation := request.GetString("activation", "")
	if activation != "" {
		params.Activation = &activation
	}

	active_severities := request.GetString("active_severities", "")
	if active_severities != "" {
		params.ActiveSeverities = &active_severities
	}

	asc := request.GetString("asc", "")
	if asc != "" {
		params.Asc = &asc
	}

	available_since := request.GetString("available_since", "")
	if available_since != "" {
		params.AvailableSince = &available_since
	}

	cwe := request.GetString("cwe", "")
	if cwe != "" {
		params.Cwe = &cwe
	}

	f := request.GetString("f", "")
	if f != "" {
		params.F = &f
	}

	facets := request.GetString("facets", "")
	if facets != "" {
		params.Facets = &facets
	}

	include_external := request.GetString("include_external", "")
	if include_external != "" {
		params.IncludeExternal = &include_external
	}

	inheritance := request.GetString("inheritance", "")
	if inheritance != "" {
		params.Inheritance = &inheritance
	}

	is_template := request.GetString("is_template", "")
	if is_template != "" {
		params.IsTemplate = &is_template
	}

	languages := request.GetString("languages", "")
	if languages != "" {
		params.Languages = &languages
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

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	qprofile := request.GetString("qprofile", "")
	if qprofile != "" {
		params.Qprofile = &qprofile
	}

	repositories := request.GetString("repositories", "")
	if repositories != "" {
		params.Repositories = &repositories
	}

	rule_key := request.GetString("rule_key", "")
	if rule_key != "" {
		params.RuleKey = &rule_key
	}

	s := request.GetString("s", "")
	if s != "" {
		params.S = &s
	}

	sansTop25 := request.GetString("sansTop25", "")
	if sansTop25 != "" {
		params.SansTop25 = &sansTop25
	}

	severities := request.GetString("severities", "")
	if severities != "" {
		params.Severities = &severities
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

	template_key := request.GetString("template_key", "")
	if template_key != "" {
		params.TemplateKey = &template_key
	}

	types := request.GetString("types", "")
	if types != "" {
		params.Types = &types
	}

	return params
}

func registerRulesShow(s *server.MCPServer) {
	tool := mcp.NewTool("rules_show",
		mcp.WithDescription("Get detailed information about a rule<br>"),
		mcp.WithString("actives",
			mcp.Description("Show rule's activations for all profiles (\"active rules\")"),
		),
		mcp.WithString("key",
			mcp.Description("Rule key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, rulesShowHandler)
}

func rulesShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRulesShow(request)
	return toResult(c.ApiRulesShow(ctx, &params, authorizationHeader))
}

func parseRulesShow(request mcp.CallToolRequest) client.ApiRulesShowParams {
	params := client.ApiRulesShowParams{}

	actives := request.GetString("actives", "")
	if actives != "" {
		params.Actives = &actives
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}

func registerRulesTags(s *server.MCPServer) {
	tool := mcp.NewTool("rules_tags",
		mcp.WithDescription("List rule tags"),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to tags that contain the supplied string."),
		),
	)

	s.AddTool(tool, rulesTagsHandler)
}

func rulesTagsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRulesTags(request)
	return toResult(c.ApiRulesTags(ctx, &params, authorizationHeader))
}

func parseRulesTags(request mcp.CallToolRequest) client.ApiRulesTagsParams {
	params := client.ApiRulesTagsParams{}

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

func registerRulesUpdate(s *server.MCPServer) {
	tool := mcp.NewTool("rules_update",
		mcp.WithDescription("Update an existing rule.<br>Requires the 'Administer Quality Profiles' permission"),
		mcp.WithString("key",
			mcp.Description("Key of the rule to update"),
			mcp.Required(),
		),
		mcp.WithString("markdown_description",
			mcp.Description("Rule description (mandatory for custom rule and manual rule) in <a href='/formatting/help'>markdown format</a>"),
		),
		mcp.WithString("markdown_note",
			mcp.Description("Optional note in <a href='/formatting/help'>markdown format</a>. Use empty value to remove current note. Note is not changed if the parameter is not set."),
		),
		mcp.WithString("name",
			mcp.Description("Rule name (mandatory for custom rule)"),
		),
		mcp.WithString("params",
			mcp.Description("Parameters as semi-colon list of <key>=<value>, for example 'params=key1=v1;key2=v2' (Only when updating a custom rule)"),
		),
		mcp.WithString("remediation_fn_base_effort",
			mcp.Description("Base effort of the remediation function of the rule"),
		),
		mcp.WithString("remediation_fn_type",
			mcp.Description("Type of the remediation function of the rule"),
		),
		mcp.WithString("remediation_fy_gap_multiplier",
			mcp.Description("Gap multiplier of the remediation function of the rule"),
		),
		mcp.WithString("severity",
			mcp.Description("Rule severity (Only when updating a custom rule)"),
		),
		mcp.WithString("status",
			mcp.Description("Rule status (Only when updating a custom rule)"),
		),
		mcp.WithString("tags",
			mcp.Description("Optional comma-separated list of tags to set. Use blank value to remove current tags. Tags are not changed if the parameter is not set."),
		),
	)

	s.AddTool(tool, rulesUpdateHandler)
}

func rulesUpdateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRulesUpdate(request)
	return toResult(c.ApiRulesUpdate(ctx, &params, authorizationHeader))
}

func parseRulesUpdate(request mcp.CallToolRequest) client.ApiRulesUpdateParams {
	params := client.ApiRulesUpdateParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	markdown_description := request.GetString("markdown_description", "")
	if markdown_description != "" {
		params.MarkdownDescription = &markdown_description
	}

	markdown_note := request.GetString("markdown_note", "")
	if markdown_note != "" {
		params.MarkdownNote = &markdown_note
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	_params := request.GetString("params", "")
	if _params != "" {
		params.Params = &_params
	}

	remediation_fn_base_effort := request.GetString("remediation_fn_base_effort", "")
	if remediation_fn_base_effort != "" {
		params.RemediationFnBaseEffort = &remediation_fn_base_effort
	}

	remediation_fn_type := request.GetString("remediation_fn_type", "")
	if remediation_fn_type != "" {
		params.RemediationFnType = &remediation_fn_type
	}

	remediation_fy_gap_multiplier := request.GetString("remediation_fy_gap_multiplier", "")
	if remediation_fy_gap_multiplier != "" {
		params.RemediationFyGapMultiplier = &remediation_fy_gap_multiplier
	}

	severity := request.GetString("severity", "")
	if severity != "" {
		params.Severity = &severity
	}

	status := request.GetString("status", "")
	if status != "" {
		params.Status = &status
	}

	tags := request.GetString("tags", "")
	if tags != "" {
		params.Tags = &tags
	}

	return params
}
