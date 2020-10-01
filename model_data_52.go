/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data52 struct {
	// 
	AllowSimultaneous bool `json:"allow_simultaneous,omitempty"`
	// 
	AskInventoryOnLaunch bool `json:"ask_inventory_on_launch,omitempty"`
	// 
	AskLimitOnLaunch bool `json:"ask_limit_on_launch,omitempty"`
	// 
	AskScmBranchOnLaunch bool `json:"ask_scm_branch_on_launch,omitempty"`
	// 
	AskVariablesOnLaunch bool `json:"ask_variables_on_launch,omitempty"`
	// 
	Description string `json:"description,omitempty"`
	// 
	ExtraVars string `json:"extra_vars,omitempty"`
	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int32 `json:"inventory,omitempty"`
	// 
	Limit string `json:"limit,omitempty"`
	// 
	Name string `json:"name"`
	// The organization used to determine access to this template.
	Organization int32 `json:"organization,omitempty"`
	// 
	ScmBranch string `json:"scm_branch,omitempty"`
	// 
	SurveyEnabled bool `json:"survey_enabled,omitempty"`
	// Personal Access Token for posting back the status to the service API
	WebhookCredential int32 `json:"webhook_credential,omitempty"`
	// Service that webhook requests will be accepted from
	WebhookService string `json:"webhook_service,omitempty"`
}
