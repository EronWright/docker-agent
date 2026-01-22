package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// CreateMessageHandler is a function type that handles sampling/createMessage requests from MCP servers.
type CreateMessageHandler func(ctx context.Context, req *mcp.CreateMessageParams) (*mcp.CreateMessageResult, error)
