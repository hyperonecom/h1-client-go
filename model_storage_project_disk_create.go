/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// StorageProjectDiskCreate struct for StorageProjectDiskCreate
type StorageProjectDiskCreate struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Size float32 `json:"size"`
	Source interface{} `json:"source,omitempty"`
	Vm string `json:"vm,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}
