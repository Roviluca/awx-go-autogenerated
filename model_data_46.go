/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data46 struct {
	// 
	AllowSimultaneous bool `json:"allow_simultaneous,omitempty"`
	// 
	AskCredentialOnLaunch bool `json:"ask_credential_on_launch,omitempty"`
	// 
	AskDiffModeOnLaunch bool `json:"ask_diff_mode_on_launch,omitempty"`
	// 
	AskInventoryOnLaunch bool `json:"ask_inventory_on_launch,omitempty"`
	// 
	AskJobTypeOnLaunch bool `json:"ask_job_type_on_launch,omitempty"`
	// 
	AskLimitOnLaunch bool `json:"ask_limit_on_launch,omitempty"`
	// 
	AskScmBranchOnLaunch bool `json:"ask_scm_branch_on_launch,omitempty"`
	// 
	AskSkipTagsOnLaunch bool `json:"ask_skip_tags_on_launch,omitempty"`
	// 
	AskTagsOnLaunch bool `json:"ask_tags_on_launch,omitempty"`
	// 
	AskVariablesOnLaunch bool `json:"ask_variables_on_launch,omitempty"`
	// 
	AskVerbosityOnLaunch bool `json:"ask_verbosity_on_launch,omitempty"`
	// 
	BecomeEnabled bool `json:"become_enabled,omitempty"`
	// Local absolute file path containing a custom Python virtualenv to use
	CustomVirtualenv string `json:"custom_virtualenv,omitempty"`
	// 
	Description string `json:"description,omitempty"`
	// If enabled, textual changes made to any templated files on the host are shown in the standard output
	DiffMode bool `json:"diff_mode,omitempty"`
	// 
	ExtraVars string `json:"extra_vars,omitempty"`
	// 
	ForceHandlers bool `json:"force_handlers,omitempty"`
	// 
	Forks int32 `json:"forks,omitempty"`
	// 
	HostConfigKey string `json:"host_config_key,omitempty"`
	// 
	Inventory int32 `json:"inventory,omitempty"`
	// The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1.
	JobSliceCount int32 `json:"job_slice_count,omitempty"`
	// 
	JobTags string `json:"job_tags,omitempty"`
	// 
	JobType string `json:"job_type,omitempty"`
	// 
	Limit string `json:"limit,omitempty"`
	// 
	Name string `json:"name"`
	// 
	Playbook string `json:"playbook,omitempty"`
	// 
	Project string `json:"project,omitempty"`
	// Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true.
	ScmBranch string `json:"scm_branch,omitempty"`
	// 
	SkipTags string `json:"skip_tags,omitempty"`
	// 
	StartAtTask string `json:"start_at_task,omitempty"`
	// 
	SurveyEnabled bool `json:"survey_enabled,omitempty"`
	// The amount of time (in seconds) to run before the task is canceled.
	Timeout int32 `json:"timeout,omitempty"`
	// If enabled, Tower will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible.
	UseFactCache bool `json:"use_fact_cache,omitempty"`
	// 
	Verbosity string `json:"verbosity,omitempty"`
	// Personal Access Token for posting back the status to the service API
	WebhookCredential int32 `json:"webhook_credential,omitempty"`
	// Service that webhook requests will be accepted from
	WebhookService string `json:"webhook_service,omitempty"`
}
