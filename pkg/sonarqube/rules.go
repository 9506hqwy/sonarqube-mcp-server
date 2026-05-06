package sonarqube

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/sonar-client-go/pkg/sonarqube"
)

func registerRulesCreate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiRulesCreateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("rules_create",
		mcp.WithDescription("Create a custom rule.<br>Requires the 'Administer Quality Profiles' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rulesCreateHandler))
}

func rulesCreateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiRulesCreateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiRulesCreate(ctx, &params, authorizationHeader))
}

func registerRulesDelete(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiRulesDeleteParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("rules_delete",
		mcp.WithDescription("Delete custom rule.<br/>Requires the 'Administer Quality Profiles' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rulesDeleteHandler))
}

func rulesDeleteHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiRulesDeleteParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiRulesDelete(ctx, &params, authorizationHeader))
}

func registerRulesRepositories(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiRulesRepositoriesParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("rules_repositories",
		mcp.WithDescription("List available rule repositories"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rulesRepositoriesHandler))
}

func rulesRepositoriesHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiRulesRepositoriesParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiRulesRepositories(ctx, &params, authorizationHeader))
}

func registerRulesSearch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiRulesSearchParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("rules_search",
		mcp.WithDescription("Search for a collection of relevant rules matching a specified query.<br/>"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rulesSearchHandler))
}

func rulesSearchHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiRulesSearchParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiRulesSearch(ctx, &params, authorizationHeader))
}

func registerRulesShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiRulesShowParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("rules_show",
		mcp.WithDescription("Get detailed information about a rule<br>"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rulesShowHandler))
}

func rulesShowHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiRulesShowParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiRulesShow(ctx, &params, authorizationHeader))
}

func registerRulesTags(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiRulesTagsParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("rules_tags",
		mcp.WithDescription("List rule tags"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rulesTagsHandler))
}

func rulesTagsHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiRulesTagsParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiRulesTags(ctx, &params, authorizationHeader))
}

func registerRulesUpdate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&client.ApiRulesUpdateParams{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("rules_update",
		mcp.WithDescription("Update an existing rule.<br>Requires the 'Administer Quality Profiles' permission"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rulesUpdateHandler))
}

func rulesUpdateHandler(ctx context.Context, request mcp.CallToolRequest, params client.ApiRulesUpdateParams) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiRulesUpdate(ctx, &params, authorizationHeader))
}
