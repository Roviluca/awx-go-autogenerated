/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Body82 struct {
	AskLimitOnLaunch     string `json:"ask_limit_on_launch,omitempty"`
	AskScmBranchOnLaunch string `json:"ask_scm_branch_on_launch,omitempty"`
	ExtraVars            string `json:"extra_vars,omitempty"`
	Inventory            int32  `json:"inventory,omitempty"`
	Limit                string `json:"limit,omitempty"`
	ScmBranch            string `json:"scm_branch,omitempty"`
}
