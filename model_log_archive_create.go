/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LogArchiveCreate struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Credential VaultCreateCredential `json:"credential,omitempty"`
	Retention float32 `json:"retention,omitempty"`
	Tag map[string]interface{} `json:"tag,omitempty"`
}
