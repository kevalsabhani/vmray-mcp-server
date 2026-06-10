package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kevalsabhani/vmray-mcp-server/config"
	"github.com/kevalsabhani/vmray-mcp-server/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("failed to load env %v", err)
		os.Exit(1)
	}

	// Load config
	cfg := config.Load()
	slog.Info("Config loaded.")

	// MCP server
	srv := mcp.NewServer(
		&mcp.Implementation{
			Name:    cfg.Name,
			Version: cfg.Version,
		},
		nil,
	)

	tools.Register(srv, cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	slog.Info("starting VMRay MCP server", "name", cfg.Name, "version", cfg.Version)
	if err := srv.Run(ctx, &mcp.StdioTransport{}); err != nil {
		slog.Error("server exited", "err", err)
		os.Exit(1)
	}
}
