package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerIssuesAddComment(s *server.MCPServer) {
	tool := mcp.NewTool("issues_add_comment",
		mcp.WithDescription("Add a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue."),
		mcp.WithInputSchema[client.ApiIssuesAddCommentParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesAddCommentHandler))
}

func issuesAddCommentHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesAddCommentParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesAddComment(ctx, &params, authorizationHeader))
}

func registerIssuesAssign(s *server.MCPServer) {
	tool := mcp.NewTool("issues_assign",
		mcp.WithDescription("Assign/Unassign an issue. Requires authentication and Browse permission on project"),
		mcp.WithInputSchema[client.ApiIssuesAssignParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesAssignHandler))
}

func issuesAssignHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesAssignParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesAssign(ctx, &params, authorizationHeader))
}

func registerIssuesAuthors(s *server.MCPServer) {
	tool := mcp.NewTool("issues_authors",
		mcp.WithDescription("Search SCM accounts which match a given query.<br/>Requires authentication.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code."),
		mcp.WithInputSchema[client.ApiIssuesAuthorsParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesAuthorsHandler))
}

func issuesAuthorsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesAuthorsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesAuthors(ctx, &params, authorizationHeader))
}

func registerIssuesBulkChange(s *server.MCPServer) {
	tool := mcp.NewTool("issues_bulk_change",
		mcp.WithDescription("Bulk change on issues. Up to 500 issues can be updated. <br/>Requires authentication."),
		mcp.WithInputSchema[client.ApiIssuesBulkChangeParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesBulkChangeHandler))
}

func issuesBulkChangeHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesBulkChangeParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesBulkChange(ctx, &params, authorizationHeader))
}

func registerIssuesChangelog(s *server.MCPServer) {
	tool := mcp.NewTool("issues_changelog",
		mcp.WithDescription("Display changelog of an issue.<br/>Requires the 'Browse' permission on the project of the specified issue."),
		mcp.WithInputSchema[client.ApiIssuesChangelogParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesChangelogHandler))
}

func issuesChangelogHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesChangelogParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesChangelog(ctx, &params, authorizationHeader))
}

func registerIssuesDeleteComment(s *server.MCPServer) {
	tool := mcp.NewTool("issues_delete_comment",
		mcp.WithDescription("Delete a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue."),
		mcp.WithInputSchema[client.ApiIssuesDeleteCommentParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesDeleteCommentHandler))
}

func issuesDeleteCommentHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesDeleteCommentParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesDeleteComment(ctx, &params, authorizationHeader))
}

func registerIssuesDoTransition(s *server.MCPServer) {
	tool := mcp.NewTool("issues_do_transition",
		mcp.WithDescription("Do workflow transition on an issue. Requires authentication and Browse permission on project.<br/>The transitions 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.<br/>The transitions involving security hotspots require the permission 'Administer Security Hotspot'."),
		mcp.WithInputSchema[client.ApiIssuesDoTransitionParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesDoTransitionHandler))
}

func issuesDoTransitionHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesDoTransitionParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesDoTransition(ctx, &params, authorizationHeader))
}

func registerIssuesEditComment(s *server.MCPServer) {
	tool := mcp.NewTool("issues_edit_comment",
		mcp.WithDescription("Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue."),
		mcp.WithInputSchema[client.ApiIssuesEditCommentParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesEditCommentHandler))
}

func issuesEditCommentHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesEditCommentParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesEditComment(ctx, &params, authorizationHeader))
}

func registerIssuesReindex(s *server.MCPServer) {
	tool := mcp.NewTool("issues_reindex",
		mcp.WithDescription("Reindex issues for a project.<br> Require 'Administer System' permission."),
		mcp.WithInputSchema[client.ApiIssuesReindexParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesReindexHandler))
}

func issuesReindexHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesReindexParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesReindex(ctx, &params, authorizationHeader))
}

func registerIssuesSearch(s *server.MCPServer) {
	tool := mcp.NewTool("issues_search",
		mcp.WithDescription("Search for issues.<br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code."),
		mcp.WithInputSchema[client.ApiIssuesSearchParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesSearchHandler))
}

func issuesSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesSearch(ctx, &params, authorizationHeader))
}

func registerIssuesSetSeverity(s *server.MCPServer) {
	tool := mcp.NewTool("issues_set_severity",
		mcp.WithDescription("Change severity.<br/>Requires the following permissions:<ul> <li>'Authentication'</li> <li>'Browse' rights on project of the specified issue</li> <li>'Administer Issues' rights on project of the specified issue</li></ul>"),
		mcp.WithInputSchema[client.ApiIssuesSetSeverityParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesSetSeverityHandler))
}

func issuesSetSeverityHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesSetSeverityParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesSetSeverity(ctx, &params, authorizationHeader))
}

func registerIssuesSetTags(s *server.MCPServer) {
	tool := mcp.NewTool("issues_set_tags",
		mcp.WithDescription("Set tags on an issue. <br/>Requires authentication and Browse permission on project"),
		mcp.WithInputSchema[client.ApiIssuesSetTagsParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesSetTagsHandler))
}

func issuesSetTagsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesSetTagsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesSetTags(ctx, &params, authorizationHeader))
}

func registerIssuesSetType(s *server.MCPServer) {
	tool := mcp.NewTool("issues_set_type",
		mcp.WithDescription("Change type of issue, for instance from 'code smell' to 'bug'.<br/>Requires the following permissions:<ul> <li>'Authentication'</li> <li>'Browse' rights on project of the specified issue</li> <li>'Administer Issues' rights on project of the specified issue</li></ul>"),
		mcp.WithInputSchema[client.ApiIssuesSetTypeParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesSetTypeHandler))
}

func issuesSetTypeHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesSetTypeParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesSetType(ctx, &params, authorizationHeader))
}

func registerIssuesTags(s *server.MCPServer) {
	tool := mcp.NewTool("issues_tags",
		mcp.WithDescription("List tags matching a given query"),
		mcp.WithInputSchema[client.ApiIssuesTagsParams](),
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesTagsHandler))
}

func issuesTagsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiIssuesTagsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiIssuesTags(ctx, &params, authorizationHeader))
}
