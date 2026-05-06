package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerAlmSettingsCountBinding(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsCountBindingParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_count_binding",
		mcp.WithDescription("Count number of project bound to an DevOps Platform setting.<br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsCountBindingHandler))
}

func almSettingsCountBindingHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsCountBindingParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsCountBinding(ctx, &params, authorizationHeader))
}

func registerAlmSettingsCreateAzure(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsCreateAzureParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_create_azure",
		mcp.WithDescription("Create Azure instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsCreateAzureHandler))
}

func almSettingsCreateAzureHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsCreateAzureParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsCreateAzure(ctx, &params, authorizationHeader))
}

func registerAlmSettingsCreateBitbucket(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsCreateBitbucketParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_create_bitbucket",
		mcp.WithDescription("Create Bitbucket instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsCreateBitbucketHandler))
}

func almSettingsCreateBitbucketHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsCreateBitbucketParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsCreateBitbucket(ctx, &params, authorizationHeader))
}

func registerAlmSettingsCreateBitbucketcloud(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsCreateBitbucketcloudParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_create_bitbucketcloud",
		mcp.WithDescription("Configure a new instance of Bitbucket Cloud. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsCreateBitbucketcloudHandler))
}

func almSettingsCreateBitbucketcloudHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsCreateBitbucketcloudParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsCreateBitbucketcloud(ctx, &params, authorizationHeader))
}

func registerAlmSettingsCreateGithub(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsCreateGithubParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_create_github",
		mcp.WithDescription("Create GitHub instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsCreateGithubHandler))
}

func almSettingsCreateGithubHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsCreateGithubParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsCreateGithub(ctx, &params, authorizationHeader))
}

func registerAlmSettingsCreateGitlab(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsCreateGitlabParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_create_gitlab",
		mcp.WithDescription("Create GitLab instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsCreateGitlabHandler))
}

func almSettingsCreateGitlabHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsCreateGitlabParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsCreateGitlab(ctx, &params, authorizationHeader))
}

func registerAlmSettingsDelete(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsDeleteParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_delete",
		mcp.WithDescription("Delete an DevOps Platform Setting.<br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsDeleteHandler))
}

func almSettingsDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsDelete(ctx, &params, authorizationHeader))
}

func registerAlmSettingsGetBinding(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsGetBindingParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_get_binding",
		mcp.WithDescription("Get DevOps Platform binding of a given project.<br/>Requires the 'Administer' permission on the project"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsGetBindingHandler))
}

func almSettingsGetBindingHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsGetBindingParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsGetBinding(ctx, &params, authorizationHeader))
}

func registerAlmSettingsList(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsListParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_list",
		mcp.WithDescription("List DevOps Platform setting available for a given project, sorted by DevOps Platform key<br/>Requires the 'Administer project' permission if the 'project' parameter is provided, requires the 'Create Projects' permission otherwise."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsListHandler))
}

func almSettingsListHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsListParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsList(ctx, &params, authorizationHeader))
}

func registerAlmSettingsListDefinitions(s *server.MCPServer) {

	tool := mcp.NewTool("alm_settings_list_definitions",
		mcp.WithDescription("List DevOps Platform Settings, sorted by created date.<br/>Requires the 'Administer System' permission"),
	)

	s.AddTool(tool, almSettingsListDefinitionsHandler)
}

func almSettingsListDefinitionsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsListDefinitions(ctx, authorizationHeader))
}

func registerAlmSettingsUpdateAzure(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsUpdateAzureParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_update_azure",
		mcp.WithDescription("Update Azure instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsUpdateAzureHandler))
}

func almSettingsUpdateAzureHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsUpdateAzureParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsUpdateAzure(ctx, &params, authorizationHeader))
}

func registerAlmSettingsUpdateBitbucket(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsUpdateBitbucketParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_update_bitbucket",
		mcp.WithDescription("Update Bitbucket instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsUpdateBitbucketHandler))
}

func almSettingsUpdateBitbucketHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsUpdateBitbucketParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsUpdateBitbucket(ctx, &params, authorizationHeader))
}

func registerAlmSettingsUpdateBitbucketcloud(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsUpdateBitbucketcloudParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_update_bitbucketcloud",
		mcp.WithDescription("Update Bitbucket Cloud Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsUpdateBitbucketcloudHandler))
}

func almSettingsUpdateBitbucketcloudHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsUpdateBitbucketcloudParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsUpdateBitbucketcloud(ctx, &params, authorizationHeader))
}

func registerAlmSettingsUpdateGithub(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsUpdateGithubParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_update_github",
		mcp.WithDescription("Update GitHub instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsUpdateGithubHandler))
}

func almSettingsUpdateGithubHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsUpdateGithubParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsUpdateGithub(ctx, &params, authorizationHeader))
}

func registerAlmSettingsUpdateGitlab(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsUpdateGitlabParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_update_gitlab",
		mcp.WithDescription("Update GitLab instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsUpdateGitlabHandler))
}

func almSettingsUpdateGitlabHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsUpdateGitlabParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsUpdateGitlab(ctx, &params, authorizationHeader))
}

func registerAlmSettingsValidate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmSettingsValidateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_settings_validate",
		mcp.WithDescription("Validate an DevOps Platform Setting by checking connectivity and permissions<br/>Requires the 'Administer System' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almSettingsValidateHandler))
}

func almSettingsValidateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmSettingsValidateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmSettingsValidate(ctx, &params, authorizationHeader))
}
