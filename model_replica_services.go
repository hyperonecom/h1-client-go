/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ReplicaServices struct {
	Type          string                 `json:"type,omitempty"`
	Name          string                 `json:"name,omitempty"`
	OneTime       bool                   `json:"oneTime,omitempty"`
	Billing       string                 `json:"billing,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty"`
	SourceService string                 `json:"sourceService,omitempty"`
	Quantity      float32                `json:"quantity,omitempty"`
}
