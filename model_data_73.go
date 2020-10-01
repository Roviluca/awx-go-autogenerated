/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data73 struct {
	// 
	Description string `json:"description,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	// The amount of time (in seconds) before the approval node expires and fails.
	Timeout int32 `json:"timeout,omitempty"`
}
