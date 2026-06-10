package tools

import (
	"github.com/kevalsabhani/vmray-mcp-server/config"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Register(srv *mcp.Server, cfg *config.Config) {
	mcp.AddTool(
		srv,
		&mcp.Tool{
			Name:        "list_samples",
			Description: "List samples",
		},
		listSamples,
	)
}
