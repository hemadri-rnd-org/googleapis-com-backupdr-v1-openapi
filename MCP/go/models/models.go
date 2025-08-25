package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListOperationsResponse represents the ListOperationsResponse schema from the OpenAPI specification
type ListOperationsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
	Operations []Operation `json:"operations,omitempty"` // A list of operations that matches the specified filter in the request.
}

// OperationMetadata represents the OperationMetadata schema from the OpenAPI specification
type OperationMetadata struct {
	Apiversion string `json:"apiVersion,omitempty"` // Output only. API version used to start the operation.
	Createtime string `json:"createTime,omitempty"` // Output only. The time the operation was created.
	Endtime string `json:"endTime,omitempty"` // Output only. The time the operation finished running.
	Requestedcancellation bool `json:"requestedCancellation,omitempty"` // Output only. Identifies whether the user has requested cancellation of the operation. Operations that have successfully been cancelled have Operation.error value with a google.rpc.Status.code of 1, corresponding to `Code.CANCELLED`.
	Statusmessage string `json:"statusMessage,omitempty"` // Output only. Human-readable status of the operation, if any.
	Target string `json:"target,omitempty"` // Output only. Server-defined resource path for the target of the operation.
	Verb string `json:"verb,omitempty"` // Output only. Name of the verb executed by the operation.
	Additionalinfo map[string]interface{} `json:"additionalInfo,omitempty"` // Output only. AdditionalInfo contains additional Info related to backup plan association resource.
}

// ListLocationsResponse represents the ListLocationsResponse schema from the OpenAPI specification
type ListLocationsResponse struct {
	Locations []Location `json:"locations,omitempty"` // A list of locations that matches the specified filter in the request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
}

// CancelOperationRequest represents the CancelOperationRequest schema from the OpenAPI specification
type CancelOperationRequest struct {
}

// ManagementServer represents the ManagementServer schema from the OpenAPI specification
type ManagementServer struct {
	Name string `json:"name,omitempty"` // Output only. Identifier. The resource name.
	Workforceidentitybasedoauth2clientid WorkforceIdentityBasedOAuth2ClientID `json:"workforceIdentityBasedOauth2ClientId,omitempty"` // OAuth Client ID depending on the Workforce Identity i.e. either 1p or 3p,
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the instance was created.
	Etag string `json:"etag,omitempty"` // Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.
	Networks []NetworkConfig `json:"networks,omitempty"` // Required. VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.
	Description string `json:"description,omitempty"` // Optional. The description of the ManagementServer instance (2048 characters or less).
	State string `json:"state,omitempty"` // Output only. The ManagementServer state.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the instance was updated.
	Workforceidentitybasedmanagementuri WorkforceIdentityBasedManagementURI `json:"workforceIdentityBasedManagementUri,omitempty"` // ManagementURI depending on the Workforce Identity i.e. either 1p or 3p.
	Labels map[string]interface{} `json:"labels,omitempty"` // Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.
	Managementuri ManagementURI `json:"managementUri,omitempty"` // ManagementURI for the Management Server resource.
	Oauth2clientid string `json:"oauth2ClientId,omitempty"` // Output only. The OAuth 2.0 client id is required to make API calls to the BackupDR instance API of this ManagementServer. This is the value that should be provided in the ‘aud’ field of the OIDC ID Token (see openid specification https://openid.net/specs/openid-connect-core-1_0.html#IDToken).
	TypeField string `json:"type,omitempty"` // Optional. The type of the ManagementServer resource.
}

// Policy represents the Policy schema from the OpenAPI specification
type Policy struct {
	Version int `json:"version,omitempty"` // Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Auditconfigs []AuditConfig `json:"auditConfigs,omitempty"` // Specifies cloud audit logging configuration for this policy.
	Bindings []Binding `json:"bindings,omitempty"` // Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Etag string `json:"etag,omitempty"` // `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
}

// Binding represents the Binding schema from the OpenAPI specification
type Binding struct {
	Members []string `json:"members,omitempty"` // Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workforce identity pool. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/group/{group_id}`: All workforce identities in a group. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All workforce identities with a specific attribute value. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/*`: All identities in a workforce identity pool. * `principal://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workload identity pool. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/group/{group_id}`: A workload identity pool group. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All identities in a workload identity pool with a certain attribute. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/*`: All identities in a workload identity pool. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `deleted:principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: Deleted single identity in a workforce identity pool. For example, `deleted:principal://iam.googleapis.com/locations/global/workforcePools/my-pool-id/subject/my-subject-attribute-value`.
	Role string `json:"role,omitempty"` // Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`. For an overview of the IAM roles and permissions, see the [IAM documentation](https://cloud.google.com/iam/docs/roles-overview). For a list of the available pre-defined roles, see [here](https://cloud.google.com/iam/docs/understanding-roles).
	Condition Expr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
}

// TestIamPermissionsRequest represents the TestIamPermissionsRequest schema from the OpenAPI specification
type TestIamPermissionsRequest struct {
	Permissions []string `json:"permissions,omitempty"` // The set of permissions to check for the `resource`. Permissions with wildcards (such as `*` or `storage.*`) are not allowed. For more information see [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
}

// WorkforceIdentityBasedManagementURI represents the WorkforceIdentityBasedManagementURI schema from the OpenAPI specification
type WorkforceIdentityBasedManagementURI struct {
	Firstpartymanagementuri string `json:"firstPartyManagementUri,omitempty"` // Output only. First party Management URI for Google Identities.
	Thirdpartymanagementuri string `json:"thirdPartyManagementUri,omitempty"` // Output only. Third party Management URI for External Identity Providers.
}

// AuditConfig represents the AuditConfig schema from the OpenAPI specification
type AuditConfig struct {
	Auditlogconfigs []AuditLogConfig `json:"auditLogConfigs,omitempty"` // The configuration for logging of each type of permission.
	Service string `json:"service,omitempty"` // Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
}

// TestIamPermissionsResponse represents the TestIamPermissionsResponse schema from the OpenAPI specification
type TestIamPermissionsResponse struct {
	Permissions []string `json:"permissions,omitempty"` // A subset of `TestPermissionsRequest.permissions` that the caller is allowed.
}

// ManagementURI represents the ManagementURI schema from the OpenAPI specification
type ManagementURI struct {
	Api string `json:"api,omitempty"` // Output only. The ManagementServer AGM/RD API URL.
	Webui string `json:"webUi,omitempty"` // Output only. The ManagementServer AGM/RD WebUI URL.
}

// WorkforceIdentityBasedOAuth2ClientID represents the WorkforceIdentityBasedOAuth2ClientID schema from the OpenAPI specification
type WorkforceIdentityBasedOAuth2ClientID struct {
	Firstpartyoauth2clientid string `json:"firstPartyOauth2ClientId,omitempty"` // Output only. First party OAuth Client ID for Google Identities.
	Thirdpartyoauth2clientid string `json:"thirdPartyOauth2ClientId,omitempty"` // Output only. Third party OAuth Client ID for External Identity Providers.
}

// AuditLogConfig represents the AuditLogConfig schema from the OpenAPI specification
type AuditLogConfig struct {
	Exemptedmembers []string `json:"exemptedMembers,omitempty"` // Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
	Logtype string `json:"logType,omitempty"` // The log type that this config enables.
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Name string `json:"name,omitempty"` // Resource name for the location, which may vary between implementations. For example: `"projects/example-project/locations/us-east1"`
	Displayname string `json:"displayName,omitempty"` // The friendly name for this location, typically a nearby city name. For example, "Tokyo".
	Labels map[string]interface{} `json:"labels,omitempty"` // Cross-service attributes for the location. For example {"cloud.googleapis.com/region": "us-east1"}
	Locationid string `json:"locationId,omitempty"` // The canonical id for this location. For example: `"us-east1"`.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata. For example the available capacity at the given location.
}

// NetworkConfig represents the NetworkConfig schema from the OpenAPI specification
type NetworkConfig struct {
	Network string `json:"network,omitempty"` // Optional. The resource name of the Google Compute Engine VPC network to which the ManagementServer instance is connected.
	Peeringmode string `json:"peeringMode,omitempty"` // Optional. The network connect mode of the ManagementServer instance. For this version, only PRIVATE_SERVICE_ACCESS is supported.
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
}

// Expr represents the Expr schema from the OpenAPI specification
type Expr struct {
	Location string `json:"location,omitempty"` // Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Title string `json:"title,omitempty"` // Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Description string `json:"description,omitempty"` // Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Expression string `json:"expression,omitempty"` // Textual representation of an expression in Common Expression Language syntax.
}

// SetIamPolicyRequest represents the SetIamPolicyRequest schema from the OpenAPI specification
type SetIamPolicyRequest struct {
	Policy Policy `json:"policy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
	Updatemask string `json:"updateMask,omitempty"` // OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
}

// ListManagementServersResponse represents the ListManagementServersResponse schema from the OpenAPI specification
type ListManagementServersResponse struct {
	Managementservers []ManagementServer `json:"managementServers,omitempty"` // The list of ManagementServer instances in the project for the specified location. If the `{location}` value in the request is "-", the response contains a list of instances from all locations. In case any location is unreachable, the response will only return management servers in reachable locations and the 'unreachable' field will be populated with a list of unreachable locations.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token identifying a page of results the server should return.
	Unreachable []string `json:"unreachable,omitempty"` // Locations that could not be reached.
}
