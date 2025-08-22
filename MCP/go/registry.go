package main

import (
	"github.com/amazon-workmail-message-flow/mcp-server/config"
	"github.com/amazon-workmail-message-flow/mcp-server/models"
	tools_messages "github.com/amazon-workmail-message-flow/mcp-server/tools/messages"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_messages.CreatePutrawmessagecontentTool(cfg),
		tools_messages.CreateGetrawmessagecontentTool(cfg),
	}
}
