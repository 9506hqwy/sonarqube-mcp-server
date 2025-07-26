package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerPermissionsAddGroup(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_group",
		mcp.WithDescription("Add a permission to a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name or group id must be provided. <br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithString("groupId",
			mcp.Description("Group id, use 'name' param instead"),
		),
		mcp.WithString("groupName",
			mcp.Description("Group name or 'anyone' (case insensitive)"),
		),
		mcp.WithString("permission",
			mcp.Description("The permission you would like to grant to the group.<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("projectId",
			mcp.Description("Project id"),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key"),
		),
	)

	s.AddTool(tool, permissionsAddGroupHandler)
}

func permissionsAddGroupHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsAddGroup(request)
	return toResult(c.ApiPermissionsAddGroup(ctx, &params, authorizationHeader))
}

func parsePermissionsAddGroup(request mcp.CallToolRequest) client.ApiPermissionsAddGroupParams {
	params := client.ApiPermissionsAddGroupParams{}

	groupId := request.GetString("groupId", "")
	if groupId != "" {
		params.GroupId = &groupId
	}

	groupName := request.GetString("groupName", "")
	if groupName != "" {
		params.GroupName = &groupName
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

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

func registerPermissionsAddGroupToTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_group_to_template",
		mcp.WithDescription("Add a group to a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'."),
		mcp.WithString("groupId",
			mcp.Description("Group id, use 'name' param instead"),
		),
		mcp.WithString("groupName",
			mcp.Description("Group name or 'anyone' (case insensitive)"),
		),
		mcp.WithString("permission",
			mcp.Description("Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsAddGroupToTemplateHandler)
}

func permissionsAddGroupToTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsAddGroupToTemplate(request)
	return toResult(c.ApiPermissionsAddGroupToTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsAddGroupToTemplate(request mcp.CallToolRequest) client.ApiPermissionsAddGroupToTemplateParams {
	params := client.ApiPermissionsAddGroupToTemplateParams{}

	groupId := request.GetString("groupId", "")
	if groupId != "" {
		params.GroupId = &groupId
	}

	groupName := request.GetString("groupName", "")
	if groupName != "" {
		params.GroupName = &groupName
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsAddProjectCreatorToTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_project_creator_to_template",
		mcp.WithDescription("Add a project creator to a permission template.<br>Requires the following permission: 'Administer System'."),
		mcp.WithString("permission",
			mcp.Description("Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsAddProjectCreatorToTemplateHandler)
}

func permissionsAddProjectCreatorToTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsAddProjectCreatorToTemplate(request)
	return toResult(c.ApiPermissionsAddProjectCreatorToTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsAddProjectCreatorToTemplate(request mcp.CallToolRequest) client.ApiPermissionsAddProjectCreatorToTemplateParams {
	params := client.ApiPermissionsAddProjectCreatorToTemplateParams{}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsAddUser(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_user",
		mcp.WithDescription("Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithString("login",
			mcp.Description("User login"),
			mcp.Required(),
		),
		mcp.WithString("permission",
			mcp.Description("The permission you would like to grant to the user<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("projectId",
			mcp.Description("Project id"),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key"),
		),
	)

	s.AddTool(tool, permissionsAddUserHandler)
}

func permissionsAddUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsAddUser(request)
	return toResult(c.ApiPermissionsAddUser(ctx, &params, authorizationHeader))
}

func parsePermissionsAddUser(request mcp.CallToolRequest) client.ApiPermissionsAddUserParams {
	params := client.ApiPermissionsAddUserParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

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

func registerPermissionsAddUserToTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_user_to_template",
		mcp.WithDescription("Add a user to a permission template.<br /> Requires the following permission: 'Administer System'."),
		mcp.WithString("login",
			mcp.Description("User login"),
			mcp.Required(),
		),
		mcp.WithString("permission",
			mcp.Description("Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsAddUserToTemplateHandler)
}

func permissionsAddUserToTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsAddUserToTemplate(request)
	return toResult(c.ApiPermissionsAddUserToTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsAddUserToTemplate(request mcp.CallToolRequest) client.ApiPermissionsAddUserToTemplateParams {
	params := client.ApiPermissionsAddUserToTemplateParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsApplyTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_apply_template",
		mcp.WithDescription("Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the following permission: 'Administer System'."),
		mcp.WithString("projectId",
			mcp.Description("Project id"),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key"),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsApplyTemplateHandler)
}

func permissionsApplyTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsApplyTemplate(request)
	return toResult(c.ApiPermissionsApplyTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsApplyTemplate(request mcp.CallToolRequest) client.ApiPermissionsApplyTemplateParams {
	params := client.ApiPermissionsApplyTemplateParams{}

	projectId := request.GetString("projectId", "")
	if projectId != "" {
		params.ProjectId = &projectId
	}

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = &projectKey
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsBulkApplyTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_bulk_apply_template",
		mcp.WithDescription("Apply a permission template to several projects.<br />The template id or name must be provided.<br />Requires the following permission: 'Administer System'."),
		mcp.WithString("analyzedBefore",
			mcp.Description("Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided."),
		),
		mcp.WithString("onProvisionedOnly",
			mcp.Description("Filter the projects that are provisioned"),
		),
		mcp.WithString("projects",
			mcp.Description("Comma-separated list of project keys"),
		),
		mcp.WithString("q",
			mcp.Description("Limit search to: <ul><li>project names that contain the supplied string</li><li>project keys that are exactly the same as the supplied string</li></ul>"),
		),
		mcp.WithString("qualifiers",
			mcp.Description("Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>TRK - Projects</li></ul>"),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsBulkApplyTemplateHandler)
}

func permissionsBulkApplyTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsBulkApplyTemplate(request)
	return toResult(c.ApiPermissionsBulkApplyTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsBulkApplyTemplate(request mcp.CallToolRequest) client.ApiPermissionsBulkApplyTemplateParams {
	params := client.ApiPermissionsBulkApplyTemplateParams{}

	analyzedBefore := request.GetString("analyzedBefore", "")
	if analyzedBefore != "" {
		params.AnalyzedBefore = &analyzedBefore
	}

	onProvisionedOnly := request.GetString("onProvisionedOnly", "")
	if onProvisionedOnly != "" {
		params.OnProvisionedOnly = &onProvisionedOnly
	}

	projects := request.GetString("projects", "")
	if projects != "" {
		params.Projects = &projects
	}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	qualifiers := request.GetString("qualifiers", "")
	if qualifiers != "" {
		params.Qualifiers = &qualifiers
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsCreateTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_create_template",
		mcp.WithDescription("Create a permission template.<br />Requires the following permission: 'Administer System'."),
		mcp.WithString("description",
			mcp.Description("Description"),
		),
		mcp.WithString("name",
			mcp.Description("Name"),
			mcp.Required(),
		),
		mcp.WithString("projectKeyPattern",
			mcp.Description("Project key pattern. Must be a valid Java regular expression"),
		),
	)

	s.AddTool(tool, permissionsCreateTemplateHandler)
}

func permissionsCreateTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsCreateTemplate(request)
	return toResult(c.ApiPermissionsCreateTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsCreateTemplate(request mcp.CallToolRequest) client.ApiPermissionsCreateTemplateParams {
	params := client.ApiPermissionsCreateTemplateParams{}

	description := request.GetString("description", "")
	if description != "" {
		params.Description = &description
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = name
	}

	projectKeyPattern := request.GetString("projectKeyPattern", "")
	if projectKeyPattern != "" {
		params.ProjectKeyPattern = &projectKeyPattern
	}

	return params
}

func registerPermissionsDeleteTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_delete_template",
		mcp.WithDescription("Delete a permission template.<br />Requires the following permission: 'Administer System'."),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsDeleteTemplateHandler)
}

func permissionsDeleteTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsDeleteTemplate(request)
	return toResult(c.ApiPermissionsDeleteTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsDeleteTemplate(request mcp.CallToolRequest) client.ApiPermissionsDeleteTemplateParams {
	params := client.ApiPermissionsDeleteTemplateParams{}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsRemoveGroup(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_group",
		mcp.WithDescription("Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group id or group name must be provided, not both.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithString("groupId",
			mcp.Description("Group id, use 'name' param instead"),
		),
		mcp.WithString("groupName",
			mcp.Description("Group name or 'anyone' (case insensitive)"),
		),
		mcp.WithString("permission",
			mcp.Description("The permission you would like to revoke from the group.<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("projectId",
			mcp.Description("Project id"),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key"),
		),
	)

	s.AddTool(tool, permissionsRemoveGroupHandler)
}

func permissionsRemoveGroupHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsRemoveGroup(request)
	return toResult(c.ApiPermissionsRemoveGroup(ctx, &params, authorizationHeader))
}

func parsePermissionsRemoveGroup(request mcp.CallToolRequest) client.ApiPermissionsRemoveGroupParams {
	params := client.ApiPermissionsRemoveGroupParams{}

	groupId := request.GetString("groupId", "")
	if groupId != "" {
		params.GroupId = &groupId
	}

	groupName := request.GetString("groupName", "")
	if groupName != "" {
		params.GroupName = &groupName
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

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

func registerPermissionsRemoveGroupFromTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_group_from_template",
		mcp.WithDescription("Remove a group from a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'."),
		mcp.WithString("groupId",
			mcp.Description("Group id, use 'name' param instead"),
		),
		mcp.WithString("groupName",
			mcp.Description("Group name or 'anyone' (case insensitive)"),
		),
		mcp.WithString("permission",
			mcp.Description("Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsRemoveGroupFromTemplateHandler)
}

func permissionsRemoveGroupFromTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsRemoveGroupFromTemplate(request)
	return toResult(c.ApiPermissionsRemoveGroupFromTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsRemoveGroupFromTemplate(request mcp.CallToolRequest) client.ApiPermissionsRemoveGroupFromTemplateParams {
	params := client.ApiPermissionsRemoveGroupFromTemplateParams{}

	groupId := request.GetString("groupId", "")
	if groupId != "" {
		params.GroupId = &groupId
	}

	groupName := request.GetString("groupName", "")
	if groupName != "" {
		params.GroupName = &groupName
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsRemoveProjectCreatorFromTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_project_creator_from_template",
		mcp.WithDescription("Remove a project creator from a permission template.<br>Requires the following permission: 'Administer System'."),
		mcp.WithString("permission",
			mcp.Description("Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsRemoveProjectCreatorFromTemplateHandler)
}

func permissionsRemoveProjectCreatorFromTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsRemoveProjectCreatorFromTemplate(request)
	return toResult(c.ApiPermissionsRemoveProjectCreatorFromTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsRemoveProjectCreatorFromTemplate(request mcp.CallToolRequest) client.ApiPermissionsRemoveProjectCreatorFromTemplateParams {
	params := client.ApiPermissionsRemoveProjectCreatorFromTemplateParams{}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsRemoveUser(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_user",
		mcp.WithDescription("Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithString("login",
			mcp.Description("User login"),
			mcp.Required(),
		),
		mcp.WithString("permission",
			mcp.Description("The permission you would like to revoke from the user.<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("projectId",
			mcp.Description("Project id"),
		),
		mcp.WithString("projectKey",
			mcp.Description("Project key"),
		),
	)

	s.AddTool(tool, permissionsRemoveUserHandler)
}

func permissionsRemoveUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsRemoveUser(request)
	return toResult(c.ApiPermissionsRemoveUser(ctx, &params, authorizationHeader))
}

func parsePermissionsRemoveUser(request mcp.CallToolRequest) client.ApiPermissionsRemoveUserParams {
	params := client.ApiPermissionsRemoveUserParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

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

func registerPermissionsRemoveUserFromTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_user_from_template",
		mcp.WithDescription("Remove a user from a permission template.<br /> Requires the following permission: 'Administer System'."),
		mcp.WithString("login",
			mcp.Description("User login"),
			mcp.Required(),
		),
		mcp.WithString("permission",
			mcp.Description("Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>"),
			mcp.Required(),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsRemoveUserFromTemplateHandler)
}

func permissionsRemoveUserFromTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsRemoveUserFromTemplate(request)
	return toResult(c.ApiPermissionsRemoveUserFromTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsRemoveUserFromTemplate(request mcp.CallToolRequest) client.ApiPermissionsRemoveUserFromTemplateParams {
	params := client.ApiPermissionsRemoveUserFromTemplateParams{}

	login := request.GetString("login", "")
	if login != "" {
		params.Login = login
	}

	permission := request.GetString("permission", "")
	if permission != "" {
		params.Permission = permission
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsSearchTemplates(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_search_templates",
		mcp.WithDescription("List permission templates.<br />Requires the following permission: 'Administer System'."),
		mcp.WithString("q",
			mcp.Description("Limit search to permission template names that contain the supplied string."),
		),
	)

	s.AddTool(tool, permissionsSearchTemplatesHandler)
}

func permissionsSearchTemplatesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsSearchTemplates(request)
	return toResult(c.ApiPermissionsSearchTemplates(ctx, &params, authorizationHeader))
}

func parsePermissionsSearchTemplates(request mcp.CallToolRequest) client.ApiPermissionsSearchTemplatesParams {
	params := client.ApiPermissionsSearchTemplatesParams{}

	q := request.GetString("q", "")
	if q != "" {
		params.Q = &q
	}

	return params
}

func registerPermissionsSetDefaultTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_set_default_template",
		mcp.WithDescription("Set a permission template as default.<br />Requires the following permission: 'Administer System'."),
		mcp.WithString("qualifier",
			mcp.Description("Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>TRK - Projects</li></ul>"),
		),
		mcp.WithString("templateId",
			mcp.Description("Template id"),
		),
		mcp.WithString("templateName",
			mcp.Description("Template name"),
		),
	)

	s.AddTool(tool, permissionsSetDefaultTemplateHandler)
}

func permissionsSetDefaultTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsSetDefaultTemplate(request)
	return toResult(c.ApiPermissionsSetDefaultTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsSetDefaultTemplate(request mcp.CallToolRequest) client.ApiPermissionsSetDefaultTemplateParams {
	params := client.ApiPermissionsSetDefaultTemplateParams{}

	qualifier := request.GetString("qualifier", "")
	if qualifier != "" {
		params.Qualifier = &qualifier
	}

	templateId := request.GetString("templateId", "")
	if templateId != "" {
		params.TemplateId = &templateId
	}

	templateName := request.GetString("templateName", "")
	if templateName != "" {
		params.TemplateName = &templateName
	}

	return params
}

func registerPermissionsUpdateTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_update_template",
		mcp.WithDescription("Update a permission template.<br />Requires the following permission: 'Administer System'."),
		mcp.WithString("description",
			mcp.Description("Description"),
		),
		mcp.WithString("id",
			mcp.Description("Id"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("Name"),
		),
		mcp.WithString("projectKeyPattern",
			mcp.Description("Project key pattern. Must be a valid Java regular expression"),
		),
	)

	s.AddTool(tool, permissionsUpdateTemplateHandler)
}

func permissionsUpdateTemplateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parsePermissionsUpdateTemplate(request)
	return toResult(c.ApiPermissionsUpdateTemplate(ctx, &params, authorizationHeader))
}

func parsePermissionsUpdateTemplate(request mcp.CallToolRequest) client.ApiPermissionsUpdateTemplateParams {
	params := client.ApiPermissionsUpdateTemplateParams{}

	description := request.GetString("description", "")
	if description != "" {
		params.Description = &description
	}

	id := request.GetString("id", "")
	if id != "" {
		params.Id = id
	}

	name := request.GetString("name", "")
	if name != "" {
		params.Name = &name
	}

	projectKeyPattern := request.GetString("projectKeyPattern", "")
	if projectKeyPattern != "" {
		params.ProjectKeyPattern = &projectKeyPattern
	}

	return params
}
