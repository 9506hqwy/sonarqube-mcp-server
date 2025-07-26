package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerAlmSettingsCountBinding(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_count_binding",
		mcp.WithDescription("Count number of project bound to an DevOps Platform setting.<br/>Requires the 'Administer System' permission"),
		mcp.WithString("almSetting",
			mcp.Description("DevOps Platform setting key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsCountBindingHandler)
}

func almSettingsCountBindingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsCountBinding(request)
	return toResult(c.ApiAlmSettingsCountBinding(ctx, &params, authorizationHeader))
}

func parseAlmSettingsCountBinding(request mcp.CallToolRequest) client.ApiAlmSettingsCountBindingParams {
	params := client.ApiAlmSettingsCountBindingParams{}

	almSetting := request.GetString("almSetting", "")
	if almSetting != "" {
		params.AlmSetting = almSetting
	}

	return params
}

func registerAlmSettingsCreateAzure(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_create_azure",
		mcp.WithDescription("Create Azure instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("Unique key of the Azure Devops instance setting"),
			mcp.Required(),
		),
		mcp.WithString("personalAccessToken",
			mcp.Description("Azure Devops personal access token"),
			mcp.Required(),
		),
		mcp.WithString("url",
			mcp.Description("Azure API URL"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsCreateAzureHandler)
}

func almSettingsCreateAzureHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsCreateAzure(request)
	return toResult(c.ApiAlmSettingsCreateAzure(ctx, &params, authorizationHeader))
}

func parseAlmSettingsCreateAzure(request mcp.CallToolRequest) client.ApiAlmSettingsCreateAzureParams {
	params := client.ApiAlmSettingsCreateAzureParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	personalAccessToken := request.GetString("personalAccessToken", "")
	if personalAccessToken != "" {
		params.PersonalAccessToken = personalAccessToken
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerAlmSettingsCreateBitbucket(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_create_bitbucket",
		mcp.WithDescription("Create Bitbucket instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("Unique key of the Bitbucket instance setting"),
			mcp.Required(),
		),
		mcp.WithString("personalAccessToken",
			mcp.Description("Bitbucket personal access token"),
			mcp.Required(),
		),
		mcp.WithString("url",
			mcp.Description("BitBucket server API URL"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsCreateBitbucketHandler)
}

func almSettingsCreateBitbucketHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsCreateBitbucket(request)
	return toResult(c.ApiAlmSettingsCreateBitbucket(ctx, &params, authorizationHeader))
}

func parseAlmSettingsCreateBitbucket(request mcp.CallToolRequest) client.ApiAlmSettingsCreateBitbucketParams {
	params := client.ApiAlmSettingsCreateBitbucketParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	personalAccessToken := request.GetString("personalAccessToken", "")
	if personalAccessToken != "" {
		params.PersonalAccessToken = personalAccessToken
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerAlmSettingsCreateBitbucketcloud(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_create_bitbucketcloud",
		mcp.WithDescription("Configure a new instance of Bitbucket Cloud. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("clientId",
			mcp.Description("Bitbucket Cloud Client ID"),
			mcp.Required(),
		),
		mcp.WithString("clientSecret",
			mcp.Description("Bitbucket Cloud Client Secret"),
			mcp.Required(),
		),
		mcp.WithString("key",
			mcp.Description("Unique key of the Bitbucket Cloud setting"),
			mcp.Required(),
		),
		mcp.WithString("workspace",
			mcp.Description("Bitbucket Cloud workspace ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsCreateBitbucketcloudHandler)
}

func almSettingsCreateBitbucketcloudHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsCreateBitbucketcloud(request)
	return toResult(c.ApiAlmSettingsCreateBitbucketcloud(ctx, &params, authorizationHeader))
}

func parseAlmSettingsCreateBitbucketcloud(request mcp.CallToolRequest) client.ApiAlmSettingsCreateBitbucketcloudParams {
	params := client.ApiAlmSettingsCreateBitbucketcloudParams{}

	clientId := request.GetString("clientId", "")
	if clientId != "" {
		params.ClientId = clientId
	}

	clientSecret := request.GetString("clientSecret", "")
	if clientSecret != "" {
		params.ClientSecret = clientSecret
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	workspace := request.GetString("workspace", "")
	if workspace != "" {
		params.Workspace = workspace
	}

	return params
}

func registerAlmSettingsCreateGithub(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_create_github",
		mcp.WithDescription("Create GitHub instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("appId",
			mcp.Description("GitHub App ID"),
			mcp.Required(),
		),
		mcp.WithString("clientId",
			mcp.Description("GitHub App Client ID"),
			mcp.Required(),
		),
		mcp.WithString("clientSecret",
			mcp.Description("GitHub App Client Secret"),
			mcp.Required(),
		),
		mcp.WithString("key",
			mcp.Description("Unique key of the GitHub instance setting"),
			mcp.Required(),
		),
		mcp.WithString("privateKey",
			mcp.Description("GitHub App private key"),
			mcp.Required(),
		),
		mcp.WithString("url",
			mcp.Description("GitHub API URL"),
			mcp.Required(),
		),
		mcp.WithString("webhookSecret",
			mcp.Description("GitHub App Webhook Secret"),
		),
	)

	s.AddTool(tool, almSettingsCreateGithubHandler)
}

func almSettingsCreateGithubHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsCreateGithub(request)
	return toResult(c.ApiAlmSettingsCreateGithub(ctx, &params, authorizationHeader))
}

func parseAlmSettingsCreateGithub(request mcp.CallToolRequest) client.ApiAlmSettingsCreateGithubParams {
	params := client.ApiAlmSettingsCreateGithubParams{}

	appId := request.GetString("appId", "")
	if appId != "" {
		params.AppId = appId
	}

	clientId := request.GetString("clientId", "")
	if clientId != "" {
		params.ClientId = clientId
	}

	clientSecret := request.GetString("clientSecret", "")
	if clientSecret != "" {
		params.ClientSecret = clientSecret
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	privateKey := request.GetString("privateKey", "")
	if privateKey != "" {
		params.PrivateKey = privateKey
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	webhookSecret := request.GetString("webhookSecret", "")
	if webhookSecret != "" {
		params.WebhookSecret = &webhookSecret
	}

	return params
}

func registerAlmSettingsCreateGitlab(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_create_gitlab",
		mcp.WithDescription("Create GitLab instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("Unique key of the GitLab instance setting"),
			mcp.Required(),
		),
		mcp.WithString("personalAccessToken",
			mcp.Description("GitLab personal access token"),
			mcp.Required(),
		),
		mcp.WithString("url",
			mcp.Description("GitLab API URL"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsCreateGitlabHandler)
}

func almSettingsCreateGitlabHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsCreateGitlab(request)
	return toResult(c.ApiAlmSettingsCreateGitlab(ctx, &params, authorizationHeader))
}

func parseAlmSettingsCreateGitlab(request mcp.CallToolRequest) client.ApiAlmSettingsCreateGitlabParams {
	params := client.ApiAlmSettingsCreateGitlabParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	personalAccessToken := request.GetString("personalAccessToken", "")
	if personalAccessToken != "" {
		params.PersonalAccessToken = personalAccessToken
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerAlmSettingsDelete(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_delete",
		mcp.WithDescription("Delete an DevOps Platform Setting.<br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("DevOps Platform Setting key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsDeleteHandler)
}

func almSettingsDeleteHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsDelete(request)
	return toResult(c.ApiAlmSettingsDelete(ctx, &params, authorizationHeader))
}

func parseAlmSettingsDelete(request mcp.CallToolRequest) client.ApiAlmSettingsDeleteParams {
	params := client.ApiAlmSettingsDeleteParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}

func registerAlmSettingsGetBinding(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_get_binding",
		mcp.WithDescription("Get DevOps Platform binding of a given project.<br/>Requires the 'Administer' permission on the project"),
		mcp.WithString("project",
			mcp.Description("Project key"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsGetBindingHandler)
}

func almSettingsGetBindingHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsGetBinding(request)
	return toResult(c.ApiAlmSettingsGetBinding(ctx, &params, authorizationHeader))
}

func parseAlmSettingsGetBinding(request mcp.CallToolRequest) client.ApiAlmSettingsGetBindingParams {
	params := client.ApiAlmSettingsGetBindingParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = project
	}

	return params
}

func registerAlmSettingsList(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_list",
		mcp.WithDescription("List DevOps Platform setting available for a given project, sorted by DevOps Platform key<br/>Requires the 'Administer project' permission if the 'project' parameter is provided, requires the 'Create Projects' permission otherwise."),
		mcp.WithString("project",
			mcp.Description("Project key"),
		),
	)

	s.AddTool(tool, almSettingsListHandler)
}

func almSettingsListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsList(request)
	return toResult(c.ApiAlmSettingsList(ctx, &params, authorizationHeader))
}

func parseAlmSettingsList(request mcp.CallToolRequest) client.ApiAlmSettingsListParams {
	params := client.ApiAlmSettingsListParams{}

	project := request.GetString("project", "")
	if project != "" {
		params.Project = &project
	}

	return params
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
	tool := mcp.NewTool("alm_settings_update_azure",
		mcp.WithDescription("Update Azure instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("Unique key of the Azure instance setting"),
			mcp.Required(),
		),
		mcp.WithString("newKey",
			mcp.Description("Optional new value for an unique key of the Azure Devops instance setting"),
		),
		mcp.WithString("personalAccessToken",
			mcp.Description("Azure Devops personal access token"),
		),
		mcp.WithString("url",
			mcp.Description("Azure API URL"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsUpdateAzureHandler)
}

func almSettingsUpdateAzureHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsUpdateAzure(request)
	return toResult(c.ApiAlmSettingsUpdateAzure(ctx, &params, authorizationHeader))
}

func parseAlmSettingsUpdateAzure(request mcp.CallToolRequest) client.ApiAlmSettingsUpdateAzureParams {
	params := client.ApiAlmSettingsUpdateAzureParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	newKey := request.GetString("newKey", "")
	if newKey != "" {
		params.NewKey = &newKey
	}

	personalAccessToken := request.GetString("personalAccessToken", "")
	if personalAccessToken != "" {
		params.PersonalAccessToken = &personalAccessToken
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerAlmSettingsUpdateBitbucket(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_update_bitbucket",
		mcp.WithDescription("Update Bitbucket instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("Unique key of the Bitbucket instance setting"),
			mcp.Required(),
		),
		mcp.WithString("newKey",
			mcp.Description("Optional new value for an unique key of the Bitbucket instance setting"),
		),
		mcp.WithString("personalAccessToken",
			mcp.Description("Bitbucket personal access token"),
		),
		mcp.WithString("url",
			mcp.Description("Bitbucket API URL"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsUpdateBitbucketHandler)
}

func almSettingsUpdateBitbucketHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsUpdateBitbucket(request)
	return toResult(c.ApiAlmSettingsUpdateBitbucket(ctx, &params, authorizationHeader))
}

func parseAlmSettingsUpdateBitbucket(request mcp.CallToolRequest) client.ApiAlmSettingsUpdateBitbucketParams {
	params := client.ApiAlmSettingsUpdateBitbucketParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	newKey := request.GetString("newKey", "")
	if newKey != "" {
		params.NewKey = &newKey
	}

	personalAccessToken := request.GetString("personalAccessToken", "")
	if personalAccessToken != "" {
		params.PersonalAccessToken = &personalAccessToken
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerAlmSettingsUpdateBitbucketcloud(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_update_bitbucketcloud",
		mcp.WithDescription("Update Bitbucket Cloud Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("clientId",
			mcp.Description("Bitbucket Cloud Client ID"),
			mcp.Required(),
		),
		mcp.WithString("clientSecret",
			mcp.Description("Optional new value for the Bitbucket Cloud client secret"),
		),
		mcp.WithString("key",
			mcp.Description("Unique key of the Bitbucket Cloud setting"),
			mcp.Required(),
		),
		mcp.WithString("newKey",
			mcp.Description("Optional new value for an unique key of the Bitbucket Cloud setting"),
		),
		mcp.WithString("workspace",
			mcp.Description("Bitbucket Cloud workspace ID"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsUpdateBitbucketcloudHandler)
}

func almSettingsUpdateBitbucketcloudHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsUpdateBitbucketcloud(request)
	return toResult(c.ApiAlmSettingsUpdateBitbucketcloud(ctx, &params, authorizationHeader))
}

func parseAlmSettingsUpdateBitbucketcloud(request mcp.CallToolRequest) client.ApiAlmSettingsUpdateBitbucketcloudParams {
	params := client.ApiAlmSettingsUpdateBitbucketcloudParams{}

	clientId := request.GetString("clientId", "")
	if clientId != "" {
		params.ClientId = clientId
	}

	clientSecret := request.GetString("clientSecret", "")
	if clientSecret != "" {
		params.ClientSecret = &clientSecret
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	newKey := request.GetString("newKey", "")
	if newKey != "" {
		params.NewKey = &newKey
	}

	workspace := request.GetString("workspace", "")
	if workspace != "" {
		params.Workspace = workspace
	}

	return params
}

func registerAlmSettingsUpdateGithub(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_update_github",
		mcp.WithDescription("Update GitHub instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("appId",
			mcp.Description("GitHub API ID"),
			mcp.Required(),
		),
		mcp.WithString("clientId",
			mcp.Description("GitHub App Client ID"),
			mcp.Required(),
		),
		mcp.WithString("clientSecret",
			mcp.Description("GitHub App Client Secret"),
		),
		mcp.WithString("key",
			mcp.Description("Unique key of the GitHub instance setting"),
			mcp.Required(),
		),
		mcp.WithString("newKey",
			mcp.Description("Optional new value for an unique key of the GitHub instance setting"),
		),
		mcp.WithString("privateKey",
			mcp.Description("GitHub App private key"),
		),
		mcp.WithString("url",
			mcp.Description("GitHub API URL"),
			mcp.Required(),
		),
		mcp.WithString("webhookSecret",
			mcp.Description("GitHub App Webhook Secret"),
		),
	)

	s.AddTool(tool, almSettingsUpdateGithubHandler)
}

func almSettingsUpdateGithubHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsUpdateGithub(request)
	return toResult(c.ApiAlmSettingsUpdateGithub(ctx, &params, authorizationHeader))
}

func parseAlmSettingsUpdateGithub(request mcp.CallToolRequest) client.ApiAlmSettingsUpdateGithubParams {
	params := client.ApiAlmSettingsUpdateGithubParams{}

	appId := request.GetString("appId", "")
	if appId != "" {
		params.AppId = appId
	}

	clientId := request.GetString("clientId", "")
	if clientId != "" {
		params.ClientId = clientId
	}

	clientSecret := request.GetString("clientSecret", "")
	if clientSecret != "" {
		params.ClientSecret = &clientSecret
	}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	newKey := request.GetString("newKey", "")
	if newKey != "" {
		params.NewKey = &newKey
	}

	privateKey := request.GetString("privateKey", "")
	if privateKey != "" {
		params.PrivateKey = &privateKey
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	webhookSecret := request.GetString("webhookSecret", "")
	if webhookSecret != "" {
		params.WebhookSecret = &webhookSecret
	}

	return params
}

func registerAlmSettingsUpdateGitlab(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_update_gitlab",
		mcp.WithDescription("Update GitLab instance Setting. <br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("Unique key of the GitLab instance setting"),
			mcp.Required(),
		),
		mcp.WithString("newKey",
			mcp.Description("Optional new value for an unique key of the GitLab instance setting"),
		),
		mcp.WithString("personalAccessToken",
			mcp.Description("GitLab personal access token"),
		),
		mcp.WithString("url",
			mcp.Description("GitLab API URL"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsUpdateGitlabHandler)
}

func almSettingsUpdateGitlabHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsUpdateGitlab(request)
	return toResult(c.ApiAlmSettingsUpdateGitlab(ctx, &params, authorizationHeader))
}

func parseAlmSettingsUpdateGitlab(request mcp.CallToolRequest) client.ApiAlmSettingsUpdateGitlabParams {
	params := client.ApiAlmSettingsUpdateGitlabParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	newKey := request.GetString("newKey", "")
	if newKey != "" {
		params.NewKey = &newKey
	}

	personalAccessToken := request.GetString("personalAccessToken", "")
	if personalAccessToken != "" {
		params.PersonalAccessToken = &personalAccessToken
	}

	url := request.GetString("url", "")
	if url != "" {
		params.Url = url
	}

	return params
}

func registerAlmSettingsValidate(s *server.MCPServer) {
	tool := mcp.NewTool("alm_settings_validate",
		mcp.WithDescription("Validate an DevOps Platform Setting by checking connectivity and permissions<br/>Requires the 'Administer System' permission"),
		mcp.WithString("key",
			mcp.Description("Unique key of the DevOps Platform settings"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, almSettingsValidateHandler)
}

func almSettingsValidateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseAlmSettingsValidate(request)
	return toResult(c.ApiAlmSettingsValidate(ctx, &params, authorizationHeader))
}

func parseAlmSettingsValidate(request mcp.CallToolRequest) client.ApiAlmSettingsValidateParams {
	params := client.ApiAlmSettingsValidateParams{}

	key := request.GetString("key", "")
	if key != "" {
		params.Key = key
	}

	return params
}
