/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data24 struct {
	// 
	Description string `json:"description,omitempty"`
	// 
	Name string `json:"name"`
	// Organization owning this inventory script
	Organization int32 `json:"organization"`
	// 
	Script string `json:"script"`
}
