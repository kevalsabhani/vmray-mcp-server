package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListSamplesInput struct {
}

type ListSamplesOutput struct {
}

func listSamples(ctx context.Context, req *mcp.CallToolRequest, input ListSamplesInput) (
	*mcp.CallToolResult,
	ListSamplesOutput,
	error,
) {
	var out ListSamplesOutput
	return nil, out, nil
}
