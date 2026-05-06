package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerAlmIntegrationsImportGitlabProject(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsImportGitlabProjectParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_import_gitlab_project",
		mcp.WithDescription("Import a GitLab project to SonarQube, creating a new project and configuring MR decoration<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsImportGitlabProjectHandler))
}

func almIntegrationsImportGitlabProjectHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsImportGitlabProjectParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsImportGitlabProject(ctx, &params, authorizationHeader))
}

func registerAlmIntegrationsListAzureProjects(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsListAzureProjectsParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_list_azure_projects",
		mcp.WithDescription("List Azure projects<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsListAzureProjectsHandler))
}

func almIntegrationsListAzureProjectsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsListAzureProjectsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsListAzureProjects(ctx, &params, authorizationHeader))
}

func registerAlmIntegrationsListBitbucketserverProjects(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsListBitbucketserverProjectsParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_list_bitbucketserver_projects",
		mcp.WithDescription("List the Bitbucket Server projects<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsListBitbucketserverProjectsHandler))
}

func almIntegrationsListBitbucketserverProjectsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsListBitbucketserverProjectsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsListBitbucketserverProjects(ctx, &params, authorizationHeader))
}

func registerAlmIntegrationsSearchAzureRepos(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsSearchAzureReposParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_search_azure_repos",
		mcp.WithDescription("Search the Azure repositories<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsSearchAzureReposHandler))
}

func almIntegrationsSearchAzureReposHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsSearchAzureReposParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsSearchAzureRepos(ctx, &params, authorizationHeader))
}

func registerAlmIntegrationsSearchBitbucketcloudRepos(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsSearchBitbucketcloudReposParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_search_bitbucketcloud_repos",
		mcp.WithDescription("Search the Bitbucket Cloud repositories<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsSearchBitbucketcloudReposHandler))
}

func almIntegrationsSearchBitbucketcloudReposHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsSearchBitbucketcloudReposParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsSearchBitbucketcloudRepos(ctx, &params, authorizationHeader))
}

func registerAlmIntegrationsSearchBitbucketserverRepos(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsSearchBitbucketserverReposParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_search_bitbucketserver_repos",
		mcp.WithDescription("Search the Bitbucket Server repositories with REPO_ADMIN access<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsSearchBitbucketserverReposHandler))
}

func almIntegrationsSearchBitbucketserverReposHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsSearchBitbucketserverReposParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsSearchBitbucketserverRepos(ctx, &params, authorizationHeader))
}

func registerAlmIntegrationsSearchGitlabRepos(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsSearchGitlabReposParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_search_gitlab_repos",
		mcp.WithDescription("Search the GitLab projects.<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsSearchGitlabReposHandler))
}

func almIntegrationsSearchGitlabReposHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsSearchGitlabReposParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsSearchGitlabRepos(ctx, &params, authorizationHeader))
}

func registerAlmIntegrationsSetPat(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiAlmIntegrationsSetPatParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("alm_integrations_set_pat",
		mcp.WithDescription("Set a Personal Access Token for the given DevOps Platform setting<br/>Only valid for Azure DevOps, Bitbucket Server, GitLab and Bitbucket Cloud Setting<br/>Requires the 'Create Projects' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(almIntegrationsSetPatHandler))
}

func almIntegrationsSetPatHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiAlmIntegrationsSetPatParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiAlmIntegrationsSetPat(ctx, &params, authorizationHeader))
}
