package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerQualityprofilesActivateRule(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_activate_rule",
		mcp.WithDescription("Activate a rule on a Quality Profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithString("key",
			mcp.Description("Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>"),
			mcp.Required(),
		),
		mcp.WithString("params",
			mcp.Description("Parameters as semi-colon list of <code>key=value</code>. Ignored if parameter reset is true."),
		),
		mcp.WithString("reset",
			mcp.Description("Reset severity and parameters of activated rule. Set the values defined on parent profile or from rule default values."),
		),
		mcp.WithString("rule",
			mcp.Description("Rule key"),
			mcp.Required(),
		),
		mcp.WithString("severity",
			mcp.Description("Severity. Ignored if parameter reset is true."),
		),
	)

	s.AddTool(tool, qualityprofilesActivateRuleHandler)
}

func qualityprofilesActivateRuleHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesActivateRule(request)
	return toResult(c.ApiQualityprofilesActivateRule(ctx, &params, authorizationHeader))
}

func parseQualityprofilesActivateRule(request mcp.CallToolRequest) client.ApiQualityprofilesActivateRuleParams {
	params := client.ApiQualityprofilesActivateRuleParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	_params := request.GetString("params", "")
	if _params != "" {
		params.Params = &_params
	}

	reset := request.GetString("reset", "")
	if reset != "" {
		params.Reset = &reset
	}

	rule := request.GetString("rule", "")
	if rule != "" {
		params.Rule = rule
	}

	severity := request.GetString("severity", "")
	if severity != "" {
		params.Severity = &severity
	}

	return params
}

func registerQualityprofilesActivateRules(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_activate_rules",
		mcp.WithDescription("Bulk-activate rules on one quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
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
		mcp.WithString("targetKey",
			mcp.Description("Quality Profile key on which the rule activation is done. To retrieve a quality profile key please see <code>api/qualityprofiles/search</code>"),
			mcp.Required(),
		),
		mcp.WithString("targetSeverity",
			mcp.Description("Severity to set on the activated rules"),
		),
		mcp.WithString("template_key",
			mcp.Description("Key of the template rule to filter on. Used to search for the custom rules based on this template."),
		),
		mcp.WithString("types",
			mcp.Description("Comma-separated list of types. Returned rules match any of the tags (OR operator)"),
		),
	)

	s.AddTool(tool, qualityprofilesActivateRulesHandler)
}

func qualityprofilesActivateRulesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesActivateRules(request)
	return toResult(c.ApiQualityprofilesActivateRules(ctx, &params, authorizationHeader))
}

func parseQualityprofilesActivateRules(request mcp.CallToolRequest) client.ApiQualityprofilesActivateRulesParams {
	params := client.ApiQualityprofilesActivateRulesParams{}

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

	targetKey := request.GetString("targetKey", "")
	if targetKey != "" {
		params.TargetKey = targetKey
	}

	targetSeverity := request.GetString("targetSeverity", "")
	if targetSeverity != "" {
		params.TargetSeverity = &targetSeverity
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

func registerQualityprofilesAddProject(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_add_project",
		mcp.WithDescription("Associate a project with a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li> <li>Administer right on the specified project</li></ul>"),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("Project key"),
			mcp.Required(),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesAddProjectHandler)
}

func qualityprofilesAddProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesAddProject(request)
	return toResult(c.ApiQualityprofilesAddProject(ctx, &params, authorizationHeader))
}

func parseQualityprofilesAddProject(request mcp.CallToolRequest) client.ApiQualityprofilesAddProjectParams {
	params := client.ApiQualityprofilesAddProjectParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	return params
}

func registerQualityprofilesBackup(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_backup",
		mcp.WithDescription("Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore."),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesBackupHandler)
}

func qualityprofilesBackupHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesBackup(request)
	return toResult(c.ApiQualityprofilesBackup(ctx, &params, authorizationHeader))
}

func parseQualityprofilesBackup(request mcp.CallToolRequest) client.ApiQualityprofilesBackupParams {
	params := client.ApiQualityprofilesBackupParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	return params
}

func registerQualityprofilesChangeParent(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_change_parent",
		mcp.WithDescription("Change a quality profile's parent.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("parentQualityProfile",
			mcp.Description("New parent profile name. <br> If no profile is provided, the inheritance link with current parent profile (if any) is broken, which deactivates all rules which come from the parent and are not overridden."),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesChangeParentHandler)
}

func qualityprofilesChangeParentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesChangeParent(request)
	return toResult(c.ApiQualityprofilesChangeParent(ctx, &params, authorizationHeader))
}

func parseQualityprofilesChangeParent(request mcp.CallToolRequest) client.ApiQualityprofilesChangeParentParams {
	params := client.ApiQualityprofilesChangeParentParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	parentQualityProfile := request.GetString("parentQualityProfile", "")
	if parentQualityProfile != "" {
		params.ParentQualityProfile = &parentQualityProfile
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	return params
}

func registerQualityprofilesChangelog(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_changelog",
		mcp.WithDescription("Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first)."),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
		mcp.WithString("since",
			mcp.Description("Start date for the changelog (inclusive). <br>Either a date (server timezone) or datetime can be provided."),
		),
		mcp.WithString("to",
			mcp.Description("End date for the changelog (exclusive, strictly before). <br>Either a date (server timezone) or datetime can be provided."),
		),
	)

	s.AddTool(tool, qualityprofilesChangelogHandler)
}

func qualityprofilesChangelogHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesChangelog(request)
	return toResult(c.ApiQualityprofilesChangelog(ctx, &params, authorizationHeader))
}

func parseQualityprofilesChangelog(request mcp.CallToolRequest) client.ApiQualityprofilesChangelogParams {
	params := client.ApiQualityprofilesChangelogParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	since := request.GetString("since", "")
	if since != "" {
		params.Since = &since
	}

	to := request.GetString("to", "")
	if to != "" {
		params.To = &to
	}

	return params
}

func registerQualityprofilesCopy(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_copy",
		mcp.WithDescription("Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithString("fromKey",
			mcp.Description("Quality profile key"),
			mcp.Required(),
		),
		mcp.WithString("toName",
			mcp.Description("Name for the new quality profile."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesCopyHandler)
}

func qualityprofilesCopyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesCopy(request)
	return toResult(c.ApiQualityprofilesCopy(ctx, &params, authorizationHeader))
}

func parseQualityprofilesCopy(request mcp.CallToolRequest) client.ApiQualityprofilesCopyParams {
	params := client.ApiQualityprofilesCopyParams{}

	fromKey := request.GetString("fromKey", "")
	if fromKey != "" {
		params.FromKey = fromKey
	}

	toName := request.GetString("toName", "")
	if toName != "" {
		params.ToName = toName
	}

	return params
}

func registerQualityprofilesCreate(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_create",
		mcp.WithDescription("Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithString("language",
			mcp.Description("Quality profile language"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("Quality profile name"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesCreateHandler)
}

func qualityprofilesCreateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesCreate(request)
	return toResult(c.ApiQualityprofilesCreate(ctx, &params, authorizationHeader))
}

func parseQualityprofilesCreate(request mcp.CallToolRequest) client.ApiQualityprofilesCreateParams {
	params := client.ApiQualityprofilesCreateParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}

func registerQualityprofilesDeactivateRule(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_deactivate_rule",
		mcp.WithDescription("Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithString("key",
			mcp.Description("Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>"),
			mcp.Required(),
		),
		mcp.WithString("rule",
			mcp.Description("Rule key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesDeactivateRuleHandler)
}

func qualityprofilesDeactivateRuleHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesDeactivateRule(request)
	return toResult(c.ApiQualityprofilesDeactivateRule(ctx, &params, authorizationHeader))
}

func parseQualityprofilesDeactivateRule(request mcp.CallToolRequest) client.ApiQualityprofilesDeactivateRuleParams {
	params := client.ApiQualityprofilesDeactivateRuleParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	rule := request.GetString("rule", "")
	if rule != "" {
		params.Rule = rule
	}

	return params
}

func registerQualityprofilesDeactivateRules(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_deactivate_rules",
		mcp.WithDescription("Bulk deactivate rules on Quality profiles.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
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
		mcp.WithString("targetKey",
			mcp.Description("Quality Profile key on which the rule deactivation is done. To retrieve a profile key please see <code>api/qualityprofiles/search</code>"),
			mcp.Required(),
		),
		mcp.WithString("template_key",
			mcp.Description("Key of the template rule to filter on. Used to search for the custom rules based on this template."),
		),
		mcp.WithString("types",
			mcp.Description("Comma-separated list of types. Returned rules match any of the tags (OR operator)"),
		),
	)

	s.AddTool(tool, qualityprofilesDeactivateRulesHandler)
}

func qualityprofilesDeactivateRulesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesDeactivateRules(request)
	return toResult(c.ApiQualityprofilesDeactivateRules(ctx, &params, authorizationHeader))
}

func parseQualityprofilesDeactivateRules(request mcp.CallToolRequest) client.ApiQualityprofilesDeactivateRulesParams {
	params := client.ApiQualityprofilesDeactivateRulesParams{}

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

	targetKey := request.GetString("targetKey", "")
	if targetKey != "" {
		params.TargetKey = targetKey
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

func registerQualityprofilesDelete(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_delete",
		mcp.WithDescription("Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesDeleteHandler)
}

func qualityprofilesDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesDelete(request)
	return toResult(c.ApiQualityprofilesDelete(ctx, &params, authorizationHeader))
}

func parseQualityprofilesDelete(request mcp.CallToolRequest) client.ApiQualityprofilesDeleteParams {
	params := client.ApiQualityprofilesDeleteParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	return params
}

func registerQualityprofilesExport(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_export",
		mcp.WithDescription("Export a quality profile."),
		mcp.WithString("exporterKey",
			mcp.Description("Output format. If left empty, the same format as api/qualityprofiles/backup is used. Possible values are described by api/qualityprofiles/exporters."),
		),
		mcp.WithString("language",
			mcp.Description("Quality profile language"),
			mcp.Required(),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name to export. If left empty, the default profile for the language is exported."),
		),
	)

	s.AddTool(tool, qualityprofilesExportHandler)
}

func qualityprofilesExportHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesExport(request)
	return toResult(c.ApiQualityprofilesExport(ctx, &params, authorizationHeader))
}

func parseQualityprofilesExport(request mcp.CallToolRequest) client.ApiQualityprofilesExportParams {
	params := client.ApiQualityprofilesExportParams{}

	exporterKey := request.GetString("exporterKey", "")
	if exporterKey != "" {
		params.ExporterKey = &exporterKey
	}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = &qualityProfile
	}

	return params
}

func registerQualityprofilesExporters(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_exporters",
		mcp.WithDescription("Lists available profile export formats."),
	)

	s.AddTool(tool, qualityprofilesExportersHandler)
}

func qualityprofilesExportersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesExporters(ctx, authorizationHeader))
}

func registerQualityprofilesImporters(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_importers",
		mcp.WithDescription("List supported importers."),
	)

	s.AddTool(tool, qualityprofilesImportersHandler)
}

func qualityprofilesImportersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesImporters(ctx, authorizationHeader))
}

func registerQualityprofilesInheritance(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_inheritance",
		mcp.WithDescription("Show a quality profile's ancestors and children."),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesInheritanceHandler)
}

func qualityprofilesInheritanceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesInheritance(request)
	return toResult(c.ApiQualityprofilesInheritance(ctx, &params, authorizationHeader))
}

func parseQualityprofilesInheritance(request mcp.CallToolRequest) client.ApiQualityprofilesInheritanceParams {
	params := client.ApiQualityprofilesInheritanceParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	return params
}

func registerQualityprofilesProjects(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_projects",
		mcp.WithDescription("List projects with their association status regarding a quality profile <br/>See api/qualityprofiles/search in order to get the Quality Profile Key"),
		mcp.WithString("key",
			mcp.Description("Quality profile key"),
			mcp.Required(),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to projects that contain the supplied string."),
		),
		mcp.WithString("selected",
			mcp.Description("Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all)."),
		),
	)

	s.AddTool(tool, qualityprofilesProjectsHandler)
}

func qualityprofilesProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesProjects(request)
	return toResult(c.ApiQualityprofilesProjects(ctx, &params, authorizationHeader))
}

func parseQualityprofilesProjects(request mcp.CallToolRequest) client.ApiQualityprofilesProjectsParams {
	params := client.ApiQualityprofilesProjectsParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
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

	selected := request.GetString("selected", "")
	if selected != "" {
		params.Selected = &selected
	}

	return params
}

func registerQualityprofilesRemoveProject(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_remove_project",
		mcp.WithDescription("Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li> <li>Administer right on the specified project</li></ul>"),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("project",
			mcp.Description("Project key"),
			mcp.Required(),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesRemoveProjectHandler)
}

func qualityprofilesRemoveProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesRemoveProject(request)
	return toResult(c.ApiQualityprofilesRemoveProject(ctx, &params, authorizationHeader))
}

func parseQualityprofilesRemoveProject(request mcp.CallToolRequest) client.ApiQualityprofilesRemoveProjectParams {
	params := client.ApiQualityprofilesRemoveProjectParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	return params
}

func registerQualityprofilesRename(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_rename",
		mcp.WithDescription("Rename a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithString("key",
			mcp.Description("Quality profile key"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("New quality profile name"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesRenameHandler)
}

func qualityprofilesRenameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesRename(request)
	return toResult(c.ApiQualityprofilesRename(ctx, &params, authorizationHeader))
}

func parseQualityprofilesRename(request mcp.CallToolRequest) client.ApiQualityprofilesRenameParams {
	params := client.ApiQualityprofilesRenameParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	return params
}

func registerQualityprofilesRestore(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_restore",
		mcp.WithDescription("Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithString("backup",
			mcp.Description("A profile backup file in XML format, as generated by api/qualityprofiles/backup or the former api/profiles/backup."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesRestoreHandler)
}

func qualityprofilesRestoreHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesRestore(request)
	return toResult(c.ApiQualityprofilesRestore(ctx, &params, authorizationHeader))
}

func parseQualityprofilesRestore(request mcp.CallToolRequest) client.ApiQualityprofilesRestoreParams {
	params := client.ApiQualityprofilesRestoreParams{}

	backup := request.GetString("backup", "")
	if backup != "" {
		params.Backup = backup
	}

	return params
}

func registerQualityprofilesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_search",
		mcp.WithDescription("Search quality profiles"),
		mcp.WithString("defaults",
			mcp.Description("If set to true, return only the quality profiles marked as default for each language"),
		),
		mcp.WithString("language",
			mcp.Description("Language key. If provided, only profiles for the given language are returned."),
		),
		mcp.WithString("project",
			mcp.Description("Project key"),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name"),
		),
	)

	s.AddTool(tool, qualityprofilesSearchHandler)
}

func qualityprofilesSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesSearch(request)
	return toResult(c.ApiQualityprofilesSearch(ctx, &params, authorizationHeader))
}

func parseQualityprofilesSearch(request mcp.CallToolRequest) client.ApiQualityprofilesSearchParams {
	params := client.ApiQualityprofilesSearchParams{}

	defaults := request.GetString("defaults", "")
	if defaults != "" {
		params.Defaults = &defaults
	}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = &language
	}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = &qualityProfile
	}

	return params
}

func registerQualityprofilesSetDefault(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_set_default",
		mcp.WithDescription("Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithString("language",
			mcp.Description("Quality profile language."),
			mcp.Required(),
		),
		mcp.WithString("qualityProfile",
			mcp.Description("Quality profile name."),
			mcp.Required(),
		),
	)

	s.AddTool(tool, qualityprofilesSetDefaultHandler)
}

func qualityprofilesSetDefaultHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseQualityprofilesSetDefault(request)
	return toResult(c.ApiQualityprofilesSetDefault(ctx, &params, authorizationHeader))
}

func parseQualityprofilesSetDefault(request mcp.CallToolRequest) client.ApiQualityprofilesSetDefaultParams {
	params := client.ApiQualityprofilesSetDefaultParams{}

	language := request.GetString("language", "")
	if language != "" {
		params.Language = language
	}

	qualityProfile := request.GetString("qualityProfile", "")
	if qualityProfile != "" {
		params.QualityProfile = qualityProfile
	}

	return params
}
