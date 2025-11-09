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
		mcp.WithInputSchema[client.ApiQualityprofilesActivateRuleParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesActivateRuleHandler))
}

func qualityprofilesActivateRuleHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesActivateRuleParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesActivateRule(ctx, &params, authorizationHeader))
}

func registerQualityprofilesActivateRules(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_activate_rules",
		mcp.WithDescription("Bulk-activate rules on one quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesActivateRulesParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesActivateRulesHandler))
}

func qualityprofilesActivateRulesHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesActivateRulesParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesActivateRules(ctx, &params, authorizationHeader))
}

func registerQualityprofilesAddProject(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_add_project",
		mcp.WithDescription("Associate a project with a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li> <li>Administer right on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesAddProjectParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesAddProjectHandler))
}

func qualityprofilesAddProjectHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesAddProjectParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesAddProject(ctx, &params, authorizationHeader))
}

func registerQualityprofilesBackup(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_backup",
		mcp.WithDescription("Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore."),
		mcp.WithInputSchema[client.ApiQualityprofilesBackupParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesBackupHandler))
}

func qualityprofilesBackupHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesBackupParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesBackup(ctx, &params, authorizationHeader))
}

func registerQualityprofilesChangeParent(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_change_parent",
		mcp.WithDescription("Change a quality profile's parent.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesChangeParentParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesChangeParentHandler))
}

func qualityprofilesChangeParentHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesChangeParentParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesChangeParent(ctx, &params, authorizationHeader))
}

func registerQualityprofilesChangelog(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_changelog",
		mcp.WithDescription("Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first)."),
		mcp.WithInputSchema[client.ApiQualityprofilesChangelogParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesChangelogHandler))
}

func qualityprofilesChangelogHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesChangelogParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesChangelog(ctx, &params, authorizationHeader))
}

func registerQualityprofilesCopy(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_copy",
		mcp.WithDescription("Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithInputSchema[client.ApiQualityprofilesCopyParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesCopyHandler))
}

func qualityprofilesCopyHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesCopyParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesCopy(ctx, &params, authorizationHeader))
}

func registerQualityprofilesCreate(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_create",
		mcp.WithDescription("Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithInputSchema[client.ApiQualityprofilesCreateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesCreateHandler))
}

func qualityprofilesCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesCreate(ctx, &params, authorizationHeader))
}

func registerQualityprofilesDeactivateRule(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_deactivate_rule",
		mcp.WithDescription("Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesDeactivateRuleParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesDeactivateRuleHandler))
}

func qualityprofilesDeactivateRuleHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesDeactivateRuleParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesDeactivateRule(ctx, &params, authorizationHeader))
}

func registerQualityprofilesDeactivateRules(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_deactivate_rules",
		mcp.WithDescription("Bulk deactivate rules on Quality profiles.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesDeactivateRulesParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesDeactivateRulesHandler))
}

func qualityprofilesDeactivateRulesHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesDeactivateRulesParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesDeactivateRules(ctx, &params, authorizationHeader))
}

func registerQualityprofilesDelete(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_delete",
		mcp.WithDescription("Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesDeleteParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesDeleteHandler))
}

func qualityprofilesDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesDelete(ctx, &params, authorizationHeader))
}

func registerQualityprofilesExport(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_export",
		mcp.WithDescription("Export a quality profile."),
		mcp.WithInputSchema[client.ApiQualityprofilesExportParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesExportHandler))
}

func qualityprofilesExportHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesExportParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesExport(ctx, &params, authorizationHeader))
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
		mcp.WithInputSchema[client.ApiQualityprofilesInheritanceParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesInheritanceHandler))
}

func qualityprofilesInheritanceHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesInheritanceParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesInheritance(ctx, &params, authorizationHeader))
}

func registerQualityprofilesProjects(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_projects",
		mcp.WithDescription("List projects with their association status regarding a quality profile <br/>See api/qualityprofiles/search in order to get the Quality Profile Key"),
		mcp.WithInputSchema[client.ApiQualityprofilesProjectsParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesProjectsHandler))
}

func qualityprofilesProjectsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesProjectsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesProjects(ctx, &params, authorizationHeader))
}

func registerQualityprofilesRemoveProject(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_remove_project",
		mcp.WithDescription("Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li> <li>Administer right on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesRemoveProjectParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesRemoveProjectHandler))
}

func qualityprofilesRemoveProjectHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesRemoveProjectParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesRemoveProject(ctx, &params, authorizationHeader))
}

func registerQualityprofilesRename(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_rename",
		mcp.WithDescription("Rename a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>"),
		mcp.WithInputSchema[client.ApiQualityprofilesRenameParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesRenameHandler))
}

func qualityprofilesRenameHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesRenameParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesRename(ctx, &params, authorizationHeader))
}

func registerQualityprofilesRestore(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_restore",
		mcp.WithDescription("Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithInputSchema[client.ApiQualityprofilesRestoreParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesRestoreHandler))
}

func qualityprofilesRestoreHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesRestoreParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesRestore(ctx, &params, authorizationHeader))
}

func registerQualityprofilesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_search",
		mcp.WithDescription("Search quality profiles"),
		mcp.WithInputSchema[client.ApiQualityprofilesSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesSearchHandler))
}

func qualityprofilesSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesSearch(ctx, &params, authorizationHeader))
}

func registerQualityprofilesSetDefault(s *server.MCPServer) {
	tool := mcp.NewTool("qualityprofiles_set_default",
		mcp.WithDescription("Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission."),
		mcp.WithInputSchema[client.ApiQualityprofilesSetDefaultParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(qualityprofilesSetDefaultHandler))
}

func qualityprofilesSetDefaultHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiQualityprofilesSetDefaultParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiQualityprofilesSetDefault(ctx, &params, authorizationHeader))
}
