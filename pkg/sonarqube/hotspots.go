package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerHotspotsSearch(s *server.MCPServer) {
	tool := mcp.NewTool("hotspots_search",
		mcp.WithDescription("Search for Security Hotpots. <br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects. <br>When issue indexation is in progress returns 503 service unavailable HTTP code."),
		mcp.WithString("branch",
			mcp.Description("Branch key. Not available in the community edition."),
		),
		mcp.WithString("cwe",
			mcp.Description("Comma-separated list of CWE numbers"),
		),
		mcp.WithString("files",
			mcp.Description("Comma-separated list of files. Returns only hotspots found in those files"),
		),
		mcp.WithString("hotspots",
			mcp.Description("Comma-separated list of Security Hotspot keys. This parameter is required unless projectKey is provided."),
		),
		mcp.WithString("inNewCodePeriod",
			mcp.Description("If 'inNewCodePeriod' is provided, only Security Hotspots created in the new code period are returned."),
		),
		mcp.WithString("onlyMine",
			mcp.Description("If 'projectKey' is provided, returns only Security Hotspots assigned to the current user"),
		),
		mcp.WithString("owaspAsvs-4.0",
			mcp.Description("Comma-separated list of OWASP ASVS v4.0 categories or rules."),
		),
		mcp.WithString("owaspAsvsLevel",
			mcp.Description("Filters hotspots with lower or equal OWASP ASVS level to the parameter value. Should be used in combination with the 'owaspAsvs-4.0' parameter."),
		),
		mcp.WithString("owaspTop10",
			mcp.Description("Comma-separated list of OWASP 2017 Top 10 lowercase categories."),
		),
		mcp.WithString("owaspTop10-2021",
			mcp.Description("Comma-separated list of OWASP 2021 Top 10 lowercase categories."),
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
		mcp.WithString("projectKey",
			mcp.Description("Key of the project or application. This parameter is required unless hotspots is provided."),
		),
		mcp.WithString("ps",
			mcp.Description("Page size. Must be greater than 0."),
		),
		mcp.WithString("pullRequest",
			mcp.Description("Pull request id. Not available in the community edition."),
		),
		mcp.WithString("resolution",
			mcp.Description("If 'projectKey' is provided and if status is 'REVIEWED', only Security Hotspots with the specified resolution are returned."),
		),
		mcp.WithString("sansTop25",
			mcp.Description("Comma-separated list of SANS Top 25 categories."),
		),
		mcp.WithString("sonarsourceSecurity",
			mcp.Description("Comma-separated list of SonarSource security categories. Use 'others' to select issues not associated with any category"),
		),
		mcp.WithString("status",
			mcp.Description("If 'projectKey' is provided, only Security Hotspots with the specified status are returned."),
		),
	)

	s.AddTool(tool, hotspotsSearchHandler)
}

func hotspotsSearchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseHotspotsSearch(request)
	return toResult(c.ApiHotspotsSearch(ctx, &params, authorizationHeader))
}

func parseHotspotsSearch(request mcp.CallToolRequest) client.ApiHotspotsSearchParams {
	params := client.ApiHotspotsSearchParams{}

	branch := request.GetString("branch", "")
	if branch != "" {
		params.Branch = &branch
	}

	cwe := request.GetString("cwe", "")
	if cwe != "" {
		params.Cwe = &cwe
	}

	files := request.GetString("files", "")
	if files != "" {
		params.Files = &files
	}

	hotspots := request.GetString("hotspots", "")
	if hotspots != "" {
		params.Hotspots = &hotspots
	}

	inNewCodePeriod := request.GetString("inNewCodePeriod", "")
	if inNewCodePeriod != "" {
		params.InNewCodePeriod = &inNewCodePeriod
	}

	onlyMine := request.GetString("onlyMine", "")
	if onlyMine != "" {
		params.OnlyMine = &onlyMine
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

	projectKey := request.GetString("projectKey", "")
	if projectKey != "" {
		params.ProjectKey = &projectKey
	}

	ps := request.GetString("ps", "")
	if ps != "" {
		params.Ps = &ps
	}

	pullRequest := request.GetString("pullRequest", "")
	if pullRequest != "" {
		params.PullRequest = &pullRequest
	}

	resolution := request.GetString("resolution", "")
	if resolution != "" {
		params.Resolution = &resolution
	}

	sansTop25 := request.GetString("sansTop25", "")
	if sansTop25 != "" {
		params.SansTop25 = &sansTop25
	}

	sonarsourceSecurity := request.GetString("sonarsourceSecurity", "")
	if sonarsourceSecurity != "" {
		params.SonarsourceSecurity = &sonarsourceSecurity
	}

	status := request.GetString("status", "")
	if status != "" {
		params.Status = &status
	}

	return params
}

func registerHotspotsShow(s *server.MCPServer) {
	tool := mcp.NewTool("hotspots_show",
		mcp.WithDescription("Provides the details of a Security Hotspot."),
		mcp.WithString("hotspot",
			mcp.Description("Key of the Security Hotspot"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, hotspotsShowHandler)
}

func hotspotsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseHotspotsShow(request)
	return toResult(c.ApiHotspotsShow(ctx, &params, authorizationHeader))
}

func parseHotspotsShow(request mcp.CallToolRequest) client.ApiHotspotsShowParams {
	params := client.ApiHotspotsShowParams{}

	hotspot := request.GetString("hotspot", "")
	if hotspot != "" {
		params.Hotspot = hotspot
	}

	return params
}
