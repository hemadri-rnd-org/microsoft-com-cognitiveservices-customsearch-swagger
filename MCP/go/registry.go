package main

import (
	"github.com/custom-search-client/mcp-server/config"
	"github.com/custom-search-client/mcp-server/models"
	tools_customsearch "github.com/custom-search-client/mcp-server/tools/customsearch"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_customsearch.CreateCustominstance_searchTool(cfg),
	}
}
