/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// ComputeProjectVmCreate struct for ComputeProjectVmCreate
type ComputeProjectVmCreate struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Image string `json:"image,omitempty"`
	Iso string `json:"iso,omitempty"`
	Username string `json:"username,omitempty"`
	UserMetadata string `json:"userMetadata,omitempty"`
	Start bool `json:"start,omitempty"`
	Credential []ComputeProjectVmCreateCredential `json:"credential,omitempty"`
	Disk []ComputeProjectVmCreateDisk `json:"disk,omitempty"`
	Netadp []ComputeProjectVmCreateNetadp `json:"netadp,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}
