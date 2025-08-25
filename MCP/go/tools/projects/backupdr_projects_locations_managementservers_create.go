package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/backup-and-dr-service-api/mcp-server/config"
	"github.com/backup-and-dr-service-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Backupdr_projects_locations_managementservers_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["managementServerId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("managementServerId=%v", val))
		}
		if val, ok := args["requestId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("requestId=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ManagementServer
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/%s/managementServers%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Operation
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateBackupdr_projects_locations_managementservers_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_managementServers",
		mcp.WithDescription("Creates a new ManagementServer in a given project and location."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The management server project and location in the format `projects/{project_id}/locations/{location}`. In Cloud Backup and DR locations map to GCP regions, for example **us-central1**.")),
		mcp.WithString("managementServerId", mcp.Description("Required. The name of the management server to create. The name must be unique for the specified project and location.")),
		mcp.WithString("requestId", mcp.Description("Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).")),
		mcp.WithString("etag", mcp.Description("Input parameter: Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.")),
		mcp.WithArray("networks", mcp.Description("Input parameter: Required. VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.")),
		mcp.WithString("description", mcp.Description("Input parameter: Optional. The description of the ManagementServer instance (2048 characters or less).")),
		mcp.WithString("state", mcp.Description("Input parameter: Output only. The ManagementServer state.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. The time when the instance was updated.")),
		mcp.WithObject("workforceIdentityBasedManagementUri", mcp.Description("Input parameter: ManagementURI depending on the Workforce Identity i.e. either 1p or 3p.")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.")),
		mcp.WithObject("managementUri", mcp.Description("Input parameter: ManagementURI for the Management Server resource.")),
		mcp.WithString("oauth2ClientId", mcp.Description("Input parameter: Output only. The OAuth 2.0 client id is required to make API calls to the BackupDR instance API of this ManagementServer. This is the value that should be provided in the ‘aud’ field of the OIDC ID Token (see openid specification https://openid.net/specs/openid-connect-core-1_0.html#IDToken).")),
		mcp.WithString("type", mcp.Description("Input parameter: Optional. The type of the ManagementServer resource.")),
		mcp.WithString("name", mcp.Description("Input parameter: Output only. Identifier. The resource name.")),
		mcp.WithObject("workforceIdentityBasedOauth2ClientId", mcp.Description("Input parameter: OAuth Client ID depending on the Workforce Identity i.e. either 1p or 3p,")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The time when the instance was created.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Backupdr_projects_locations_managementservers_createHandler(cfg),
	}
}
