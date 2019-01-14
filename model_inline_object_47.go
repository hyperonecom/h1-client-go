/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject47 struct {
	Name string `json:"name"`
	Size float32 `json:"size"`
	Service string `json:"service"`
	Credential VaultCredential `json:"credential,omitempty"`
	Snapshot string `json:"snapshot,omitempty"`
	Tag map[string]interface{} `json:"tag,omitempty"`
}
