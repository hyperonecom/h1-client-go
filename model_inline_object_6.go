/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject6 struct {
	Name string `json:"name,omitempty"`
	Action string `json:"action,omitempty"`
	Priority float32 `json:"priority,omitempty"`
	Filter []string `json:"filter,omitempty"`
	External []string `json:"external,omitempty"`
	Internal []string `json:"internal,omitempty"`
}