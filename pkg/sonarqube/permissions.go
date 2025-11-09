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
		mcp.WithInputSchema[client.ApiPermissionsAddGroupParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsAddGroupHandler))
}

func permissionsAddGroupHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsAddGroupParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsAddGroup(ctx, &params, authorizationHeader))
}

func registerPermissionsAddGroupToTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_group_to_template",
		mcp.WithDescription("Add a group to a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsAddGroupToTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsAddGroupToTemplateHandler))
}

func permissionsAddGroupToTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsAddGroupToTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsAddGroupToTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsAddProjectCreatorToTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_project_creator_to_template",
		mcp.WithDescription("Add a project creator to a permission template.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsAddProjectCreatorToTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsAddProjectCreatorToTemplateHandler))
}

func permissionsAddProjectCreatorToTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsAddProjectCreatorToTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsAddProjectCreatorToTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsAddUser(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_user",
		mcp.WithDescription("Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiPermissionsAddUserParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsAddUserHandler))
}

func permissionsAddUserHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsAddUserParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsAddUser(ctx, &params, authorizationHeader))
}

func registerPermissionsAddUserToTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_add_user_to_template",
		mcp.WithDescription("Add a user to a permission template.<br /> Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsAddUserToTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsAddUserToTemplateHandler))
}

func permissionsAddUserToTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsAddUserToTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsAddUserToTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsApplyTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_apply_template",
		mcp.WithDescription("Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsApplyTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsApplyTemplateHandler))
}

func permissionsApplyTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsApplyTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsApplyTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsBulkApplyTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_bulk_apply_template",
		mcp.WithDescription("Apply a permission template to several projects.<br />The template id or name must be provided.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsBulkApplyTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsBulkApplyTemplateHandler))
}

func permissionsBulkApplyTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsBulkApplyTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsBulkApplyTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsCreateTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_create_template",
		mcp.WithDescription("Create a permission template.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsCreateTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsCreateTemplateHandler))
}

func permissionsCreateTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsCreateTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsCreateTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsDeleteTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_delete_template",
		mcp.WithDescription("Delete a permission template.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsDeleteTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsDeleteTemplateHandler))
}

func permissionsDeleteTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsDeleteTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsDeleteTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsRemoveGroup(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_group",
		mcp.WithDescription("Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group id or group name must be provided, not both.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiPermissionsRemoveGroupParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsRemoveGroupHandler))
}

func permissionsRemoveGroupHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsRemoveGroupParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsRemoveGroup(ctx, &params, authorizationHeader))
}

func registerPermissionsRemoveGroupFromTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_group_from_template",
		mcp.WithDescription("Remove a group from a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsRemoveGroupFromTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsRemoveGroupFromTemplateHandler))
}

func permissionsRemoveGroupFromTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsRemoveGroupFromTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsRemoveGroupFromTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsRemoveProjectCreatorFromTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_project_creator_from_template",
		mcp.WithDescription("Remove a project creator from a permission template.<br>Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsRemoveProjectCreatorFromTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsRemoveProjectCreatorFromTemplateHandler))
}

func permissionsRemoveProjectCreatorFromTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsRemoveProjectCreatorFromTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsRemoveProjectCreatorFromTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsRemoveUser(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_user",
		mcp.WithDescription("Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>"),
		mcp.WithInputSchema[client.ApiPermissionsRemoveUserParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsRemoveUserHandler))
}

func permissionsRemoveUserHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsRemoveUserParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsRemoveUser(ctx, &params, authorizationHeader))
}

func registerPermissionsRemoveUserFromTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_remove_user_from_template",
		mcp.WithDescription("Remove a user from a permission template.<br /> Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsRemoveUserFromTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsRemoveUserFromTemplateHandler))
}

func permissionsRemoveUserFromTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsRemoveUserFromTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsRemoveUserFromTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsSearchTemplates(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_search_templates",
		mcp.WithDescription("List permission templates.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsSearchTemplatesParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsSearchTemplatesHandler))
}

func permissionsSearchTemplatesHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsSearchTemplatesParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsSearchTemplates(ctx, &params, authorizationHeader))
}

func registerPermissionsSetDefaultTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_set_default_template",
		mcp.WithDescription("Set a permission template as default.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsSetDefaultTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsSetDefaultTemplateHandler))
}

func permissionsSetDefaultTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsSetDefaultTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsSetDefaultTemplate(ctx, &params, authorizationHeader))
}

func registerPermissionsUpdateTemplate(s *server.MCPServer) {
	tool := mcp.NewTool("permissions_update_template",
		mcp.WithDescription("Update a permission template.<br />Requires the following permission: 'Administer System'."),
		mcp.WithInputSchema[client.ApiPermissionsUpdateTemplateParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(permissionsUpdateTemplateHandler))
}

func permissionsUpdateTemplateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiPermissionsUpdateTemplateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiPermissionsUpdateTemplate(ctx, &params, authorizationHeader))
}
