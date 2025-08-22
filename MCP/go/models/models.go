package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetRawMessageContentResponse represents the GetRawMessageContentResponse schema from the OpenAPI specification
type GetRawMessageContentResponse struct {
	Messagecontent interface{} `json:"messageContent"`
}

// PutRawMessageContentResponse represents the PutRawMessageContentResponse schema from the OpenAPI specification
type PutRawMessageContentResponse struct {
}

// PutRawMessageContentRequest represents the PutRawMessageContentRequest schema from the OpenAPI specification
type PutRawMessageContentRequest struct {
	Content interface{} `json:"content"`
}

// RawMessageContent represents the RawMessageContent schema from the OpenAPI specification
type RawMessageContent struct {
	S3reference interface{} `json:"s3Reference"`
}

// GetRawMessageContentRequest represents the GetRawMessageContentRequest schema from the OpenAPI specification
type GetRawMessageContentRequest struct {
}

// S3Reference represents the S3Reference schema from the OpenAPI specification
type S3Reference struct {
	Key interface{} `json:"key"`
	Objectversion interface{} `json:"objectVersion,omitempty"`
	Bucket interface{} `json:"bucket"`
}
