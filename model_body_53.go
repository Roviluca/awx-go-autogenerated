/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Body53 struct {
	// Allow changing the SCM branch or revision in a job template that uses this project.
	AllowOverride string `json:"allow_override,omitempty"`
	Credential    int32  `json:"credential,omitempty"`
	// Local absolute file path containing a custom Python virtualenv to use
	CustomVirtualenv string `json:"custom_virtualenv,omitempty"`
	Description      string `json:"description,omitempty"`
	// Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.
	LocalPath string `json:"local_path,omitempty"`
	Name      string `json:"name"`
	// The organization used to determine access to this template.
	Organization int32 `json:"organization,omitempty"`
	// Specific branch, tag or commit to checkout.
	ScmBranch string `json:"scm_branch,omitempty"`
	// Discard any local changes before syncing the project.
	ScmClean string `json:"scm_clean,omitempty"`
	// Delete the project before syncing.
	ScmDeleteOnUpdate string `json:"scm_delete_on_update,omitempty"`
	// For git projects, an additional refspec to fetch.
	ScmRefspec string `json:"scm_refspec,omitempty"`
	// Specifies the source control system used to store the project.
	ScmType string `json:"scm_type,omitempty"`
	// The number of seconds after the last project update ran that a new project update will be launched as a job dependency.
	ScmUpdateCacheTimeout int32 `json:"scm_update_cache_timeout,omitempty"`
	// Update the project when a job is launched that uses the project.
	ScmUpdateOnLaunch string `json:"scm_update_on_launch,omitempty"`
	// The location where the project is stored.
	ScmUrl string `json:"scm_url,omitempty"`
	// The amount of time (in seconds) to run before the task is canceled.
	Timeout int32 `json:"timeout,omitempty"`
}