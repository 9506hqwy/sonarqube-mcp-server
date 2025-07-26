package sonarqube

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerMonitoringMetrics(s *server.MCPServer) {
	tool := mcp.NewTool("monitoring_metrics",
		mcp.WithDescription("Return monitoring metrics in Prometheus format. Support content type 'text/plain' (default) and 'application/openmetrics-text'. this endpoint can be access using a Bearer token, that needs to be defined in sonar.properties with the 'sonar.web.systemPasscode' key."),
	)

	s.AddTool(tool, monitoringMetricsHandler)
}

func monitoringMetricsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ApiMonitoringMetrics(ctx, authorizationHeader))
}
