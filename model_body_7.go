/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Body7 struct {
	BecomeEnabled string `json:"become_enabled,omitempty"`
	Credential    int32  `json:"credential,omitempty"`
	DiffMode      string `json:"diff_mode,omitempty"`
	ExtraVars     string `json:"extra_vars,omitempty"`
	Forks         int32  `json:"forks,omitempty"`
	Inventory     int32  `json:"inventory,omitempty"`
	JobType       string `json:"job_type,omitempty"`
	Limit         string `json:"limit,omitempty"`
	ModuleArgs    string `json:"module_args,omitempty"`
	ModuleName    string `json:"module_name,omitempty"`
	Verbosity     string `json:"verbosity,omitempty"`
}