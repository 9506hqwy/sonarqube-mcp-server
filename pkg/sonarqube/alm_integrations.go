package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerAlmIntegrationsImportGitlabProject(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_import_gitlab_project",
		mcp.WithDescription("Import a GitLab project to SonarQube, creating a new project and configuring MR decoration<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
		mcp.WithString("gitlabProjectId",
			mcp.Description("GitLab project ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almIntegrationsImportGitlabProjectHandler)
}

func almIntegrationsImportGitlabProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsImportGitlabProject(request)
	return toResult(c.ApiAlmIntegrationsImportGitlabProject(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsImportGitlabProject(request mcp.CallToolRequest) client.ApiAlmIntegrationsImportGitlabProjectParams {
	params := client.ApiAlmIntegrationsImportGitlabProjectParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	gitlabProjectId := request.GetString("gitlabProjectId", "")
	if gitlabProjectId != "" {
		params.GitlabProjectId = gitlabProjectId
	}

	return params
}

func registerAlmIntegrationsListAzureProjects(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_list_azure_projects",
		mcp.WithDescription("List Azure projects<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almIntegrationsListAzureProjectsHandler)
}

func almIntegrationsListAzureProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsListAzureProjects(request)
	return toResult(c.ApiAlmIntegrationsListAzureProjects(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsListAzureProjects(request mcp.CallToolRequest) client.ApiAlmIntegrationsListAzureProjectsParams {
	params := client.ApiAlmIntegrationsListAzureProjectsParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	return params
}

func registerAlmIntegrationsListBitbucketserverProjects(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_list_bitbucketserver_projects",
		mcp.WithDescription("List the Bitbucket Server projects<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almIntegrationsListBitbucketserverProjectsHandler)
}

func almIntegrationsListBitbucketserverProjectsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsListBitbucketserverProjects(request)
	return toResult(c.ApiAlmIntegrationsListBitbucketserverProjects(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsListBitbucketserverProjects(request mcp.CallToolRequest) client.ApiAlmIntegrationsListBitbucketserverProjectsParams {
	params := client.ApiAlmIntegrationsListBitbucketserverProjectsParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	return params
}

func registerAlmIntegrationsSearchAzureRepos(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_search_azure_repos",
		mcp.WithDescription("Search the Azure repositories<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
		mcp.WithString("projectName",
			mcp.Description("Project name filter"),
		),
		mcp.WithString("searchQuery",
			mcp.Description("Search query filter"),
		),
	)

	s.AddTool(tool, almIntegrationsSearchAzureReposHandler)
}

func almIntegrationsSearchAzureReposHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsSearchAzureRepos(request)
	return toResult(c.ApiAlmIntegrationsSearchAzureRepos(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsSearchAzureRepos(request mcp.CallToolRequest) client.ApiAlmIntegrationsSearchAzureReposParams {
	params := client.ApiAlmIntegrationsSearchAzureReposParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	projectName := request.GetString("projectName", "")
	if projectName != "" {
		params.ProjectName = &projectName
	}

	searchQuery := request.GetString("searchQuery", "")
	if searchQuery != "" {
		params.SearchQuery = &searchQuery
	}

	return params
}

func registerAlmIntegrationsSearchBitbucketcloudRepos(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_search_bitbucketcloud_repos",
		mcp.WithDescription("Search the Bitbucket Cloud repositories<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 100"),
		),
		mcp.WithString("repositoryName",
			mcp.Description("Repository name filter"),
		),
	)

	s.AddTool(tool, almIntegrationsSearchBitbucketcloudReposHandler)
}

func almIntegrationsSearchBitbucketcloudReposHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsSearchBitbucketcloudRepos(request)
	return toResult(c.ApiAlmIntegrationsSearchBitbucketcloudRepos(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsSearchBitbucketcloudRepos(request mcp.CallToolRequest) client.ApiAlmIntegrationsSearchBitbucketcloudReposParams {
	params := client.ApiAlmIntegrationsSearchBitbucketcloudReposParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	repositoryName := request.GetString("repositoryName", "")
	if repositoryName != "" {
		params.RepositoryName = &repositoryName
	}

	return params
}

func registerAlmIntegrationsSearchBitbucketserverRepos(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_search_bitbucketserver_repos",
		mcp.WithDescription("Search the Bitbucket Server repositories with REPO_ADMIN access<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
		mcp.WithString("projectName",
			mcp.Description("Project name filter"),
		),
		mcp.WithString("repositoryName",
			mcp.Description("Repository name filter"),
		),
	)

	s.AddTool(tool, almIntegrationsSearchBitbucketserverReposHandler)
}

func almIntegrationsSearchBitbucketserverReposHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsSearchBitbucketserverRepos(request)
	return toResult(c.ApiAlmIntegrationsSearchBitbucketserverRepos(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsSearchBitbucketserverRepos(request mcp.CallToolRequest) client.ApiAlmIntegrationsSearchBitbucketserverReposParams {
	params := client.ApiAlmIntegrationsSearchBitbucketserverReposParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	projectName := request.GetString("projectName", "")
	if projectName != "" {
		params.ProjectName = &projectName
	}

	repositoryName := request.GetString("repositoryName", "")
	if repositoryName != "" {
		params.RepositoryName = &repositoryName
	}

	return params
}

func registerAlmIntegrationsSearchGitlabRepos(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_search_gitlab_repos",
		mcp.WithDescription("Search the GitLab projects.<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
		mcp.WithString("p",
			mcp.Description("1-based page number"),
		),
		mcp.WithString("projectName",
			mcp.Description("Project name filter"),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0 and less or equal than 500"),
		),
	)

	s.AddTool(tool, almIntegrationsSearchGitlabReposHandler)
}

func almIntegrationsSearchGitlabReposHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsSearchGitlabRepos(request)
	return toResult(c.ApiAlmIntegrationsSearchGitlabRepos(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsSearchGitlabRepos(request mcp.CallToolRequest) client.ApiAlmIntegrationsSearchGitlabReposParams {
	params := client.ApiAlmIntegrationsSearchGitlabReposParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	p := request.GetString("p", "")
	if p != "" {
		params.P = &p
	}

	projectName := request.GetString("projectName", "")
	if projectName != "" {
		params.ProjectName = &projectName
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	return params
}

func registerAlmIntegrationsSetPat(s *server.MCPServer) {
	tool := mcp.NewTool("alm_integrations_set_pat",
		mcp.WithDescription("Set a Personal Access Token for the given DevOps Platform setting<br/>Only valid for Azure DevOps, Bitbucket Server, GitLab and Bitbucket Cloud Setting<br/>Requires the 'Create Projects' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
		mcp.WithString("pat",
			mcp.Description("Personal Access Token"),
			mcp.Required(),
		),
		mcp.WithString("username",
			mcp.Description("Username"),
		),
	)

	s.AddTool(tool, almIntegrationsSetPatHandler)
}

func almIntegrationsSetPatHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmIntegrationsSetPat(request)
	return toResult(c.ApiAlmIntegrationsSetPat(ctx, &params, authorizationHeader))
}

func parseAlmIntegrationsSetPat(request mcp.CallToolRequest) client.ApiAlmIntegrationsSetPatParams {
	params := client.ApiAlmIntegrationsSetPatParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	pat := request.GetString("pat", "")
	if pat != "" {
		params.Pat = pat
	}

	username := request.GetString("username", "")
	if username != "" {
		params.Username = &username
	}

	return params
}
