/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AgentResourceEvent struct {
	Id string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	CreatedOn string `json:"createdOn,omitempty"`
	Project string `json:"project,omitempty"`
}
