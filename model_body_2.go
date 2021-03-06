/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Body2 struct {
	Description string `json:"description,omitempty"`
	// Enter injectors using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax.
	Injectors *interface{} `json:"injectors,omitempty"`
	// Enter inputs using either JSON or YAML syntax. Refer to the Ansible Tower documentation for example syntax.
	Inputs *interface{} `json:"inputs,omitempty"`
	Kind   string       `json:"kind"`
	Name   string       `json:"name"`
}
