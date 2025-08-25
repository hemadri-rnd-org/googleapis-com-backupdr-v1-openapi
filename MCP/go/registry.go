package main

import (
	"github.com/backup-and-dr-service-api/mcp-server/config"
	"github.com/backup-and-dr-service-api/mcp-server/models"
	tools_projects "github.com/backup-and-dr-service-api/mcp-server/tools/projects"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_projects.CreateBackupdr_projects_locations_managementservers_setiampolicyTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_managementservers_testiampermissionsTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_operations_deleteTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_operations_getTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_listTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_operations_listTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_operations_cancelTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_managementservers_listTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_managementservers_createTool(cfg),
		tools_projects.CreateBackupdr_projects_locations_managementservers_getiampolicyTool(cfg),
	}
}
