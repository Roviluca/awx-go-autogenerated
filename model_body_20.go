/*
 * Ansible Tower API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Body20 struct {
	Credential      int32  `json:"credential,omitempty"`
	Name            string `json:"name"`
	PodSpecOverride string `json:"pod_spec_override,omitempty"`
	// List of exact-match Instances that will be assigned to this group
	PolicyInstanceList []string `json:"policy_instance_list,omitempty"`
	// Static minimum number of Instances that will be automatically assign to this group when new instances come online.
	PolicyInstanceMinimum int32 `json:"policy_instance_minimum,omitempty"`
	// Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.
	PolicyInstancePercentage int32 `json:"policy_instance_percentage,omitempty"`
}
